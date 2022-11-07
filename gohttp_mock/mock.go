package gohttpmock

import (
	"fmt"
	"net/http"

	"github.com/serapmtr/go-httpclient/core"
)

// mocking key => method + url + body

// The Mock structure provides a clean way to configure HTTP mocks based on
// the combination between request method, URL and request body
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
