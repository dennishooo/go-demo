package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main () {
	var favBook = "harry potter"
	var amICool = true
	fmt.Println("hello world")
	fmt.Println(favBook)
	fmt.Println(amICool)

	router := gin.Default()
	router.Run()


}