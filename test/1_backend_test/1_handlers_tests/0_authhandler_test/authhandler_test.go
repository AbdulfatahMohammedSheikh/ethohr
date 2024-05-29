package authhandlertest

import (
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
)

// create new user
func Test0Signup(t *testing.T) {

	repo, err := testrunner.GetConfig()

	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	// testCases := []testrunner.TestCase{
	// 	testrunner.NewTestCase(
	//            "singup new user",
	//        ),
	// }

}

// login
