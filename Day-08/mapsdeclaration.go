package main

import "fmt"

func main() {
	fmt.Println("Declaration of map")
	// Using var keyword without initializing it
	var m map[string]int
	fmt.Println(m)

	// Using map literals
	m2 := map[string]int{
		"biraj":   19,
		"aadarsh": 20,
		"kiran":   21,
	}
	fmt.Println(m2)

	// Using the make function
	m3 := make(map[string]int)
	m3["biraj"] = 19
	m3["aadarsh"] = 20
	fmt.Println("The maps declared using make function is", m3)

	// Using make function with specific size
	m4 := make(map[string]int, 10) //maps with the initial capacity of 10
	m4["biraj"] = 19
	m4["aadarsh"] = 20
	m4["kiran"] = 21
	m4["ram"] = 22

	fmt.Println("The maps declared using make function with specific size", m4)

	// Using struct as a map values
	type Students struct {
		Name string
		Age  int
	}
	m5 := make(map[string]Students)
	m5["biraj"] = Students{
		Name: "Biraj",
		Age:  20,
	}
	fmt.Println(m5)

	// Using nested maps
	m6 := make(map[string]map[int]string)
	m6["biraj"] = map[int]string{1: "Biraj", 2: "Ram", 3: "Shyam"}
	fmt.Println(m6)

}
