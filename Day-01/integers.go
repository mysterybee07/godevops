package main

import "fmt"

// Variable Types
// 1. Basic Types: int, float64, bool, etc.
// 2. Composite Types: array, struct, pointer, slice, map, channel, interface etc
func main() {
	fmt.Println("Signed Integers\n")
	var a int8 = 127    //ranges from -128 to 127
	var b int16 = 32767 //ranges from -32768 to 32767
	fmt.Printf("a=%v \n b=%v \n", a, b)

	fmt.Println("\nUnSigned Integers\n")
	var c uint8 = 255    //ranges from 0 to 255
	var d uint16 = 65535 //ranges from 0 to 65535
	fmt.Printf("c=%v \n d=%v \n", c, d)

	fmt.Println("Machine Dependent Types\n")
	// Size depends on the system architecture

	var e int = 12345678
	var f uint = 123456
	var g uintptr = 0xdeadbeef
	fmt.Printf("e=%v \n f=%v \n g=%v \n", e, f, g)

}
