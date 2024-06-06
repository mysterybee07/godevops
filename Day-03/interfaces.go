package main

// Interface Types
// Method Sets:
// Interfaces define sets of methods. A type implements an interface by implementing its methods.
// They are central to Go's type system and polymorphism

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func main() {
	fmt.Println("Interfaces:")
	var r Shape = Rectangle{10, 20}
	fmt.Println("The area of rectangle is ", r.Area())
	fmt.Println("The Perimeter of rectangle is ", r.Perimeter())
	fmt.Println(r)

	var s Shape = Square{10}
	fmt.Println("\nThe area of square is ", s.Area())
	fmt.Println("The Perimeter of square is ", s.Perimeter())
	fmt.Println(s)
}
