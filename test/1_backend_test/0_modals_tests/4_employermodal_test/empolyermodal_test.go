package employermodaltest

import (
	"fmt"
	"math/rand"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

var id string
var addeddTag string

func TestCreateEmployer(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	roles := rolemodal.All(repo)
	var role string

	for _, r := range roles {
		if r.Name == "ngo" {
			role = r.Id
			break
		}
	}

	r, users, err := usermodal.Roles(repo, role)

	if nil != err {
		t.Fatal(err.Error())
	}

	if r.Id != role {
		message := fmt.Sprintf("role id : %v , but get : %v ", role, r.Id)
		t.Fatal(message)
	}

	target := rand.Intn(len(users) - 1)
	selectedUser := users[target]
	employer := employermodal.New(
		selectedUser.Id,
		faker.Name(),
		faker.Word(),
		faker.Paragraph(),
		"",
		faker.Phonenumber(),
	)

	err = employer.Create(repo)

	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestAllEmployers(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	emlpoyers, err := employermodal.All(repo)
	target := emlpoyers[0].Id
	id = target

	e, err := employermodal.Index(repo, target)

	if nil != err {
		t.Fatal(err.Error())
	}

	if e.Id != target {
		t.Fatal("error which comparing target with index quary ")
	}

}
func TestUpdate(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()
	e, err := employermodal.Index(repo, id)

	if nil != err {
		t.Fatal(err.Error())
	}
	e.Name = "updated name"

	err = employermodal.Update(repo, e)

	if nil != err {
		t.Fatal(err.Error())
	}
}

// TODO: added tests  for added and removing tags

func TestDelete(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()
	err = employermodal.Delete(repo, id)

	if nil != err {
		t.Fatal(err.Error())
	}
}
