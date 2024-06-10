package main

import "fmt"

// Arrays
// An array has a fixed size defined at compile time. For example, var arr [5]int declares an array of five integers.

func main() {
	fmt.Println("Arrays")
	var arr [5]int = [5]int{1, 2, 3, 4} //by default the array will be of size 5 as it adds 0 as fifth element
	fmt.Println(arr[3])
}
