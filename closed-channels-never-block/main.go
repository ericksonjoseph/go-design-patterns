package main

import (
	"fmt"
	"time"
)

func main() {

	const n = 5

	lifeline := make(chan struct{})
	done := make(chan int)

	for i := 0; i < n; i++ {

		i := i
		fmt.Println("launching ", i)

		go func() {

			select {

			case <-lifeline:
				fmt.Printf("Shutting %d down\n", i)
				done <- i
			}
		}()
	}

	time.Sleep(2 * time.Second)

	close(lifeline)

	for i := 0; i < n; i++ {
		fmt.Printf("Job %d had completed!\n", <-done)
	}

	fmt.Println("exiting")
}
