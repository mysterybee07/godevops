package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// Complex operations
	var complex1 complex128 = complex(3, 4)
	var complex2 complex128 = complex(5, 6)
	fmt.Println("Complex addition:", complex1+complex2)
	fmt.Println("Complex subtraction:", complex1-complex2)
	fmt.Println("Complex multiplication:", complex1*complex2)
	fmt.Println("Complex division:", complex1/complex2)

	// Math Functions
	fmt.Println("powers:", math.Pow(2, 3))
	fmt.Println("Square root:", math.Sqrt(16))
	fmt.Println("Logarithm:", math.Log(math.E))
	fmt.Println("Sin", math.Sin(math.Pi/4))
	fmt.Println("Tan", math.Tan(math.Pi/4))

	// string to Number Parsing

	numStr := "3.14159"
	parsedFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Error parsing string to float: ", err)
	} else {
		fmt.Println("Parsed float:", parsedFloat)
	}
}
