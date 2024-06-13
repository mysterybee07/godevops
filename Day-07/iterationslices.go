package main

import "fmt"

func main() {
	fmt.Println("Slices iteration")
	slice := []string{"apple", "banana", "pineapple", "jackfruit", "cherry"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println(index, value)
	}
}
