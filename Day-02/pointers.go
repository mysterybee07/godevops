package main

import "fmt"

// Pointers: Reference Types
// A pointer holds the memory address of a variable.
// for example: var p *int is a pointer to an integer variable

func main() {
	fmt.Println("Pointers")
	var s string = "pointer to a variable"
	var p *string = &s
	fmt.Println(s)  //prints the value of s
	fmt.Println(p)  //prints the memory address of s
	fmt.Println(*p) //prints the value at the memory address stored in p
}
