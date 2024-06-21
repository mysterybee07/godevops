package main

import "fmt"

func SecondLargest(arr []int) int {
	largest := arr[0]
	secondLargest := 0
	if len(arr) < 2 {
		return 0
	}
	if len(arr) == 2 {
		if arr[0] < arr[1] {
			secondLargest = arr[0]
		} else {
			secondLargest = arr[1]
		}

	} else {
		for i, _ := range arr {
			if largest > arr[i] {
				secondLargest = largest
				largest = arr[i]
			} else if largest > arr[i] && secondLargest < arr[i] {
				secondLargest = arr[i]
			}
		}

	}
	return secondLargest
}

func main() {
	arr := []int{22, 11}
	secondLargest := SecondLargest(arr)
	fmt.Println("Second largest element is: ", secondLargest)
}
