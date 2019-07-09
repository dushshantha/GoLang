package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var firstThree = nums[0:3]
	var lastTwo = nums[3:5]
	fmt.Println(firstThree, lastTwo)
	fmt.Println("Size of firstThree : ", len(firstThree))

	// Appending data to slices
	firstThree = append(firstThree, lastTwo...)
	fmt.Println(firstThree, lastTwo)

	firstThree = append(firstThree, 6, 7)
	fmt.Println(firstThree, nums)
}
