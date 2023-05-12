package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PostRequest()
}

func PostRequest() {
	const url string = "https://jsonplaceholder.typicode.com/posts"

	// fake json payload
	payload := strings.NewReader(`
		{
			"title": "foo",
			"body": "bar",
			"userId": 1
		}
	`)

	// send post request
	response, err := http.Post(url, "application/json", payload)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
