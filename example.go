package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/serapmtr/go-httpclient.git/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	getUrls()
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `jsom:"last_name"`
}

func getUrls() {
	// client := gohttp.HttpClient{}  // you can't reach an interface directly

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
