package emplpyermigration

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func SetUp(a *surreal.AppRepository) {

	counter := 0
	// user role index is 2
	roles := rolemodal.All(a)

	var role string
	for _, v := range roles {
		if v.Name == "ngo" {
			role = v.Id

		}
	}

	user, err := usermodal.Find(a, " role = $role", map[string]interface{}{
		"role": role,
	})

	if nil != err {
		panic(err.Error())
	}

	for counter < 4 {

		e := employermodal.New(
			user[0].Id,
			faker.Name(),
			faker.Word(),
			faker.Paragraph(),
			faker.Name(),
			faker.Phonenumber(),
		)

		err := e.Create(a)

		if nil != err {
			panic(err.Error())
		}

		counter++
	}
}

func Down(a *surreal.AppRepository) {

	_, err := a.Db.Delete("employers")

	if nil != err {
		panic(err)
	}

}
