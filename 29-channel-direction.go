package main

import "fmt"

func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channeldir() {
	separator("Channel direction")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
