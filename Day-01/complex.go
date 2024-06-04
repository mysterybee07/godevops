package main

import "fmt"

func main() {
	fmt.Println("Float type\n")
	var a float32 = 3.14     //single precision floating numbers
	var b float64 = 3.141592 //double precision floating numbers
	fmt.Printf("a=%v \n b=%v \n", a, b)

	fmt.Println("Complex type\n")
	var c complex64 = 3 + 4i  //complex number with float32 real and imaginary parts
	var d complex128 = 3 + 4i //complex number with float64 real and imaginary parts
	fmt.Printf("c=%v \n d=%v \n", c, d)
}
