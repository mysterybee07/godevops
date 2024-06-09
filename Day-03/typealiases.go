package main

import "fmt"

// Type Aliases:
// Custome Names for Existing Types
// Useful for readability or when transitioning code.
// For example, type byteSize int64 creates a ByteSize type as an alias for int64

func main() {
	fmt.Println("Type Aliases")
	type ByteSize int64
	var b ByteSize = 1024
	fmt.Println(b)
}
