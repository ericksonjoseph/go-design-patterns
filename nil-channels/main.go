package main

import (
	"fmt"
	"time"
)

func main() {

	var isClosed1, isClosed2 bool

	exit := make(chan bool)
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	// Close channel 1 after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closing 1")
		close(ch1)
	})

	// Close channel 2 after 6 seconds
	time.AfterFunc(6*time.Second, func() {
		fmt.Println("closing 2")
		close(ch2)
	})

	// Exit after 8 seconds
	time.AfterFunc(8*time.Second, func() {
		fmt.Println("time")
		exit <- true
	})

	for !isClosed1 || !isClosed2 {

		fmt.Println(".")

		// A closed channel will always be selected
		// So we need to make the channel nil to be ignored
		select {
		case i, open := <-ch1:
			fmt.Println("ch1 received. i:", i, "open:", open)
			if !open {
				// If you comment the following line, you'll see that the
				// second case may never be called
				ch1 = nil
				isClosed1 = true
			}
		case i, open := <-ch2:
			fmt.Println("ch2 received. i:", i, "open:", open)
			if !open {
				ch2 = nil
				isClosed2 = true
			}
		}
	}

	fmt.Println("loop over")

	var a bool

	a = <-exit

	fmt.Println("exited", a)
}
