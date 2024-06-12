package main

import "fmt"

func main() {
	//
	fmt.Println("Slices in Go")
	fmt.Println("Without initializing it")
	var slice []int
	fmt.Println(slice)

	fmt.Println("With initializing it")
	slice2 := []int{1, 2, 3, 4}
	fmt.Println(slice2)

	fmt.Println("with make function")
	slice3 := make([]int, 5)
	slice3[0] = 4
	slice3[1] = 5
	slice3[2] = 6
	slice3[3] = 7
	slice3[4] = 8
	fmt.Println(slice3)

	fmt.Println("with make function with capacity")
	slice4 := make([]int, 3, 5)
	slice4[0] = 4
	slice4[1] = 5
	slice4[2] = 6
	fmt.Println(slice4)
	slice4 = append(slice4, 7)
	slice4 = append(slice4, 8)
	slice4 = append(slice4, 9)
	fmt.Println("Slice after appending", slice4)
	fmt.Println(cap(slice4)) //on appending elements more than its capacity it doubles the capapcity which help in memory management

	fmt.Println("From an array")
	array := [5]int{1, 2, 3, 4, 5}
	slice5 := array[1:3]
	fmt.Println(slice5)

}
