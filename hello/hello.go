package main

import (
	"fmt"

	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main()  {
	log.SetPrefix("greetings: ")
    // log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("hello world")
	fmt.Println(quote.Go())
	fmt.Println(message)

}