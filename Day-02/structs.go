package main

import "fmt"

// Structs: Custom Types
// Structs are used to create custom types with mixed data types. For example, type Person struct{Name string, age int}

func main() {
	type Person struct {
		name string
		age  int
	}
	fmt.Println("Structs")
	// var p Person
	// p.name = "John"
	// p.age = 20
	// fmt.Println(p)
	p := Person{
		name: "Biraj",
		age:  20,
	}
	fmt.Println(p)

	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
}
