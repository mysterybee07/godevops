package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// open the file
	file, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var cleanedLines []string
	// Read each line trim space and store in cleaned lines
	for scanner.Scan() {
		line := scanner.Text()
		cleanedLine := strings.TrimSpace(line)
		cleanedLines = append(cleanedLines, cleanedLine)
	}
	// Handle any error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	file, err = os.Create("file1.txt")
	if err != nil {
		fmt.Println("Error opening file for writing:", err)
		return
	}
	defer file.Close()

	// Write the cleaned lines back to the file
	for _, line := range cleanedLines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File cleaned and updated successfully!")
	// fmt.Println("File cleaned successfully!")
}
