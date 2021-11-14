package main

import "fmt"

var n int

func init()  {
	// Can be used for file bootstrap, setup, etc..
	fmt.Println("Init funcion running")
	n = 10
}

func main() {
	fmt.Println("Main function running")
	fmt.Println(n)
}
