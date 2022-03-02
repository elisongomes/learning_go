package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			ch1 <- "Message to channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "Message to channel 2"
		}
	}()

	/*
	*
	* Without `select` has an block until receive each message on channel
	*
	**/
	//for {
	//	message1 := <-ch1
	//	fmt.Println(message1)
	//
	//	message2 := <-ch2
	//	fmt.Println(message2)
	//}

	for {
		select {
		case message1 := <-ch1:
			fmt.Println(message1)
		case message2 := <-ch2:
			fmt.Println(message2)
		}
	}
}
