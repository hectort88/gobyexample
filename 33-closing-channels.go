package main

import "fmt"

func closingchannels() {
	separator("Closing Channels")

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("Received all the jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs
	for j := 1; j < 4; j++ {
		jobs <- j
		fmt.Println("Sent job:", j)
	}

	// close the jobs channel
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
