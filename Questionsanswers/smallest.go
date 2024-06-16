// Find the Minimum Element:
// Given an array of integers, find and return the smallest element.

package main

import "fmt"

func main() {
	arr := []int{22, 11, 33, 25, 16, 6, 55}
	smallest := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}
	fmt.Println("The smallest element in the array is ", smallest)
}
