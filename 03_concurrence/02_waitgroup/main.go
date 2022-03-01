package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// goroutines that will be run at same time
	waitGroup.Add(2)

	go func() {
		go writeText("Hello world")

		waitGroup.Done()
	}()

	go func() {
		writeText("World is good")

		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func writeText(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
