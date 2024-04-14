package usermodaltest

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

// NOTE: no need to test createUser due to the use of the same functionaliy in authmodal Singup

func TestIndex(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	q := "select * from users ;"
	m := map[string]interface{}{}

	users, err := surreal.Find[usermodal.User](r, q, m, []usermodal.User{})

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, u := range users {

		user, err := usermodal.Index(r, u.Id)

		if nil != err {
			t.Fatal(err.Error())
		}

		if reflect.DeepEqual(user, u) == false {

			t.Fatal(fmt.Sprintf("%v not qual to %v", u, user))
		}
	}
}

func TestUpdateUserEmail(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	q := "select * from users ;"
	m := map[string]interface{}{}

	users, err := surreal.Find[usermodal.User](r, q, m, []usermodal.User{})

	if nil != err {
		t.Fatal(err.Error())
	}

	for i := 0; i < len(users); i++ {
		orderEmail := users[i].Email
		users[i].Email = faker.Email()

		if orderEmail == users[i].Email {
			t.Fatal("no update counld  happen due to email never changing")
		}
		err = usermodal.Update(r, users[i])

		if nil != err {
			t.Fatal(err.Error())
		}
	}
}

func TestFindUsersByRole(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	roles := rolemodal.All(r)

	if 0 == len(roles) {
		t.Fatal("No roles in db")
	}

	for _, role := range roles {
		roleobj, _, err := usermodal.Roles(r, role.Id)

		if nil != err {
			t.Fatal(err.Error())
		}

		if roleobj != role {

			t.Fatal("roles object is not equal with role")
		}
	}

}

// TODO: implement
// add tag

func TestAddTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	roles := rolemodal.All(r)

	adminRoleId := ""
	// TODO: remove this _
	_ = adminRoleId
	for _, roleobj := range roles {
		if roleobj.Name == "admin" {
			adminRoleId = roleobj.Id
			break
		}
		fmt.Println(roleobj)
	}

	users, err := usermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, user := range users {

		if user.Role == adminRoleId {
			continue
		}

		tagIndex := rand.Intn(len(tags) - 1)

		err = usermodal.AddTag(r, user.Id, tags[tagIndex].Id)

		if nil != err {
			t.Fatal(err.Error())
		}
	}

}

// removes the first tag for each user
func TestRemoveTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	users, err := usermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, user := range users {
		if 0 == len(user.Tags) {
			continue
		}

		tag, err := tagmodal.Index(r, user.Tags[0])

		if nil != err {
			t.Fatal(err.Error())
		}
		err = usermodal.RemoveTag(r, user.Id, tag.Id)

		if nil != err {
			t.Fatal(err.Error())
		}
	}

}

func TestDeleteUsers(t *testing.T) {
	// get all roles and avoid deleting the one wiht admin

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	roles := rolemodal.All(r)

	adminRoleId := ""
	for _, roleobj := range roles {
		if roleobj.Name == "admin" {
			adminRoleId = roleobj.Id
			break
		}
		fmt.Println(roleobj)
	}

	if "" == adminRoleId {
		t.Fatal("did not find the admin role id ")
	}

	users, err := usermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, u := range users {
		if u.Role != adminRoleId {
			usermodal.Delete(r, u.Id)
		}
	}
}
