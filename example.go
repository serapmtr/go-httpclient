package main

import (
	"fmt"

	"github.com/serapmtr/go-httpclient.git/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {
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

	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}

	fmt.Println(user.FirstName)

	// fmt.Println(response.StatusCode)

	// bytes, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(bytes))
}

// func createUser(user User) {
// 	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(response.StatusCode)

// 	bytes, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(bytes))
// }
