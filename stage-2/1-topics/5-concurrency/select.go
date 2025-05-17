package main

import (
	"fmt"
	"time"
)

func main() {
	// run with "time go run select.go" to get the total time

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- 2
	}()

	for r := range 2 { // looping over the number of routines will lead to deadlock
		r++
		select {
		case msg1 := <-c1:
			fmt.Println("c1 received", msg1) // as soon the second passed will be printed
		case msg2 := <-c2:
			fmt.Println("c2 received", msg2) // will take 3 seconds to be printed
		case msg3 := <-c2:
			fmt.Println("c3 received", msg3) // won't be printed, no more data to be retrieved
		}
	}
}
