package main

import (
	"fmt"
)

// To find the largest element in the array
func main() {
	arr := [7]int{12, 27, 12, 14, 33, 37, 42}
	flag := false
	var largest int = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
			flag = true
		}
	}
	if flag == true {
		fmt.Println("Largest element is: ", largest)
	} else {
		fmt.Println("Array is empty")
	}

	// Sum of array elements
	var sum int
	for _, value := range arr {
		sum = sum + value
		// flag = true
	}
	fmt.Println(sum)

	// Reverse an array
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}

	// Remove the duplicates
	var unique []int

	for i := 0; i < len(arr); i++ {
		isDuplicate := false
		for j := 0; j < len(unique); j++ {
			if unique[j] == arr[i] {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			unique = append(unique, arr[i])
		}
	}

	fmt.Println("Array with Unique Elements:", unique)

	// check if the array is sorted
	isSorted := true
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			isSorted = false
			break
		}
	}
	if isSorted {
		fmt.Println("Array is sorted:", arr)
	} else {
		fmt.Println("Array is not sorted:", arr)
	}
	// Sorting an array
	// Bubble Sort
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Array is sorted:", arr)

	// Returning the index of the target element
	target := 12
	var index int
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			index = i
		}
	}
	fmt.Println("Index of the target element is:", index)

	// Merging two sorted arrays
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr1 = append(arr1, arr2...)
	fmt.Println(arr1)

	// how many times the target element has been repeated
	count := 0
	// target number is defined above
	for i := 0; i < len(arr1); i++ {
		if arr[i] == target {
			count++
		}
	}
	fmt.Println("The target element", target, "has been repeated", count, "times")

}
