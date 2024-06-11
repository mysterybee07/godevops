package main

import "fmt"

func main() {
	// Array declared without any initialization
	var arr [5]int
	fmt.Println(arr)

	// Array declared with initialization
	var arr2 [5]string = [5]string{"ram", "shyam", "how"}
	fmt.Println(arr2)

	// Shorthand declaration
	arr3 := [3]int{5, 6, 7}
	fmt.Println(arr3)

	// Initializing with specific elements
	arr4 := [4]int{1: 2, 3: 4}
	fmt.Println(arr4)

	// Using the ... operator to infer length
	arr5 := [...]int{1, 2, 3, 4}
	fmt.Println(arr5)

	// Arrays of Arrays(Multi-Dimensional arrays)
	arr6 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr6)

	// pointer Arr
	var pointerArr [4]*int
	num1, num2, num3, num4 := 4, 5, 6, 7
	pointerArr[0] = &num1
	pointerArr[1] = &num2
	pointerArr[2] = &num3
	pointerArr[3] = &num4
	fmt.Println("Array with Pointers:")
	for i := 0; i < len(pointerArr); i++ {
		fmt.Printf("pointerArr[%d]: %d	\n", i, *pointerArr[i])
	}

	// Array of Structs
	type Students struct {
		Name string
		Age  int
	}

	students := [...]Students{
		{
			Name: "Biraj",
			Age:  20,
		},
		{
			Name: "Shyam",
			Age:  21,
		},
		{
			Name: "Ram",
			Age:  22,
		},
	}
	fmt.Println(students)

}
