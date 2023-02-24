package main

import "fmt"

// channels are pipes that connect concurrent goroutines
// Values can be sent into channels from one goroutine and receive those values into another goroutine.

// By default channels are unbuffered,
// meaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready to receive the sent value.

// Using a channel, the main task will wait until the asynchronous task is complete.
// It blocks until it receives a notification from the worker on the channel.

// When the goroutine completes its work, it will send a value through the channel,
// which will be read before operating on the numbers array.

// Channels can be buffered if the intention is to prevent blocking further execution
// until a value is eventually read from the channel to free it up.

func sum(s []int, c chan int, name string) {
	sum :=0
	for _, v := range s {
	fmt.Println("doing", name, v)

		sum += v
	}
	c <- sum
}

func Channel() {
	s := []int{1, 2, 3, 4}

	c := make(chan int)

	go sum(s[:len(s)/2], c, "first half")
	go sum(s[len(s)/2:], c, "second half")
	x, y := <-c, <-c

	fmt.Println(x, y)
}