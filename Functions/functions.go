package main

import "fmt"

func main(){
	fmt.Println(calc(2,5))
}

func calc(num1 int, num2 int)(int, int) {
	return num1 + num2, num1 * num2
}