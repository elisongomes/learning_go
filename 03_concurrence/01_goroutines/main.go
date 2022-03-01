package main

import (
	"fmt"
	"time"
)

func main() {
	go writeText("Hello world")
	writeText("World is good")
}

func writeText(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
