package main

import "fmt"

// represents a record
type Employee struct {
	Name     string
	Position string
	Salary   float64
}

func UpdateSalary(emp *Employee, newSalary float64) {
	fmt.Println(emp)
	emp.Salary = newSalary
}

//emp is a pointer variable variable that points to the memory adress of employee

func main() {
	emp := Employee{
		Name:     "Biraj",
		Position: "Software Engineer",
		Salary:   100000,
	}

	fmt.Println("Initial Salary: ", emp.Salary)

	// fmt.Println(emp)
	UpdateSalary(&emp, 200000)
	fmt.Println("Updated Salary: ", emp)

}
