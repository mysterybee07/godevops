package main

import (
	"fmt"
)

func main() {
	// Declarations using var keyword
	var map1 map[string]int
	fmt.Println("Map is declared using var keyword", map1)

	// Using map literals
	map2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("Map is declared using map literals", map2)

	// Creating a map with make function
	map3 := make(map[string]int)
	map3["one"] = 1
	map3["two"] = 2
	fmt.Println("Map is created using make function", map3)

	// Using make function with specific size
	map4 := make(map[string]int, 5)
	map4["four"] = 4
	map4["five"] = 5
	map4["six"] = 6
	fmt.Println("Map is created using make function with specific size", map4)

	// Using Contruct as Map values
	type Student struct {
		Name string
		Age  int
	}
	map5 := make(map[string]Student)
	map5["student1"] = Student{
		Name: "Biraj Pudasaini",
		Age:  20,
	}
	map5["student2"] = Student{
		Name: "Shyam Pudasaini",
		Age:  21,
	}

	fmt.Println("Map is created using Construct as Map values", map5)

	// Declarations using Nested Map
	map6 := make(map[string]map[string]int)
	map6["fruit"] = map[string]int{
		"apple":  10,
		"banana": 20,
	}
	map6["vegetable"] = map[string]int{
		"carrot": 5,
		"potato": 7,
	}

	fmt.Println("Map is declared using Nested Map", map6)
}
