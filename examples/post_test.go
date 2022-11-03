package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/serapmtr/go-httpclient.git/gohttp"
)

func TestCreateRepo(t *testing.T) {
	gohttp.FlushMocks()

	gohttp.AddMock(gohttp.Mock{
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
}
