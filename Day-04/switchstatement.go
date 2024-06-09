package main

import (
	"fmt"
)

// The switch statement in Go is used for multiple conditional executions.
// Its a cleaner way to write a sequence of if-else if-else statements

func main() {

	// fmt.Println("The switch statement")
	// // days := 8
	// var days int
	// fmt.Println("Enter a number between 1 and 7: ")
	// fmt.Scanf("%d", &days)
	// switch days {
	// case 1:
	// 	fmt.Println("Sunday")
	// case 2:
	// 	fmt.Println("Monday")
	// case 3:
	// 	fmt.Println("Tuesday")
	// case 4:
	// 	fmt.Println("Wednesday")
	// case 5:
	// 	fmt.Println("Thursday")
	// case 6:
	// 	fmt.Println("Friday")
	// case 7:
	// 	fmt.Println("Saturday")
	// default:
	// 	fmt.Println("Invalid day")
	// 	break
	// }

	// today := time.Now().Weekday()

	// switch today := time.Now().Weekday(); today {
	// case time.Monday:
	// 	fmt.Println("Monday")
	// case time.Tuesday:
	// 	fmt.Println("Tuesday")
	// case time.Wednesday:
	// 	fmt.Println("Wednesday")
	// case time.Thursday:
	// 	fmt.Println("Thursday")
	// case time.Friday:
	// 	fmt.Println("Friday")
	// case time.Saturday:
	// 	fmt.Println("Saturday")
	// case time.Sunday:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	x := 50

	switch {
	case x < 0:
		fmt.Println("Negative")
	case x == 0:
		fmt.Println("Zero")
	case x > 0:
		fmt.Println("Positive")
	default:
		fmt.Println("Invalid")
		break
	}

}
