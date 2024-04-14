package main

import (
	surreal "github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	emplpyermigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/emplpyerMigration"
	rolemigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/roleMigration"
	tagmigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/tagMigration"
	usermigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/userMigration"
	logger "github.com/sirupsen/logrus"
)

var log = logger.New()

func main() {


	config := surreal.NewApp()
	log.Trace(logger.TraceLevel)

	repo, err := surreal.NewAppRepository(config.DB)

	if nil != err {
		log.Fatalf("failed to creat app : %v", err)
	}

	log.Info("connecting to database ")
	defer func() {
		log.Info("closing connection with  database")
		repo.Close()
	}()

	rolemigration.SetUp(repo)
	usermigration.SetUp(repo)
	tagmigration.SetUp(repo)
	emplpyermigration.SetUp(repo)

}
