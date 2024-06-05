package main

import "fmt"

// Maps: Key-value Pairs
// Maps are used to store data in key-value pairs. For example, map[string]int is a map with strig keys and integer values. For
func main() {
	fmt.Println("Map")
	m := map[string]int{"biraj": 19, "aadarsh": 20, "kiran": 21}
	fmt.Println(m)
	fmt.Println(m["biraj"]) //this is how we access the value use variable and keys
	fmt.Println(m["aadarsh"])
}
