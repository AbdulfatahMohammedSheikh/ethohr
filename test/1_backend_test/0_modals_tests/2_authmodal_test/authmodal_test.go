package authmodaltest

import (
	"math/rand"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	authmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/authModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func Test0Singup(t *testing.T) {
	repo, err := testrunner.GetConfig()
	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	roles := rolemodal.All(repo)
	if 0 == len(roles) {
		t.Fatal("no roles was created yet")
	}

	for _, role := range roles {

		// phone number will be used as a password as will
		phone := faker.Phonenumber()

		err = authmodal.SignUp(
			repo,
			faker.Name(),
			faker.Email(),
			phone,
			phone,
			role.Id,
		)

		if nil != err {
			t.Fatal(err.Error())
		}
	}

}

func Test01Login(t *testing.T) {
	repo, err := testrunner.GetConfig()
	if nil != err {
		t.Fatal(err.Error())
	}

	defer repo.Close()

	t.Run("successfully login", func(t *testing.T) {
		users, err := usermodal.All(repo)
		if nil != err {
			t.Fatal(err.Error())
		}
		target := rand.Intn(len(users) - 1)

		user := users[target]
        // FIXME: remove the _ = user and make the auth maog work 
        _ = user

		// here we are using the phone number as the password
		// _, err = authmodal.Login(repo, user.Email, user.Phone)
		//
		// if nil != err {
		// 	t.Fatal(err.Error())
		// }

	})

	t.Run("fail to login", func(t *testing.T) {
        // FIXME: make this test work

		// _, err = authmodal.Login(repo, "email", "password")
		//
		// if nil == err {
		// 	t.Fatal("test was supposed to fail due to wrong email and password")
		// }

	})

}
