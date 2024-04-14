package taghandertest

import (
	"net/url"
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	"github.com/bxcodec/faker/v3"
)

func TestCreateTag(t *testing.T) {

	testCases := []testrunner.TestCase{
		testrunner.NewTestCase("Create new tag", "/tag", testrunner.POST, tagmodal.MapIt("", faker.Name()), 200),
		testrunner.NewTestCase("failed to create tag due to empty input", "/tag", testrunner.POST, "", 401),
	}

	for _, testCase := range testCases {
		testrunner.HttpRunner(t, testCase)
	}
}

func TestAllTeag(t *testing.T) {

	testCase := testrunner.NewTestCase("get all tag", "/tags", testrunner.GET, "", 200)

	testrunner.HttpRunner(t, testCase)
}

func TestIndex(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, tag := range tags {

		successTest := testrunner.NewTestCase("get tag with id : "+tag.Id, "/tag?id="+tag.Id, testrunner.GET, "", 200)
		failedTest := testrunner.NewTestCase("failed to get tag due to using tag  name instade of id", "/tag?id="+tag.Name, testrunner.GET, "", 401)
		testrunner.HttpRunner(t, successTest)
		testrunner.HttpRunner(t, failedTest)
	}
}

func TestUpdateTags(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, tag := range tags {

		successTest := testrunner.NewTestCase("update tag with id : "+tag.Id, "/tag", testrunner.PATCH, tagmodal.MapIt(
			tag.Id,
			faker.Name(),
		), 200)
		failedTest := testrunner.NewTestCase("failed to update tag due to using tag  name instade of id", "/tag", testrunner.PATCH,
			"", 401)
		testrunner.HttpRunner(t, successTest)
		testrunner.HttpRunner(t, failedTest)
	}
}

func TestUDeleteTag(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}

	tags, err := tagmodal.All(r)

	if nil != err {
		t.Fatal(err.Error())
	}

	for _, tag := range tags {

		successTest := testrunner.NewTestCase("Delete tag with id : "+tag.Id, "/delete/tag", testrunner.POST,
			url.Values{
				"id": {tag.Id},
			}.Encode(),
			200)
		testrunner.HttpRunner(t, successTest)

	}
}
