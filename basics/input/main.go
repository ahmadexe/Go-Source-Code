package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type of result %T\n", reader)
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", input)
}
