package main

import (
	"fmt"
	"time"
)

func timers() {
	separator("Timers")

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 was fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 was fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Tmer 2 stoped")
	}

	time.Sleep(2 * time.Second)
}
