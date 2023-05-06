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

	// if we do this [5] i.e provide a fixed length an array will be declared
	var num []int = []int{1, 2, 3, 4, 5}
	num = append(num, 0)
	fmt.Println(num)
	sort.Ints(num)
	fmt.Println(num)
	fmt.Println(sort.IntsAreSorted(num))

	var techs = []string{"Golang", "Python", "Kotlin", "Swift", "Dart"}
	fmt.Println(techs)
	sort.Strings(techs)
	fmt.Println(techs)
	fmt.Println(sort.StringsAreSorted(techs))

	// Removing elements
	techs = append(techs[:1], techs[2:]...)
	fmt.Println(techs)

	// Remove last 
	techs = techs[:len(techs)-1]

	// Remove using name of element from a slice
	for i, v := range techs {
		if v == "Kotlin" {
			techs = append(techs[:i], techs[i+1:]...)
		}
	}
	fmt.Println(techs)


}