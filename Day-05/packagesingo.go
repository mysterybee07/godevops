package main

import (
	"fmt"

	"github.com/mysterybee07/godevops/Day-05/shape"
)

// What is Package: Packages in Go are a way to organize and encapsulate code.
// Each package can contain multiple Go files and each belong to one package.

// Creating a package: To create a package, declare it at the top of your go file

// Exported Names: In Go, a name is exported if it begins with a capital letter.
// For example Printer is an exported name and can be used from other packages. while printer is not

// Importing Packages: You can use code from other packages by impoting them

// Organizing Code with packages:
// Packages are a great way to structure your Go application. Each pacjage should have a single purpose.
// For instance a net/http package in the standard library is used for Gttp request and resposes.

// Dependency Management:
// Go Modules are used for managing dependencies.
// You can specify the required libraries in you go.mod file and Go will handle the donloading and updating of these libraries
func main() {

	fmt.Println(shape.SArea(4))
	fmt.Println(shape.SPerimeter(4))

	fmt.Println(shape.Area(4))
	fmt.Println(shape.Perimeter(4))

}
