package employerhandlertest

import (
	"math/rand"
	"net/url"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/bxcodec/faker/v3"
)

func TestCreate(t *testing.T) {
	// creater the user

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}
	// get all users with role employer

	roles := rolemodal.All(r)

	if 0 == len(roles) {
		t.Fatal("no tags yet")
	}

	var employerRoleId string

	for _, role := range roles {
		if "ngo" == role.Name {
			employerRoleId = role.Id
			break
		}
	}
	if "" == employerRoleId {
		t.Fatal("could not find ngo role id")
	}

	// get all users with role employer
	_, users, err := usermodal.Roles(r, employerRoleId)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(users) {
		t.Fatal("no users with role: " + employerRoleId)
	}

	for _, employer := range users {
		testCase := testrunner.NewTestCase(
			"create empolyer for user: "+employer.Id,
			"/employer",
			testrunner.POST,
			url.Values{
				"name":     {faker.Name()},
				"user_id":  {employer.Id},
				"meto":     {faker.Word()},
				"about":    {faker.Paragraph()},
				"location": {faker.Word()},
				"phone":    {faker.Phonenumber()},
			}.Encode(),
			200,
		)
		testrunner.HttpRunner(t, testCase)
	}

	testCase := testrunner.NewTestCase(
		"faild to create new employer due to missing id input",
		"/employer",
		testrunner.POST,
		url.Values{
			"name":     {faker.Name()},
			"meto":     {faker.Word()},
			"about":    {faker.Paragraph()},
			"location": {faker.Word()},
			"phone":    {faker.Phonenumber()},
		}.Encode(),
		404,
	)
	testrunner.HttpRunner(t, testCase)
}

func TestGetEmployerById(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}
	if 0 == len(employers) {
		t.Fatal("no employers")
	}
	id := employers[0].Id
	testCase := testrunner.NewTestCase("get employer by id : "+id, "/employer?id="+id, testrunner.GET,
		"",
		200,
	)
	testrunner.HttpRunner(t, testCase)
}
func TestGetEmployers(t *testing.T) {
	testCase := testrunner.NewTestCase("getting all employers", "/employers", testrunner.GET,
		"",
		200,
	)
	testrunner.HttpRunner(t, testCase)
}

func TestUpdate(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(employers) {
		t.Fatal("no employers")
	}

	index := rand.Intn(len(employers) - 1)
	employer := employers[index]
	employer.Phone = faker.Phonenumber()
	testCase := testrunner.NewTestCase("get employer by id : "+employer.Id, "/employer", testrunner.PATCH,
		url.Values{
			"id":       {employer.Id},
			"user_id":  {employer.UserId},
			"name":     {faker.Name() + " updated " + employer.Name},
			"meto":     {faker.Word()},
			"about":    {faker.Paragraph()},
			"location": {faker.Word()},
			"phone":    {faker.Phonenumber()},
		}.Encode(),
		200,
	)
	testrunner.HttpRunner(t, testCase)

}

func TestAddTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(tags) {
		t.Fatal("no tags yet")
	}

	tag := tags[0]

	employers, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(employers) {
		t.Fatal("no employers yet")
	}

	testCase := testrunner.NewTestCase("update add tag to employer :"+employers[0].Id, "/employer/tag", testrunner.POST,
		url.Values{
			"id":  {employers[0].Id},
			"tag": {tag.Id},
		}.Encode(),
		200,
	)
	testrunner.HttpRunner(t, testCase)

}

// TODO: make this work

// func TestDeleteTag(t *testing.T) {
//
// 	r, err := testrunner.GetConfig()
// 	defer r.Close()
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
//
// 	employers, err := employermodal.All(r)
//
// 	if nil != err {
// 		t.Fatal(err.Error())
// 	}
//
// 	if 0 == len(employers) {
// 		t.Fatal("no emploters yet")
// 	}
//
// 	id := employers[0].Id
// 	// should test id tag is ""
// 	tag := employers[0].Tags[0]
//
// 	testCase := testrunner.NewTestCase("update add tag to employer :"+employers[0].Id, "/employer/delete/tag", testrunner.POST,
// 		url.Values{
// 			"id":  {id},
// 			"tag": {tag},
// 		}.Encode(),
// 		200,
// 	)
// 	testrunner.HttpRunner(t, testCase)
//
// }

func TestDeleteEmployer(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	e, err := employermodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	if 0 == len(e) {
		t.Fatal("no employers in db")
	}

    testCase := testrunner.NewTestCase("deleting employer : "+e[0].Id, "/delete/employer", testrunner.POST,
		url.Values{
			"id":  {e[0].Id},
		}.Encode(),
		200,
	)
	testrunner.HttpRunner(t, testCase)

}
