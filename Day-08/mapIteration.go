package main

import "fmt"

func main() {
	fmt.Println("Iterating over the maps")
	mapVariables := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	mapVariables["key4"] = 4

	for key, value := range mapVariables {
		fmt.Println("The value for key", key, "is", value)
	}

}
