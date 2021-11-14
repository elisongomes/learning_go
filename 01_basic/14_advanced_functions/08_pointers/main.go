package main

import "fmt"

func signalInverter(num *int) {
	*num = *num * -1
}

func main() {
	num1 := 20
	fmt.Println(num1)
	signalInverter(&num1)
	fmt.Println(num1)
}
