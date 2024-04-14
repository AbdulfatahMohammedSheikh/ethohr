package employermodaltest

import (
	"math/rand"
	"reflect"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func TestCreate(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	// get roles
	roles := rolemodal.All(r)
	var employerRoleId string

	for _, role := range roles {
		if role.Name == "ngo" {
			employerRoleId = role.Id
			break
		}
	}

	if "" == employerRoleId {
		t.Fatal("no roles for ngo")
	}

	emplyers, err := usermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, emplyer := range emplyers {
		if employerRoleId == emplyer.Role {

			e := employermodal.New(
				emplyer.Id,
				faker.Name()+" ngo",
				faker.Word(),
				faker.Paragraph(),
				faker.Name()+" city",
				faker.Phonenumber(),
			)
			err = e.Create(r)

			if nil != err {
				t.Fatal(err.Error())
			}
		}
	}

}

func TestIndex(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if len(employers) == 0 {
		t.Fatal("no employers yet")
	}

	for _, employer := range employers {

		e, err := employermodal.Index(r, employer.Id)

		if nil != err {
			t.Fatal(err.Error())
		}

		if !reflect.DeepEqual(e, employer) {
			t.Fatal("error employer ids conflict")
		}

	}
}

func TestAll(t *testing.T) {
	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}
	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if len(employers) == 0 {
		t.Fatal("no employers yet")
	}

}

func TestUpdate(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if len(employers) == 0 {
		t.Fatal("no employers yet")
	}

	for _, employer := range employers {

		employer.Phone = faker.Phonenumber()
		employer.Meto = employer.Meto + " " + faker.Word()
		err = employermodal.Update(r, employer)

		if nil != err {
			t.Fatal(err.Error())
		}
	}
}

func TestAddTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if len(employers) == 0 {
		t.Fatal("no employers yet")
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, employer := range employers {
		selectedTag := rand.Intn(len(tags) - 1)
		err = employermodal.AddTag(r, employer.Id, tags[selectedTag].Id)
		if nil != err {
			t.Fatal(err.Error())
		}
	}

}

func TestRemoveTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if len(employers) == 0 {
		t.Fatal("no employers yet")
	}

	for _, employer := range employers {
		if 0 < len(employer.Tags) {

			err = employermodal.RemoveTag(r, employer.Id, employer.Tags[0])

			if nil != err {
				t.Fatal(err.Error())
			}
			break
		}

	}
}
func TestDeleteEmployer(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(employers) {
		t.Fatal("there is no employer in db yet")
	}

	for _, employer := range employers {
		err = employermodal.Delete(r, employer.Id)
		if nil != err {
			t.Fatal(err.Error())
		}
	}

}
