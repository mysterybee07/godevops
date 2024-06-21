package main

import (
	"fmt"
	"math"
)

func SecondLargest(arr []int) int {
	if len(arr) < 2 {
		fmt.Println("Array must contain at least two distinct elements")
		return 0
	}

	largest := math.MinInt32
	secondLargest := math.MinInt32

	for _, num := range arr {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest && num != largest {
			secondLargest = num
		}
	}

	if secondLargest == math.MinInt32 {
		fmt.Println("No second largest element found")
		return 0
	}

	return secondLargest
}

func main() {
	arr1 := []int{20, 15, 17, 18, 19}
	fmt.Println("Second largest element is:", SecondLargest(arr1))

	arr2 := []int{10, 10, 10, 10, 10}
	fmt.Println("Second largest element is:", SecondLargest(arr2))
}
