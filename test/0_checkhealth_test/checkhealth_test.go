package checkhealthtest

import (
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	emplpyermigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/emplpyerMigration"
	rolemigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/roleMigration"
	tagmigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/tagMigration"
	usermigration "github.com/AbdulfatahMohammedSheikh/backend/migrations/userMigration"
)

func TestCheckHealth(t *testing.T) {
	testCase := testrunner.NewTestCase("cheack health test", "/health", testrunner.GET, "", 200)
	testrunner.HttpRunner(t, testCase)
}

func TestSetup(t *testing.T) {

	repo, err := testrunner.GetConfig()
	defer repo.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	rolemigration.SetUp(repo)
	usermigration.SetUp(repo)
	tagmigration.SetUp(repo)
	emplpyermigration.SetUp(repo)
}

func TestInit(t *testing.T) {

	repo, err := testrunner.GetConfig()
	defer repo.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	rolemigration.Down(repo)
	tagmigration.Down(repo)
	usermigration.Down(repo)
	emplpyermigration.Down(repo)
}

