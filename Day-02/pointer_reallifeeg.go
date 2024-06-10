package main

import "fmt"

// represents a record
//
//	type Employee struct {
//		Name     string
//		Position string
//		Salary   float64
//	}
type Student struct {
	Name   string
	RollNo string
	Fee    float64
}

func UpdateFee(fee *Student, newFee float64) {
	fee.Fee = newFee
}

// func UpdateSalary(emp *Employee, newSalary float64) {
// 	fmt.Println(emp)
// 	emp.Salary = newSalary
// }

//emp is a pointer variable variable that points to the memory adress of employee

func main() {
	// emp := Employee{
	// 	Name:     "Biraj",
	// 	Position: "Software Engineer",
	// 	Salary:   100000,
	// }
	student := Student{
		Name:   "Biraj",
		RollNo: "123",
		Fee:    100000,
	}
	// fmt.Println("Initial Salary: ", emp.Salary)
	fmt.Println(student.Fee)
	UpdateFee(&student, 2000)
	// fmt.Println(emp)
	// UpdateSalary(&emp, 200000)
	// fmt.Println("Updated Salary: ", emp)
	fmt.Println(student.Fee)

}
