package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variable1 int = 0
	var variable2 int = variable1

	variable1++

	fmt.Println(variable1, variable2)

	var variable3 int
	var pointer1 *int

	variable3 = 100
	pointer1 = &variable3

	fmt.Println(variable3, pointer1)

	variable3 = 150

	// The "*" is for dereferencing
	fmt.Println(variable3, pointer1, *pointer1)
}
