package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	res, err := http.Get("https://lco.dev")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status code is", res.Status)
	fmt.Println("Content length is", res.ContentLength)
	fmt.Println("Content type is", res.Header.Get("Content-Type"))
	fmt.Println("Server is", res.Header.Get("Server"))
	fmt.Println("Date is", res.Header.Get("Date"))
	fmt.Println("Content encoding is", res.Header.Get("Content-Encoding"))
	fmt.Println("Content length is", res.Header.Get("Content-Length"))
	fmt.Println("Transfer encoding is", res.Header.Get("Transfer-Encoding"))
	fmt.Println("Connection is", res.Header.Get("Connection"))
	fmt.Println("Strict transport security is", res.Header.Get("Strict-Transport-Security"))
	fmt.Println("X-Frame-Options is", res.Header.Get("X-Frame-Options"))
	fmt.Println("X-Content-Type-Options is", res.Header.Get("X-Content-Type-Options"))
	fmt.Println("X-XSS-Protection is", res.Header.Get("X-XSS-Protection"))
	fmt.Println("Referrer-Policy is", res.Header.Get("Referrer-Policy"))
	fmt.Println("Feature-Policy is", res.Header.Get("Feature-Policy"))
	fmt.Println("Content-Security-Policy is", res.Header.Get("Content-Security-Policy"))
	fmt.Println("Cache-Control is", res.Header.Get("Cache-Control"))

	fmt.Println("----------------------------------------------------------------------------------")
	
	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	println(content)
}