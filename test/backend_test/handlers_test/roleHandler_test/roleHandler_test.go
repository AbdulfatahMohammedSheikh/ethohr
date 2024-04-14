package rolehandlertest

import (
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
)

func TestGetAllRoles(t *testing.T) {
	testCase := testrunner.NewTestCase("Get all roles", "/roles", testrunner.GET, "", 200)
	testrunner.HttpRunner(t, testCase)
}

func TestIndex(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}
	// get all users
	users, err := usermodal.All(r)
	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(users) {
		t.Fatal("no users in db yet")
	}

	id := users[0].Role

	// take the first user role
	testCase := testrunner.NewTestCase("Get all users with role id"+id, "/role?id="+id, testrunner.GET, "", 200)
	testrunner.HttpRunner(t, testCase)

	// should fail due to incorrect role id
	testCase = testrunner.NewTestCase("fail due to incorrect role id"+users[0].Id, "/role?id="+users[0].Id, testrunner.GET, "", 401)
	testrunner.HttpRunner(t, testCase)
}
