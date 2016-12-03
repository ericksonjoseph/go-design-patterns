package main

import (
	"fmt"
	"time"
)

var (
	shutdown = make(chan struct{})
	done     = make(chan int)
)

func main() {
	const n = 5

	// Start up the goroutines...
	for i := 0; i < n; i++ {
		i := i
		go func() {
			select {
			case <-shutdown:
				done <- i
			}
		}()
	}

	panic("stacks")

	time.Sleep(1 * time.Second)

	// Close the channel. All goroutines will immediately "unblock".
	close(shutdown)

	for i := 0; i < n; i++ {
		fmt.Println("routine", <-done, "has exited!")
	}
}
