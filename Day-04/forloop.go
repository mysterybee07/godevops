package main

import "fmt"

// Loops statements
// Go has only one looping construct, the for loop. It can be
func main() {
	fmt.Println("Loops in go")
	// Basic loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for as a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	if i == 0 {
	// 		break
	// 	}
	// }

	// for loop with range (for slices, arrays, maps, strings)
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// for loop with range (for strings)
	for index, runeValue := range "hello" {
		fmt.Println(index, runeValue)
	}

}
