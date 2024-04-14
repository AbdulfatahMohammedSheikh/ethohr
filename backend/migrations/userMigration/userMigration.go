package usermigration

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	authmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/authModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func SetUp(a *surreal.AppRepository) {

	roles := rolemodal.All(a)

	for _, role := range roles {

		// NOTE: the password and the number are the same to make things easy later
		number := faker.Phonenumber()
		u := usermodal.New(
			faker.Name(),
			faker.Email(),
			number,
			number,
			role.Id,
		)

		err := authmodal.SignUp(
			a,
			u.Name,
			u.Email,
			u.Password,
			u.Phone,
			role.Id,
		)
		if nil != err {
			panic(err.Error())
		}

	}

}


func Down(a *surreal.AppRepository) {

	_, err := a.Db.Delete("users")

	if nil != err {
		panic(err)
	}

}
