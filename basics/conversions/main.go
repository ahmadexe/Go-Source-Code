package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating 1-5: ")
	var input, _ = reader.ReadString('\n')

	// Handeling errors, if any
	inputint, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	inputint++
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Your rating is: ", inputint)
	}
}
