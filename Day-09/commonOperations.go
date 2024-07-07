package main

import "fmt"

func modifyValue(ptr *int) {
	*ptr = 30
}
func main() {
	fmt.Println("Common Operations in Pointer")
	// Dereferencing a pointer
	a := 20
	ptrA := &a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", ptrA)
	fmt.Println("Dereferencing a pointer gives:", *ptrA)

	// Changing the value at a pointer
	*ptrA = 35
	fmt.Println("After changing the value, a:", a)

	// Comparing Pointers
	b := 20
	ptrB := &b
	fmt.Println("PtrA and PtrB point to the same address", ptrB == ptrA)
	p1 := &a
	p2 := &b
	p3 := &b
	fmt.Println("P2 and P1 point to the same address", p1 == p2)
	fmt.Println("P2 and P3 point to the same address", p2 == p3)

	// Passing Pointers to Functions
	fmt.Println("value of a:", a)
	modifyValue(ptrA)
	fmt.Println("After modifying value in function, a:", a)
}
