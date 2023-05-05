package main

import "fmt"

func main() {
	var username string = "Ahmad"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isAhmad bool = true
	fmt.Println(isAhmad)
	fmt.Printf("Variable is of type: %T\n", isAhmad)

	// we can use uint8, uint16, uint32, uint64 and int8, int16, int32, int64 alongisde pointers
	var age int = 20
	fmt.Println(age)
	fmt.Printf("Variable is of type: %T\n", age)

	// upto 7 decimal points 
	var height float32 = 1.75
	fmt.Println(height)
	fmt.Printf("Variable is of type: %T\n", height)

	// upto 15 decimal points
	var weight float64 = 70.5
	fmt.Println(weight)
	fmt.Printf("Variable is of type: %T\n", weight)
}
