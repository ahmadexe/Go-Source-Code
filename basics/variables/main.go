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

	// Default values
	// bool: false
	// int and uint: 0
	// int8, int16, int32, int64: 0
	// uint8, uint16, uint32, uint64, uintptr: 0
	// float32 and float64: 0.0
	// complex64 and complex128: (0 + 0i)
	// string: ""
	// error: nil
	// struct, array, and slice: all of their elements are initialized to their zero values recursively

	var foo = "bar"
	fmt.Println(foo)
	// foo = 3 will throw an error, type once decided won't change
}
