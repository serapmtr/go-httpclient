package gohttp

import (
	"testing"
)

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

		if string(body) != `["one","two"]` {
			t.Error("Invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {})

}
