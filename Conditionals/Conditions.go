package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Println(i)
		case 1:
			fmt.Println(i)
		case 2:
			fmt.Println(i)
		}

		if i == 0 {
			fmt.Println("If called")
		} else if i == 1 {
			fmt.Println("else if called")
		} else {
			fmt.Println("else called")
		}
	}

}
