package offermigration

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	offermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/offerModal"
	"github.com/bxcodec/faker/v3"
)

func SetUp(a *surreal.AppRepository) {

	// get all employers

	employers, err := employermodal.All(a)

	if nil != err {
		panic(err.Error())
	}

	if len(employers) == 0 {
		panic("no employers yet")
	}

	employer := employers[0]

	for i := 0; i < 4; i++ {

		requirements := []string{
			faker.Word(),
			faker.Word(),
			faker.Word(),
		}

		duty := []string{
			faker.Word(),
			faker.Word(),
			faker.Word(),
		}
		offer := offermodal.New(employer.Id, employer.Name, faker.Word(), "*****", "******", requirements, duty)

		err = offer.Create(a)

		if nil != err {
			panic(err.Error())
		}

	}

}

func Down(a *surreal.AppRepository) {

	_, err := a.Db.Delete("Offers")

	if nil != err {
		panic(err)
	}

}
