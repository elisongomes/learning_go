package main

import "fmt"

func closure() func() {
	text := "Inside closure function"

	func1 := func() {
		fmt.Println(text)
	}
	return func1
}
func main() {
	fmt.Println("Functions: Closure")
	text := "Inside main function"
	fmt.Println(text)

	funcClosure := closure()
	funcClosure()
}
