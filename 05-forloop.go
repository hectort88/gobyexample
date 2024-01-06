package main

import "fmt"

func forloops() {
	// for like while
	separator("For Loops")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// classic c for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	counter := 0

	for {
		if counter == 3 {
			break
		}
		fmt.Println("loop", counter)
		counter++
	}

	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
