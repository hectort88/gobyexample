package main

import "fmt"

func buffchannel() {
	separator("Channel Buffered")
	messages := make(chan string, 2)

	messages <- "Buffered"
	messages <- "Channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
