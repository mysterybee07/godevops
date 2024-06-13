package main

import "fmt"

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}

}

func main() {
	originalSlice := []int{1, 2, 3, 4, 5}

	modifySlice(originalSlice)

	fmt.Println(originalSlice)

}
