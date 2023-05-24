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

func DecodeJson() {
	randJson := []byte(`
	{
		"name": "Flutter",
		"price": 12,
		"platform": "Udemy",
		"tags": [
				"flutter",
				"mobile"
		]
	}`)

	if (json.Valid(randJson)) {
		var course course
		err := json.Unmarshal(randJson, &course)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", course)
	} else {
		fmt.Println("Invalid JSON")
	}

	var someMap map[string]interface{}
	err := json.Unmarshal(randJson, &someMap)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\n", someMap)
	}

}

func main() {
	// EncodeJson()
	DecodeJson()
}
