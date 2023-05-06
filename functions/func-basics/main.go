package main

func main()  {
	c := add(1, 2)
	println(c)

	hi()

	total, msg := superAdders(1, 2, 3, 4, 5, 6)
	println(total, msg)
}

func add(a int, b int) int {
	return a + b
}

func hi() {
	println("hi")
}

func superAdders(val ...int) (int, string) {
	total := 0
	for _, num := range val {
		total += num
	}
	return total, "hi"
} 