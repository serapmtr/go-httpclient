package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")

	client.Headers = commonHeaders

	// Execution
	requstHeaders := make(http.Header)
	requstHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requstHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid content type recieved")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid user agent recieved")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid request id recieved")
	}
}

func TestGetRequestBody(t *testing.T) {

	client := httpClient{}
	t.Run("noBodyNilResponse", func(t *testing.T) {

		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("No error expected when passing a nil body")
		}

		if body != nil {
			t.Error("No body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("No error expected when marshaling slice as json")
		}

		if string(body) != `["one", "two"]` {
			t.Error("Invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {})

}
