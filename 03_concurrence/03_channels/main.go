package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go writeText("Hello world!", ch1)
	go writeText("Go lang is funny!", ch2)

	// Example 1: channel iterate with manual check if is open
	for {
		//receive value on channel
		message, openned := <-ch1
		//verify if channel is still openned
		if !openned {
			break
		}
		fmt.Println(message)
	}

	// Example 2: channel iterate with range, automatic check if is open
	for message := range ch2 {
		fmt.Println(message)
	}

}

func writeText(text string, ch chan string) {
	for i := 0; i < 5; i++ {
		//send value to channel
		ch <- fmt.Sprintf("%s - %d", text, i)
		time.Sleep(time.Second)
	}
	// close channel to prevent deadlock
	close(ch)
}
