package main

import "fmt"

func main() {
	fmt.Println("Common Operations in Map")

	// Checking Length
	m := map[string]int{
		"biraj":   19,
		"aadarsh": 20,
		"kiran":   21,
	}
	length := len(m)
	fmt.Println("Length of the map is:", length)

	// Adding map variables
	m["Ram"] = 24
	fmt.Println("After adding Ram:", m)

	// Retrieving the map variables
	values := m["biraj"]
	fmt.Println("Value of biraj:", values)

	// Checking the existence of the map variables
	value, exists := m["aadarsh"]
	if exists {
		fmt.Println("Value of aadarsh:", value)
	} else {
		fmt.Println("Key aadarsh does not exist in the map")
	}

	// Deleting the map variable
	delete(m, "biraj")
	fmt.Println("After deleting biraj:", m)
}
