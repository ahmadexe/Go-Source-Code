package main

import "fmt"


func main()  {
	languages := make(map[string]string)
	languages["py"] = "python"
	languages["js"] = "javascript"
	languages["d"] = "dart"
	languages["go"] = "golang"

	for key, value := range languages {
		println(key, value)
	}

	// delete key
	delete(languages, "js")
	fmt.Println(languages)
}