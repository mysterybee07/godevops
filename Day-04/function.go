package main

// Functions:
// A function in Go is a block of code that performs a specific task.
// It can take input parameters and return one or more values
// Types: Basic Function, Named Return Values, Multiple Return Values, Variadic Functions
import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum_of_num(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	x := add(20, 30)
	fmt.Println(x)
	a, b := swap("Hello", "Biraj")
	fmt.Println(a, b)

	sum := 100
	m, n := split(sum)
	fmt.Println(m, n)

	fmt.Println("variadic function:")
	sum_result := sum_of_num(1, 2, 3, 4)
	fmt.Println(sum_result)
}
