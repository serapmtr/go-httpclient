package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	gohttpmock "github.com/serapmtr/go-httpclient.git/gohttp_mock"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeoutFromGithub", func(t *testing.T) {
		gohttpmock.MockupServer.DeleteMocks()

		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			Error: errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("No repo expected when we get a timeout from github")
		}
		if err == nil {
			t.Error("An error is expecting when we get a timeout from github")
		}
		if err.Error() != "timeout from github" {
			fmt.Println(err.Error())
			t.Error("invalid error message")
		}
	})

	t.Run("timeoutFromGithub", func(t *testing.T) {
		gohttpmock.MockupServer.DeleteMocks()

		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		if err != nil {
			t.Error("No error expected when we get valid response from github")
		}
		if repo == nil {
			t.Error("A valid repo was expected at this point")
		}
		if repo.Name != repository.Name {
			t.Error("invalid repository name obtained from github")
		}
	})

}
