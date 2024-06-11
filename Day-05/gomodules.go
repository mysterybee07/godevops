package main

import (
	"fmt"

	"github.com/google/uuid"
)

// What is Go Module?
// Go Modules is the official dependency management system in Go, introduced in version 1.11.
// It replaces the older GOPATH-based approach to managing dependencies and allows for more straightforward management of project dependencies

// Initializing a New Module
// To start a new project with Go Modules, you first initialize a new module.
// This is done using the go mod init command followed by your modules name(often the repository path).

func main() {
	id := uuid.New()
	fmt.Println(id)
}

// The go.sum File
// It contains the cryptographic checksums of the content of specific version of modules.
// It ensures the integrity and authenticity of your modules dependencies
