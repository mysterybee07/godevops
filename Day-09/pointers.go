package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang")

	// Using var
	var a *int
	fmt.Println(a)

	// Pointer to existing variable
	var b int = 9
	var p *int = &b
	fmt.Println(p)

	// creating a pointer with new
	ptr := new(int)
	*ptr = 10
	fmt.Println(ptr)

	// Pointer to Pointer
	ptr2 := &ptr
	fmt.Println(**ptr2)

}
