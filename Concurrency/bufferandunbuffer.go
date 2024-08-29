package main

import "fmt"

func main() {
	msgCh := make(chan int, 1)
	msgCh <- 10

	msg := <-msgCh
	fmt.Println(msg)

}
