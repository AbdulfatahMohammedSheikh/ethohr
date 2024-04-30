package testrunner

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/AbdulfatahMohammedSheikh/backend/core/router"
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

var log = logger.New()

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	UPDATE = "UPDATE"
	PATCH  = "PATCH"
)

type TestCase struct {
	Name   string
	Input  string
	Output int
	Url    string
	Method string
}

func NewTestCase(name, url, Method, input string, output int) TestCase {

	return TestCase{
		Name:   name,
		Input:  input,
		Output: output,
		Url:    url,
		Method: Method,
	}
}

func GetConfig() (*surreal.AppRepository, error) {

	config := surreal.NewApp()

	return surreal.NewAppRepository(config.DB)
}

func NewLogger() *logger.Logger {
	return logger.New()
}

// TODO: refactor this so it takes an [TestCase}
func HttpRunner(t *testing.T, testCase TestCase) {

	repo, err := GetConfig()

	if nil != err {
		log.Fatalf("failed to creat app : %v", err)
	}

	log.Info("connecting to database ")
	defer func() {
		repo.Close()
	}()
	r := gin.Default()
	router.SetRouter(r, repo, log)

	t.Run(testCase.Name, func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(testCase.Method, testCase.Url, strings.NewReader(testCase.Input))

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		defer req.Body.Close()
		r.ServeHTTP(w, req)

		expected := testCase.Output
		actual := w.Code

		if actual != expected {
			data, _ := io.ReadAll(w.Result().Body)
			t.Fatalf("Expected: %d, Actual: %d, Body %v", expected, actual, string(data))
		}

	})

}
