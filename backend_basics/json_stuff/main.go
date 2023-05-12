package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	courses := []course{
		{"Flutter", 12, "Udemy", "1234", []string{"flutter", "mobile"}},
		{"Hyperledger Fabric", 15, "Udemy", "12345", []string{"blockchain"}},
		{"React", 10, "Udemy", "123456", []string{"react", "frontend"}},
		{"Go", 10, "Udemy", "1234567", nil},
	}

	prettyJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(prettyJson))
}

func main() {
	EncodeJson()
}
