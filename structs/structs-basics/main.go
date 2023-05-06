package main

import "fmt"

type User struct {
	Username string
	Email string
	Password string
}

func main()  {
	ahmad := User{"Ahmad", "muahmad710@gmail.com", "123456"}
	fmt.Println(ahmad.Username)
	fmt.Println(ahmad)
}