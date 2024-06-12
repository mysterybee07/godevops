package main

import "fmt"

func main() {
	nums := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Iterating over the array using traditionla for loop")
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	fmt.Println("Iterating over the array using for range")
	for _, value := range nums {
		fmt.Println(value)
	}
}
