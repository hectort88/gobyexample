package main

import "fmt"

func nonblockingchannel() {
	separator("Non-Blocking Channel Operations")

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("No received messages")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent Message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("No activity")
	}
}
