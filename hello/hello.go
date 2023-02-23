package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main()  {
	fmt.Println("hello world")
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("dennis"))

}