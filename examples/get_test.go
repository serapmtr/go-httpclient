package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/serapmtr/go-httpclient.git/gohttp"
)

// always this function starts first
func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples' ")

	// Tell the HTTP library to mock any further request
	// from here
	gohttp.StartMockServer()

	// we have to do that for starting other tests
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// with this, we clean all mocks, so we can recreate them
		gohttp.FlushMocks()

		// we use mock because we can't control what
		// this request returns. Application will
		//  be up and running id we don't

		// if we delete the mock, then api will call
		// actual http call
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})
		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message recieved")
		}

	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})
		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message recieved")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})
		endpoints, err := GetEndpoints()

		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got `%s`", err.Error()))
		}
		if endpoints == nil {
			t.Error("endpoints were expected at this point")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current_user_url")
		}

	})

}
