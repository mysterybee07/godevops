package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Error opening file1: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cont bool
	contains := "spaces"
	for scanner.Scan() {
		line := scanner.Text()
		contain := strings.Contains(line, contains)
		cont = contain
	}
	fmt.Println("Contains 'spaces'", cont)
}
