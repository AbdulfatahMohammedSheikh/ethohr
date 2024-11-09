package offermodaltest

import (
	// "reflect"
	"reflect"
	"testing"

	"math/rand"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	offermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/offerModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	"github.com/bxcodec/faker/v3"
)

// helper function for creating  tags
func tagGenerator(a *surreal.AppRepository) (*tagmodal.Tag, error) {

	tags, err := tagmodal.All(a)

	if nil != err {
		return nil, err
	}
	// target := rand.Intn(len(*tags) - 1)

	target := rand.Intn(len(tags))
	tag := tags[target]

	return &tag, nil
}

// TODO: add all method
func TestCreateOffer(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	emplyer, err := employermodal.All(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	offer := offermodal.New(
		emplyer[0].Id,
		emplyer[0].Name,
		faker.Sentence(),
		faker.Paragraph(),
		faker.Sentence(),
		// TODO:  change the data
		[]string{
			"have good language",
		},
		[]string{
			"work hard",
		},
	)

	err = offer.Create(repo)

	if nil != err {
		t.Fatal(err)
	}

}

func TestAll(t *testing.T) {
	// get all the offers

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	offers, err := offermodal.All(repo)

	if nil != err {
		t.Fatal(err)
	}

	if len(*offers) == 0 {
		t.Fatal("no offers yet")
	}

}

func TestIndex(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	offers, err := offermodal.All(repo)

	_ = offers
	if nil != err {
		t.Fatal(err)
	}

	// iterate to see the first offer

	t.Run("get the first offer in the list", func(t *testing.T) {
		offer := (*offers)[0]

		index, err := offermodal.Index(repo, offer.Id)

		if nil != err {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(*index, offer) {
			t.Fatalf("get: %v  \n, expeced : %v", index, (*offers)[0])
		}

	})

}

func TestShowOffer(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	offers, err := offermodal.All(repo)

	if nil != err {
		t.Fatal(err)
	}

	if len(*offers) == 0 {
		t.Fatal("no offers yet")
	}

	_, err = offermodal.ShowOffer(repo, (*offers)[0].Id)

	if nil != err {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	offers, err := offermodal.All(repo)

	if nil != err {
		t.Fatal(err)
	}

	target := rand.Intn(len(*offers) - 1)

	offer := (*offers)[target]

	t.Run("will not update due to non changing data ", func(t *testing.T) {
		err := offermodal.Update(repo, offer)
		if nil != err {
			t.Fatal(err)
		}
	})

	t.Run("will update offer", func(t *testing.T) {
		offer.EmployerName += "updated"
		err := offermodal.Update(repo, offer)
		if nil != err {
			t.Fatal(err)
		}
	})

	t.Run("should fail due to wrong id", func(t *testing.T) {
		offer.EmployerName += "updated"
		offer.Id += "?"
		err := offermodal.Update(repo, offer)
		if nil == err {
			t.Fatal("expeted test to fail but it did not")
		}
	})
}

func TestDelete(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	emplyer, err := employermodal.All(repo)

	if nil != err {
		t.Fatal(err.Error())
	}

	offer := offermodal.New(
		emplyer[0].Id,
		emplyer[0].Name,
		faker.Sentence(),
		faker.Paragraph(),
		faker.Sentence(),
		// TODO:  change the data
		[]string{
			"have good language",
		},
		[]string{
			"work hard",
		},
	)

	// this may not work
	offer.Id = "Offers:delete"
	offer.Create(repo)

	if nil != err {
		t.Fatal(err)
	}

	_, err = repo.Db.Delete("Offers:delete")

	if nil != err {
		t.Fatal(err)
	}
}

func TestAddTag(t *testing.T) {
	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	offers, err := offermodal.All(repo)

	if nil != err {
		t.Fatal(err)
	}

	if len(*offers) == 0 {
		t.Fatal("no offers yet")
	}

	target := rand.Intn(len(*offers) - 1)

	tag, err := tagGenerator(repo)

	if nil != err {
		t.Fatal(err)
	}

	err = offermodal.AddTag(repo, (*offers)[target].Id, (*tag).Id)

	if nil != err {
		t.Fatal(err)
	}
}

// TODO: implemet this
func TestRemoveTag(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	// add the tag
	offers, err := offermodal.All(repo)

	if nil != err {
		t.Fatal(err)
	}

	if len(*offers) == 0 {
		t.Fatal("no offers yet")
	}

	target := rand.Intn(len(*offers) - 1)

	tag, err := tagGenerator(repo)

	if nil != err {
		t.Fatal(err)
	}
	offerId := (*offers)[target].Id

	err = offermodal.AddTag(repo, offerId, (*tag).Id)

	if nil != err {
		t.Fatal(err)
	}

	// removing the tag

	err = offermodal.RemoveTag(repo, (*offers)[target].Id, (*tag).Id)

	if nil != err {
		t.Fatal(err)
	}
}
