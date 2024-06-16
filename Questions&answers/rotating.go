// Rotate an Array:
// Rotate the elements of an array to the right by a given number of steps.

package main

import "fmt"

func Reverse(arr []int, first, last int) {
	for first < last {
		arr[first], arr[last] = arr[last], arr[first]
		first++
		last--
	}
}

func Rotate(arr []int, step int) {
	step = step % len(arr)
	Reverse(arr, 0, len(arr)-1)
	Reverse(arr, 0, step-1)
	Reverse(arr, step, len(arr)-1)
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}
	steps := 2
	Rotate(arr, steps)
	fmt.Println(arr)
}
