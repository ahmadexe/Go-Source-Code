package main

import (
	"fmt"
	"sort"
)

func main()  {
	
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)
	fruits = append(fruits, "papaya", "mango", "kiwi")
	fmt.Println(fruits)

	var users = make([]string, 5)
	users[0] = "Rizky"
	users[1] = "Khaironi"
	users[2] = "Rizky Khaironi"
	users[3] = "Rizky Khaironi Putra"
	users[4] = "Rizky Khaironi Putra Pratama"
	fmt.Println(users)

	// Appends reallocates memory
	users = append(users, "Rizky Khaironi Putra Pratama Siahaan")
	fmt.Println(users)

	var num []int = []int{1, 2, 3, 4, 5}
	num = append(num, 0)
	fmt.Println(num)
	sort.Ints(num)
	fmt.Println(num)
	fmt.Println(sort.IntsAreSorted(num))
}