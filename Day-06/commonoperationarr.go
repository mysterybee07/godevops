package main

import "fmt"

// Copying arrays to functions
func modifyArrays(arr [5]int, index, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {
	// Length of an array
	arr1 := [5]int{1, 2, 3, 4}
	arr3 := [5]int{1, 2, 3, 4}
	fmt.Println("The length of an array is", len(arr1))

	// Copying an array
	arr2 := arr1
	fmt.Println("The copied array is", arr2)

	// Modifying elements of an array
	arr1[3] = 5
	fmt.Println("The modified array is", arr1)

	// copying arrays to functions
	fmt.Println(modifyArrays([5]int(arr1), 2, 9))

	// Comparing arrays
	isEqual := arr2 == arr3

	fmt.Println(isEqual)

}
