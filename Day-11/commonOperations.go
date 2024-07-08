package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, Iam " + p.Name
}

// Type Assertions
func PrintDetails(i interface{}) {
	str, ok := i.(string)
	if ok {
		println(str)
	} else {
		println("Invalid Type")
	}
}

// Interface with type switch
func Describer(i interface{}) {
	switch v := i.(type) {
	case int:
		println("The value is an integer", v)
	case string:
		println("The value is a string", v)
	case Person:
		println("The value is a person", v.Greet())
	default:
		println("Unknown type")
	}
}

func main() {
	// implementing an interface
	person := Person{Name: "John Doe"}
	var greeter Greeter = person
	fmt.Println(greeter.Greet())

	// type assertions
	PrintDetails(person)
	PrintDetails(10)
	PrintDetails("Hello, World!")

	// interface with type switch
	Describer(person)
	Describer(10)
	Describer("Hello, World!")
	Describer(true) // Invalid Type
}
