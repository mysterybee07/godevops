# How to Handle Concurrency in Go?
Concurrency is handled primarily through two key features: goroutines and channels. Understanding crucial writing efficient and effective Go programs:

# Goroutines
   go myFunction()
   To create a goroutine, you simply use the go keyword followed by a function call. This will start a new goroutine that will execute myFunction() concurrently with the rest of the program.

# Features of Goroutines:
   a. Independence: Each goroutine runs independently.
   b. Non-blocking: The go statement itself is non-blocking. The program will not wait for the goroutine to complete; instead, it will continue to the next line of code.
   c. Shared Memory: Goroutines within the same process share the same memory space, so care must be taken to synchronize the shared values.

# Channels
channels are the conduits through which goroutines communicate with each other. They allow the safe passing of data between goroutines, ensuring that only one goroutine accesses a particular piece of data ata a time.

1. Creating a Channel
   ch := make(chan int)
   You create a <- channel using the make function, above creates a channel for passing integer creates a channel for passing integers.

2. Sending and Receiving Channels
   Sending: ch <- v sends v to channel ch.
   Receiving: v:= <-ch receives data from ch and assigns it to v.

Features of Channels:
1. Synchronization: Sending and receiving operations on a channel are blocking by default. The send operation blocks until another goroutine reads from the channel, and vice versa.
2. Buffering: Channels can be bufferred. A buffered channel has a capacity and doesnot block until the buffer is full(for sends) and empty (for receives)
3. Directional: Channels can be directional, either send-only or receive-only, enhancing type safety.

# Basic Example

package main

func printCounts(label string, count int, ch chan int) {
    for i:=0; i<count; i++{
        ch <- i
        fmt.Println(label,i)
        time.Sleep(time.Millisecond*500)//simulate work
    }
    close(ch)
}

func main(){
    ch := make(chan int)
    go printCounts("goroutine", 5,ch)

    for value:= range ch{
        fmt.Println("Main received:",value)
    }
}

In this example, printCounts is a function that runs in a goroutine and communicates with the main function via a channel. The main function waits and processes messages from the channel until its closed.