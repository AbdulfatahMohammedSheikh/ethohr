package authmodaltest

import (
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	authmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/authModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func TestSignup(t *testing.T) {
	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal("faited to run Singup Test due to: ", err.Error())
	}

	roles := rolemodal.All(r)
	for _, role := range roles {

		// using the phone number to be also the password
		phone := faker.Phonenumber()
		err = authmodal.SignUp(r,
			faker.NAME,
			faker.Email(),
			phone,
			phone,
			role.Id,
		)
		if nil != err {
			t.Fatal(err.Error())
		}
	}

	phone := faker.Phonenumber()
	err = authmodal.SignUp(r,
		faker.NAME,
		faker.Email(),
		phone,
		phone,
		faker.Email(), // fake role
	)

	if nil == err {
		t.Fatal("should faill due to wrong role id")
	}

}

func TestLogin(t *testing.T) {
	// TODO: implemetn login

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal("faited to run login  Test due to: ", err.Error())
	}

	users, err := usermodal.All(r)
	_ = users
	if nil != err {
		t.Fatal(err.Error())
	}

	for _, user := range users {

		_, err = authmodal.Login(r, user.Email, user.Phone)

		if nil != err {
			t.Fatal("error while login : " + err.Error())
		}

	}
}
