package main

import "fmt"

// Goroutine Communication
// channels are used for communication between goroutines (concurrent threads).
// The can be synchronous and asynchronous

func main() {
	fmt.Println("Channels")
	channels := make(chan int) //unbuffered channels of integers
	fmt.Println("channels:", channels)
}
