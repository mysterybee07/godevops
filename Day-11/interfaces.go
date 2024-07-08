package main

import "fmt"

// Single Method Interface
type Speaker interface {
	speak() string
}

type Dog struct {
	Name string
}

func (d Dog) speak() string {
	return d.Name + "says Woof!"
}

// Embedding Interface
type Animal interface {
	Speaker
	Move() string
}
type Bird struct {
	Name string
}

func (b Bird) speak() string {
	return b.Name + "says Chirp!"
}
func (b Bird) move() string {
	return b.Name + " is flying!"
}

// Empty Interface
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

// Interface with multiple methods
type Vehicle interface {
	Start() string
	Stop() string
}
type Car struct {
	Model string
}

func (c Car) start() string {
	return c.Model + " is starting!"
}
func (c Car) stop() string {
	return c.Model + " is stopping!"
}

// Interface as a Constraint Generic
func Describe[T Speaker](t T) {
	fmt.Println(t.speak())
}
func main() {
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweety"}
	car := Car{Model: "Toyota"}
	fmt.Println(dog.speak())
	fmt.Println(bird.speak())
	fmt.Println(bird.move())

	PrintAnything("An empty Interface can hold anything")
	PrintAnything(42)

	fmt.Println(car.start())
	fmt.Println(car.stop())

	Describe(dog)
	Describe(bird)
}
