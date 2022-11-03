package examples

import (
	"net/http"
	"time"

	"github.com/serapmtr/go-httpclient.git/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	currentClient := http.Client{}
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetHttpClient(&currentClient).
		Build()

	return client
}
