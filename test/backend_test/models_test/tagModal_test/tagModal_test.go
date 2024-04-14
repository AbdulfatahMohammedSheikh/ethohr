package tagmodaltest

import (
	"fmt"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	"github.com/bxcodec/faker/v3"
)

func TestCreateTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	err = tagmodal.Create(r, faker.Name())

	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestGetAllTeags(t *testing.T) {
	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}
	_, err = tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestGetTagByIndex(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, tag := range tags {
		result, err := tagmodal.Index(r, tag.Id)
		if nil != err {
			t.Fatal(err.Error())
		}

		if result.Id != tag.Id {
			t.Fatal(fmt.Sprintf("tag id is : %v but get id : %v", tag.Id, result.Id))
		}
	}
}

func TestUpdateTagName(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for i := 0; i < len(tags); i++ {

			tags[i].Name = faker.Name()
			tagmodal.Update(r, tags[i])
	}
}

func TestDeleteTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for i := 0; i < len(tags); i++ {

		if tags[i].Name != "admin" && tags[i].Name != "ngo" && tags[i].Name != "user" {
			err = tagmodal.Delete(r, tags[i].Id)

			if nil != err {
				t.Fatal(err.Error())
			}

		}
	}

}
