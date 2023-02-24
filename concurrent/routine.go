package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i:=0;i<10;i++ {
	time.Sleep(time.Millisecond * 500)

		fmt.Println(from, ":", i)
	}
}

func Routine(){
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("wtf is going on")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	// fmt.Println(time.Month.String(1))
}

// goroutine is a function that is capable of runing concurrently with other function
// it is a lightweight thread of execution
// Goroutines run concurrently, but not necessarily in parallel. 


