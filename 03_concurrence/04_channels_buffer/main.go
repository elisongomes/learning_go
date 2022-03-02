package main

import (
	"fmt"
)

func main() {

	//channel with capacity = 2
	ch := make(chan string, 2)

	ch <- "Hello world"
	ch <- "World is good"

	message1 := <-ch
	message2 := <-ch
	fmt.Println(message1)
	fmt.Println(message2)
}
