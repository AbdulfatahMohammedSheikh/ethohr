package tagmigration

import (
	"math/rand"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
)

func SetUp(a *surreal.AppRepository) {
	tags := []string{"tech", "spart", "health", "education", "nature", "hr", "enginering", "oil"}

	for i := range tags {
		// check if the tag name alredy in db
		err := tagmodal.HasTagWithName(a, tags[i])
		if nil != err {
			continue
		} else {

			err = tagmodal.Create(a, tags[i])
			if nil != err {
				panic(err.Error())
			}
		}
	}

	// add same tags to the users

	users, err := usermodal.All(a)
	if nil != err {
		panic(err.Error())
	}
	allTags, err := tagmodal.All(a)
	if nil != err {
		panic(err.Error())
	}

	for _, user := range users {
		firstTag := rand.Intn(len(allTags) - 1)
		seoncdTag := rand.Intn(len(allTags) - 1)
		usermodal.AddTag(a, user.Id, allTags[firstTag].Id)
		usermodal.AddTag(a, user.Id, allTags[seoncdTag].Id)

	}
}

func Down(a *surreal.AppRepository) {

	_, err := a.Db.Delete("tags")

	if nil != err {
		panic(err)
	}

}
