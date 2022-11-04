package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("Invalid content type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("Invalid user agent header")
	}

	if ContentTypeJson != "application/json" {
		t.Error("Invalid content type json header")
	}

	if ContentTypeXml != "application/xml" {
		t.Error("Invalid content type xml header")
	}

	if ContentTypeOctetStream != "application/octet-stream" {
		t.Error("Invalid content type octet stream header")
	}
}
