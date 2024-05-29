package authmodal

import (
	"errors"
	"time"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// TODO: refactort this latter

type Auth struct {
	// TODO: make user id the index
	User  string `json:"user"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func Create(a *surreal.AppRepository, user, role, token string) error {

	authData := Auth{
		User:  user,
		Role:  role,
		Token: token,
	}
	_, err := a.Db.Create("auths", authData)

	if nil != err {
		return err
	}
	return nil
}

// [user] -> id of the user
func New(user, token string) *Auth {
	return &Auth{
		User:  user,
		Token: token,
	}
}

// login
func Login(a *surreal.AppRepository, email, password string) (*string, *string, *string, error) {

	user, err := surreal.Find(a,
		"select * from users where email = $email limit 1; ",
		map[string]interface{}{
			"email": email},
		[]usermodal.User{})

	if nil != err {
		return nil, nil, nil, err
	}

	if 0 == len(user) {
		return nil, nil, nil, errors.New("no such user")
	}

	role := user[0].Role
	r, err := rolemodal.HasRoleWithId(a, role)

	if nil != err {
		return nil, nil, nil, err

	}
	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))

	if nil != err {
		return nil, nil, nil, errors.New("wrong email or password")

	}
	id := user[0].Id

	// TODO: add the secrete to tomal file
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": id,

		// token is valied for a 30 days
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//TODO: the SignedString method cannot take string directly
	t, err := token.SignedString([]byte("sfsd"))

	err = Create(a, user[0].Id, user[0].Role, t)

	if nil != err {
		return nil, nil, nil, errors.New("fiald to create token in db error is :   " + err.Error())
	}

	return &t, &r.Name, &id, err
}

// singup
func SignUp(a *surreal.AppRepository, name, email, password, phone, role string) error {

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := usermodal.New(
		name,
		email,
		phone,
		string(hash),
		role,
	)

	err := user.Create(a)

	if nil != err {
		return err
	}

	return nil
}

// TODO: implement signout method
// signout
