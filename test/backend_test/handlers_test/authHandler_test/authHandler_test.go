package authhandlertest

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

var createdUsers []string

func TestSignupHandler(t *testing.T) {
	// get list of all roles
	r, err := testrunner.GetConfig()

    defer r.Close()
	if nil != err {
		t.Fatal(err.Error())
	}

	roles := rolemodal.All(r)

	var testCase testrunner.TestCase

	for _, role := range roles {
		// sime test
		// TODO: store the user and passoerd in /tmp/
		user := usermodal.New(
			faker.Name(),
			faker.Email(),
			faker.Phonenumber(),
			faker.Password(),
			role.Id,
		)

		pattern := fmt.Sprintf("%s/%s", user.Email, user.Password)
		createdUsers = append(createdUsers, pattern)

		testCase = testrunner.TestCase{
			Name:   fmt.Sprintf("createing user with role %v ", role.Name),
			Input:  user.Encode(),
			Method: testrunner.POST,
			Url:    "/auth/signup",
			Output: 200,
		}
		_ = role

		testrunner.HttpRunner(t, testCase)
	}
}

func TestLoginHander(t *testing.T) {

	if 0 == len(createdUsers) {
		t.Fatal("now user was created yeat")
	}

	for tt := range createdUsers {
		result := strings.Split(createdUsers[tt], "/")

		if 0 == len(result) {
			t.Fatal("failed to get name or password")
		}

		newCase := testrunner.NewTestCase(
			fmt.Sprintf("login as user %s", result[0]),
			"/auth/login",
			testrunner.POST,
			url.Values{
				"email":    {result[0]},
				"password": {result[1]},
			}.Encode(),
			200,
		)
		testrunner.HttpRunner(t, newCase)
	}

	newCase := testrunner.NewTestCase(
		"must fail login",
		"/auth/login",
		testrunner.POST,
		url.Values{
			"email":    {"newemail"},
			"password": {"the passwordj js;f "},
		}.Encode(),

		401,
	)
	testrunner.HttpRunner(t, newCase)
}
