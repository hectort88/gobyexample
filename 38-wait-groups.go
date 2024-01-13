package main

import (
	"fmt"
	"sync"
	"time"
)

func workerwg(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep((time.Second))
	fmt.Printf("Worker %d done\n", id)
}

func waitgroup() {
	separator("Wait Group")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			workerwg(i)
		}()
	}

	wg.Wait()
}
