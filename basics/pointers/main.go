package main

import "fmt"

func main()  {
	var num int = 5
	var numPointer *int = &num

	fmt.Println("num: ", *numPointer)
	fmt.Println("numPointer: ", numPointer)
}