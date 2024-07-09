package main

import (
	"fmt"
	"strings"
)

func main() {
	// String concatenation
	str1 := "Hello"
	str2 := "World"
	concatenation := str1 + " " + str2
	fmt.Println(concatenation)
	// concate:=
	// Splitting a string
	sentence := "Go is statically typed open source programming language."
	words := strings.Split(sentence, " ")
	fmt.Println(words)

	// Replacing a substring
	replaced := strings.Replace(sentence, "Go", "Golang", 1)
	fmt.Println(replaced)

	// Upper case and lower case
	upper := strings.ToUpper(sentence)
	lower := strings.ToLower(sentence)
	title := strings.ToTitle(sentence)
	fmt.Println(title)
	fmt.Println(upper)
	fmt.Println(lower)

	//Trimming
	spaceyString := "   trim me   "
	trimmed := strings.TrimSpace(spaceyString)
	fmt.Println(trimmed)

	// substring using slicing
	// Go doesnot have built in substring methods
	start := 3
	end := 11
	if end <= len(sentence) && start <= end {
		substring := sentence[start:end]
		fmt.Println("Substring:", substring)
	} else {
		fmt.Println("Invalid range of substring")
	}

	// checking if string contains a substring
	contains := strings.Contains(sentence, "source")
	fmt.Println("Contains 'source':", contains)

	// Finding a substring index
	index := strings.Index(sentence, "source")
	fmt.Println(index)
}
