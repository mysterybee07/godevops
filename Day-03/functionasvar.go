package main

import "fmt"

// Function as Variable
// Function can be assigned to variables.
// For instance, var f func(int) int declares a f as a variable that can hold a function taking an int and returning an int

func main() {
	var v func(int) int //here the variable v holds a function
	v = func(x int) int {
		return x + 1 //defining a function
	}
	fmt.Println(v(1))
}
