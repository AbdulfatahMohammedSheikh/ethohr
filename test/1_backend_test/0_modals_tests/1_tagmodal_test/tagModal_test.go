package tagmodaltest

import (
	"fmt"
	"math/rand"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	"github.com/bxcodec/faker/v3"
)

func TestCreateTag(t *testing.T) {

	tags := []string{"gaming", "business", "housing", "IT"}
	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()
	limit := rand.Intn(len(tags) - 1)

	for i := 0; i <= limit; i++ {

		err = tagmodal.Create(repo, tags[i])

		if nil != err {
			t.Fatal(err.Error())
		}

	}
}

func TestAllTags(t *testing.T) {
	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	tags, err := tagmodal.All(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(tags) {
		t.Fatal("there is no tags created yet")
	}
}

func TestIndex(t *testing.T) {
	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	tags, err := tagmodal.All(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(tags) {
		t.Fatal("there is no tags created yet")
	}

	t.Run("Find tag successfully by id", func(t *testing.T) {
		target := rand.Intn(len(tags) - 1)
		id := tags[target].Id
		tag, err := tagmodal.Index(repo, id)

		if nil != err {
			t.Fatal(err.Error())
		}

		if tag.Id != id {
			message := fmt.Sprintf("quary id : %s , result id : %s", id, tag.Id)
			t.Fatal(message)
		}
	})

	t.Run("Fail to get tag by id", func(t *testing.T) {
		tag, err := tagmodal.Index(repo, "id")

		if nil == err {
			t.Fatal("test should have fail due to wrong id")
		}

		if tag.Id != "" {
			t.Fatal("tag shoud have no value in it due to wrong id")
		}
	})
}

func TestUpdateTagName(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	tags, err := tagmodal.All(repo)
	target := rand.Intn(len(tags) - 1)

	NewName := faker.Name() + " tag"
	newTag := tagmodal.Tag{
		Id:   tags[target].Id,
		Name: NewName,
	}

	err = tagmodal.Update(repo, newTag)

	if nil != err {
		t.Fatal(err.Error())
	}

	updatedTag, err := tagmodal.Index(repo, newTag.Id)

	if nil != err {
		t.Fatal(err.Error())
	}

	if updatedTag.Name != newTag.Name {
		t.Fatal("fail to updated tags")
	}
}

func TestDeleteTag(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	tags, err := tagmodal.All(repo)
	if nil != err {
		t.Fatal(err.Error())
	}
	id := tags[len(tags)-1].Id

	err = tagmodal.Delete(repo, id)

	if nil != err {
		t.Fatal(err.Error())
	}

}
