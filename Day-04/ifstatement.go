package main

import "fmt"

// Control Structure (if, switch, loops)
// if Statement: The if statement in Go is used for conditional execution of code blocks.
// It can optionally include an else block for alternative execution when the if condition is not met.

func main() {
	x := 99

	if x > 100 {
		fmt.Println("x is greater than 100")
	} else if x == 100 {
		fmt.Println("x is equal to 100")
	} else {
		fmt.Println("x is less than 100")
	}
	// if x := 99; x < 100 { In this case x is not accessible outside of this if block
	// 	fmt.Println("x is less than 100")
	// }

}
