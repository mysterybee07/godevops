// Message Communication with for Select (daemon)
package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {

	//you can name your for loop
	// running:
	for {
		// block here until someone is sending a message to the channel
		select {
		case msg := <-s.msgch:
			fmt.Printf("received message from : %s \n%s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("The server is doing a graceful shutdown")
			return
			// default:
		}
	}
	// fmt.Println("the server is shutdown")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Biraj Pudasaini",
		Payload: payload,
	}

	msgch <- msg
	// fmt.Println("sending message")
}

func gracefulQuickServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		go sendMessageToServer(s.msgch, "Server Started.......")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		go gracefulQuickServer(s.quitch)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Program exited.......")
	// select {}
}
