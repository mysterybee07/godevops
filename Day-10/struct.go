package main

import "fmt"

func main() {
	// Basic Declarations of a struct
	type Person struct {
		Name string
		Age  int
	}

	// Declaring and instantiating together
	person1 := Person{
		Name: "John Doe",
		Age:  30,
	}
	fmt.Println(person1)
	fmt.Println(person1.Name)
	fmt.Println(person1.Age)

	// Using the new Keyword
	person2 := new(Person)
	person2.Name = "Jane Doe"
	person2.Age = 35

	fmt.Println(person2)

	// declaration using anynomous struct
	person3 := struct {
		Country string
		Code    int
	}{
		Country: "USA",
		Code:    1,
	}
	fmt.Println(person3)

	// Nested Structs
	type Address struct {
		City  string
		State string
	}
	type PersonDetails struct {
		personalInfo Person
		Address      Address
	}
	person4 := PersonDetails{
		personalInfo: Person{
			Name: "Bob Smith",
			Age:  40,
		},
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	fmt.Println(person4)

	// Embedded (Anonymous) Fields
	type Manager struct {
		Person
		Department string
	}

	mgr := Manager{
		Person: Person{
			Name: "Alice Johnson",
			Age:  35,
		},
		Department: "Sales",
	}
	fmt.Println(mgr)

}
