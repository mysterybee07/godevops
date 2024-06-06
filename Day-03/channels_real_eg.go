package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func ProcessOrder(orderID int, StatusChannel chan Order) {
	// simulating some processing time
	time.Sleep(time.Second * 2)

	// sending process order status back through channel
	StatusChannel <- Order{ID: orderID, Status: "Completed"}
}
func main() {
	// Creating a channel for communicating order status
	statusChannel := make(chan Order)

	// starting go routines for processing order
	for i := 1; i <= 5; i++ {
		go ProcessOrder(i, statusChannel)
	}

	// receiving and printing order status back through channel
	for i := 1; i <= 5; i++ {
		order := <-statusChannel
		fmt.Println(order)
	}

}
