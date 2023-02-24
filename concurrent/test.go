package main

import "fmt"
func main () {
	// unbuffered := make(chan int)
	// go func() {<-unbuffered}()

	// unbuffered <- 123

	// a:= <-unbuffered

	// fmt.Println(a)
	c := make(chan int)

	close(c)
	fmt.Println(<-c)
}