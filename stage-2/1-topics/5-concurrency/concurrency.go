package main

import (
	"fmt"
	"time"
)

func normalFunc(message string) {
	for r := range 5 {
		fmt.Println(message, "", r)
	}
}

func worker(done chan bool) {
	fmt.Println("started worker")
	time.Sleep(time.Second)
	fmt.Println("finished worker")

	done <- true
}

func ping(pings chan<- string, msg string) {
	// only accept sending
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	// pings is a receiver and pongs a sender
	msg := <-pings
	pongs <- msg
}

func main() {
	//calling normal
	normalFunc("normal func")

	//calling with goroutine
	go normalFunc("goroutine func")

	// goroutine with an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("first routine")

	time.Sleep(time.Second)
	fmt.Printf("\nfinished goroutines, starting channels\n\n")

	/*
		IMPORTANT: when making chan T and dont giving a V will not be buffered
		e.g.: make(chan T, V) == buffered // make(chan T) == not buffered
			chan T == bi-directional
			chan<- T == send
			<-chan T == receive
			Trying sending on a receive chan or receive in a sending channel lead to compile-time error
	*/

	message := make(chan string)
	go func() {
		message <- "pong pong"
	}()

	msg := <-message
	fmt.Println(msg)

	bufferedChan := make(chan string, 2)

	bufferedChan <- "ping"
	bufferedChan <- "pong"
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	done := make(chan bool, 1)
	go worker(done)
	<-done // important line! get out the values?!

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "ping msg success")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Printf("\nrange over channels\n\n")

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	close(queue)
	//queue <- "three" // trying to send on closed channel lead to panic

	for msg := range queue { // trying to range over a non-closed channel leads to deadlock
		fmt.Println(msg)
	}

}
