package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	PerformGetReq()
}

func PerformGetReq()  {
	const url string = "https://jsonplaceholder.typicode.com/posts"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}