package main

import (
	"fmt"
	"time"
)

func main()  {
	var time = time.Now()
	fmt.Printf("Hello, World! %T\n", time)
	fmt.Println("Time is ", time)
	
	// To format time we only use these dates/numbers, weird but true
	fmt.Println("Time is ", time.Format("01-02-2006"))
	fmt.Println("Time is ", time.Format("01-02-2006 15:04:05"))
	fmt.Println("Time is ", time.Format("01-02-2006 15:04:05 Monday"))

}