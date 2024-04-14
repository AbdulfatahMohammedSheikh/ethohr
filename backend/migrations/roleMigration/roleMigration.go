package rolemigration

import (
	logger "github.com/sirupsen/logrus"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
)

func SetUp(a *surreal.AppRepository) {
	var log = logger.New()
	roles := []string{"admin", "user", "ngo"}

	for _, r := range roles {

		role := rolemodal.New(r)
		err := role.Create(a)

		if nil != err {
			log.Error(err.Error())
		}
	}
}

func Down(a *surreal.AppRepository) {

	_, err := a.Db.Delete("roles")

	if nil != err {
		panic(err)
	}

}
