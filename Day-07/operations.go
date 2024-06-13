package main

import (
	"fmt"
	"sort"
)

func main() {
	// Length and capacity of slice
	numbers := []int{}
	length := len(numbers)
	fmt.Println(length)
	capacity := cap(numbers)
	fmt.Println(capacity)

	// Appending to the slice
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers)

	moreNumbers := []int{6, 7, 8}
	numbers = append(numbers, moreNumbers...)
	fmt.Println(numbers)

	// Slicing a slice
	slicedNumbers := numbers[1:3]
	fmt.Println("The sliced numbers are", slicedNumbers)

	slicedNumbers = numbers[1:]
	fmt.Println("The again sliced numbers are", slicedNumbers)

	// Deleting a slice
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

	// Checking if slice is empty
	if len(numbers) == 0 {
		fmt.Println("The slice is empty")
	} else {
		fmt.Println("The slice is not empty")
	}

	// Sorting a slice
	sort.Ints(numbers)
	fmt.Println(numbers)

	// Finding the element
	target := 8
	found := false

	for _, value := range numbers {
		if value == target {
			found = true
			break
		}
	}
	if found {
		fmt.Println("The element is found")
	} else {
		fmt.Println("The element is not found")
	}

}
