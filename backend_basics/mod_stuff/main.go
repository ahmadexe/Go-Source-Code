package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")
	hello()
}

func hello() {
	fmt.Println("Hello!")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Request deals with the parameters of the request, URL, headers, body, etc.
	// ResponseWriter is used to fill in the HTTP response. Here our simple response is just “hello\n”.
	w.Write([]byte("<h1>Hello World</h1>"))
}
