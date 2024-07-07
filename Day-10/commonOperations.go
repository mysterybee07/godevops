package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) DisplayInfo() {
	println("Hello, my name is", p.Name)
	println("I am", p.Age, "years old")
}

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	//Instantiating
	person := Person{
		Name: "John Doe",
		Age:  30,
	}
	person.DisplayInfo()

	// Modifying
	person.Age = 35
	person.DisplayInfo()

	// Pointers to struct
	person1 := &Person{
		Name: "Jane Doe",
		Age:  35,
	}
	person1.DisplayInfo()

	// Tagging
	employee := Employee{
		ID:   1,
		Name: "Alice Johnson",
	}
	t := reflect.TypeOf(employee)
	field, _ := t.FieldByName("Name")
	fmt.Println("Tag on Name field:", field.Tag.Get("json"))
	println(employee.Name)
	println(employee.ID)
}
