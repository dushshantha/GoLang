package main

import "fmt"

func main() {
	var x int
	x = 4
	fmt.Println("x:", x)

	var y = 5
	fmt.Println("y:", y)
	fmt.Println("x-y = ", x-y)

	var a, b = 3, "string"
	fmt.Println(a, b)

	// Also, yo can omit var altogether
	s := "String"
	fmt.Println(s)

	//Constants
	const c int = 9
	fmt.Println(c)
}
