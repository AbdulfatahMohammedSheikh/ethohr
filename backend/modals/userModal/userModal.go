package usermodal

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"slices"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	log "github.com/sirupsen/logrus"
)

type User struct {
	// TODO: make [Name] and [Email] unique
	Id    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	// [Password] is the hash of the user password not the actual one.
	Password string   `json:"password"`
	Role     string   `json:"role"`
	IsBand   bool     `json:"IsBand"`
	Tags     []string `json:"tags"`
}

// TODO: check if pointers are better option than raw objects
func New(name, email, phone, password, role string) User {
	return User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
		Role:     role,
		IsBand:   false,
	}
}

func (u *User) Encode() string {
	return url.Values{
		"name":     {u.Name},
		"email":    {u.Email},
		"phone":    {u.Phone},
		"password": {u.Password},
		"role":     {u.Role},
		"isBand":   {"false"},
	}.Encode()
}

func (u *User) Create(a *surreal.AppRepository) error {

	_, err := rolemodal.HasRoleWithId(a, u.Role)

	if nil != err {
		log.Error(err.Error())
		return err
	}

	// TODO: make sure that name is unique
	_, err = a.Db.Create("users", u)

	if nil != err {
		log.Error(err.Error())
		return err
	}

	return nil
}

// Returns User based on the quary [q] and kets [m]
func Find(a *surreal.AppRepository, q string, m map[string]interface{}) ([]User, error) {
	query := "select * from users where " + q
	user, err := surreal.Find[User](a, query, m, []User{})
	return user, err
}

func All(a *surreal.AppRepository) ([]User, error) {

	q := "select * from users ;"
	m := map[string]interface{}{}
	users, err := surreal.Find[User](a, q, m, []User{})
	if nil != err {
		return []User{}, err
	}
	if 0 == len(users) {
		return []User{}, errors.New("no user has been created yet")
	}

	return users, nil
}

func Index(a *surreal.AppRepository, id string) (User, error) {

	query := "select * from users where id = $id limit 1"

	m := map[string]interface{}{
		"id": id,
	}
	user, err := surreal.Find[User](a, query, m, []User{})

	if 0 == len(user) {

		log.Errorf("no users with id: %v", id)
		return user[0], errors.New("check the information you entered")
	}

	return user[0], err
}

func Update(a *surreal.AppRepository, u User) error {

	query := "select * from users where id = $id limit 1 ;"
	id := u.Id
	m := map[string]interface{}{
		"id": id,
	}
	user, err := surreal.Find[User](a, query, m, []User{})

	if nil != err {
		return err
	}

	if 0 == len(user) {
		return errors.New(fmt.Sprintf("cannot find user with id : %v", id))
	}

	if "" == user[0].Id {
		return errors.New(fmt.Sprintf("cannot find user with id : %v", id))
	}

	if reflect.DeepEqual(u, user[0]) {
		return nil
	}

	// [password , email]  can only changed by thier own functions
	u.Password = user[0].Password
	u.Email = user[0].Email

	err = surreal.Update(a, id, u)

	if nil != err {
		return err
	}
	return nil
}

func Delete(a *surreal.AppRepository, id string) error {

	_, err := Index(a, id)
	if nil != err {
		return err
	}
	err = surreal.Delete(a, id)
	return err
}

// this will be used to show all the user with given
func Roles(a *surreal.AppRepository, id string) (rolemodal.Role, []User, error) {
	query := "select * from roles where id = $id limit 1"
	m := map[string]interface{}{
		"id": id,
	}
	roles, err := surreal.Find[rolemodal.Role](a, query, m, []rolemodal.Role{})

	if nil != err {
		return rolemodal.Role{}, []User{}, err
	}

	if 0 == len(roles) {

		return rolemodal.Role{}, []User{}, err
	}

	if "" == roles[0].Id {

		return rolemodal.Role{}, []User{}, err
	}

	query = "select * from users where role = $role  "
	m = map[string]interface{}{
		"role": roles[0].Id,
	}
	users, err := surreal.Find[User](a, query, m, []User{})

	if nil != err {
		return rolemodal.Role{}, []User{}, err
	}

	if 0 == len(users) {

		return rolemodal.Role{}, []User{}, err
	}

	if "" == users[0].Id {

		return rolemodal.Role{}, []User{}, err
	}
	return roles[0], users, nil
}

// TODO: add tag , remove tag
func AddTag(a *surreal.AppRepository, id, tag string) error {

	u, err := Index(a, id)
	if nil != err {
		return err
	}

	for i := 0; i < len(u.Tags); i++ {
		if u.Tags[i] == tag {
			return errors.New("tags already added")
		}
	}

	// TODO: user tag += tagid instead of this
	u.Tags = append(u.Tags, tag)

	err = Update(a, u)

	if nil != err {
		return err
	}
	return nil
}

// TODO: add and remove tag
func RemoveTag(a *surreal.AppRepository, id, tag string) error {

	user, err := Index(a, id)

	if nil != err {
		return err
	}

	t, err := tagmodal.Index(a, tag)

	if nil != err {
		return err
	}

	// extra level of checking
	if t.Id != tag {
		return errors.New("conflict in tag id")
	}

	if !slices.Contains(user.Tags, tag) {
		return errors.New("user has not tag with given id: " + tag)

	}
	// TODO: user tag  -= tagid
	deletetdTagIndex := slices.Index(user.Tags, tag)
	slices.Replace(user.Tags, deletetdTagIndex, deletetdTagIndex+1, "")

	Update(a, user)

	return nil
}
