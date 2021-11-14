package main

import "fmt"

func func1() {
	fmt.Println("Running function 1")
}

func func2() {
	fmt.Println("Running function 2")
}

func studentApproved(n1, n2 float32) bool {
	defer fmt.Println("Calculated average. Result will be returned.")
	fmt.Println("Entering on funcion for calc")

	median := (n1 + n2) / 2
	if median >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Functions: Defer")
	// Defer execute on last moment of method
	defer func1()
	func2()

	fmt.Println(studentApproved(7, 8))
}
