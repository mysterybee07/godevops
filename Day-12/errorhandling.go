package main

import (
	"errors"
	"fmt"
)

// custom error is a custom error type
type CustomError struct {
	msg string
}

// Error implements the error interface for customError
func (e *CustomError) Error() string {
	return e.msg
}

// SomeFunction demonstrates returning a custom error
func SomeFunction(flag bool) error {
	if !flag {
		return &CustomError{"Custom error occuring"}
	}
	return nil
}

// SafeFunction demonstrates using panic and recover
func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	panic("A problem occurred")
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}
func main() {
	// Simple error return
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// another example with valid division

	result1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	// Example checking and handling error with custom errors
	err2 := SomeFunction(false)
	if err2 != nil {
		fmt.Println(err2)
	}

	// Example using panic and recover
	SafeFunction()
}
