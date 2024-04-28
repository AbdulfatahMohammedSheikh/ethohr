package usermodaltest

import (
	"math/rand"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func TestCreate(t *testing.T) {
	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	roles := rolemodal.All(repo)

	for _, role := range roles {
		phone := faker.Phonenumber()

		user := usermodal.New(

			faker.Name(),
			faker.Email(),
			phone,
			phone,
			role.Id,
		)
		err = user.Create(repo)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
}

func TestGetAllUsers(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	users, err := usermodal.All(repo)
	if nil != err {
		t.Fatal(err.Error())
	}

	if len(users) == 0 {
		t.Fatal("no users in db yet")
	}
}

func TestIndex(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	users, err := usermodal.All(repo)
	if nil != err {
		t.Fatal(err.Error())
	}

	if len(users) == 0 {
		t.Fatal("no users in db yet")
	}

	index, err := usermodal.Index(repo, users[0].Id)

	if nil != err {
		t.Fatal(err.Error())
	}

	if index.Id != users[0].Id {
		t.Fatal("error while compareing user and user index quary")
	}
}

func TestUpdateName(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	phone := faker.Phonenumber()
	role := rolemodal.All(repo)
	name := "will be updated"

	user := usermodal.New(
		name,
		faker.Email(),
		phone,
		phone,
		role[0].Id,
	)
	err = user.Create(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	q := "select * from users where name = $name limit 1;"
	m := map[string]interface{}{
		"name": name,
	}
	users, err := usermodal.Find(repo, q, m)

	if nil != err {
		t.Fatal(err.Error())
	}

	if users[0].Name == user.Name {
		t.Fatal("update did not happen")
	}
}

func TestDeleteUser(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	phone := faker.Phonenumber()
	role := rolemodal.All(repo)
	name := "will be deleted"

	user := usermodal.New(
		name,
		faker.Email(),
		phone,
		phone,
		role[0].Id,
	)
	err = user.Create(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	q := "select * from users where name = $name limit 1;"
	m := map[string]interface{}{
		"name": name,
	}
	users, err := usermodal.Find(repo, q, m)

	if nil != err {
		t.Fatal(err.Error())
	}

	err = usermodal.Delete(repo, users[0].Id)

	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestAddTag(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	tags, err := tagmodal.All(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	users, err := usermodal.All(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	target := rand.Intn(len(users))

	err = usermodal.AddTag(repo, users[target].Id, tags[0].Id)

	if nil != err {
		t.Fatal(err.Error())
	}
}

// TODO: make this test work

// func TestRemoveTag(t *testing.T) {
//
// 	repo, err := testrunner.GetConfig()
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
//
// 	defer repo.Close()
//
// 	phone := faker.Phonenumber()
// 	role := rolemodal.All(repo)
// 	name := "will have a tag in it"
//
// 	user := usermodal.New(
// 		name,
// 		faker.Email(),
// 		phone,
// 		phone,
// 		role[0].Id,
// 	)
// 	err = user.Create(repo)
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
//
// 	q := "select * from users where name = $name limit 1;"
// 	m := map[string]interface{}{
// 		"name": name,
// 	}
// 	users, err := usermodal.Find(repo, q, m)
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
//
// 	tags, err := tagmodal.All(repo)
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
// 	tag := tags[0].Id
//
// 	err = usermodal.AddTag(repo, users[0].Id, tag)
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
//
// 	err = usermodal.RemoveTag(repo, users[0].Id, tag)
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
// }
