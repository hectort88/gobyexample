// A goroutine is a lightweight thread of execution.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutines() {
	separator("Go Routines")

	// running synchronously.
	f("direct")

	// This will execute concurrently with the calling one.
	go f("goroutine1")
	go f("goroutine2")
	go f("goroutine3")

	// start a goroutine for an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
