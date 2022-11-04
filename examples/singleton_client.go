package examples

import (
	"net/http"
	"time"

	"github.com/serapmtr/go-httpclient.git/gohttp"
	"github.com/serapmtr/go-httpclient.git/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Serap-Computer"). // Who is calling or sending HTTP request
		Build()

	return client
}
