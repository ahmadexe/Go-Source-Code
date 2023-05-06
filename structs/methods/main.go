package main

import "fmt"

func main()  {
	ahmad := User{"Ahmad", "12345"}
	ahmad.SetUn("!Ahmad")
	fmt.Println(ahmad.Username)
	fmt.Println(ahmad.GetPwd())
}

type User struct {
	Username string
	Password string
}

func (u *User) SetUn(n string) {
	u.Username = n
}

func (u User) GetPwd() string {
	return u.Password
}
