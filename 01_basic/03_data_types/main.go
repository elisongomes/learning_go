package main

import (
	"errors"
	"fmt"
)

func main() {

	var number1 int64 = 10000000000
	fmt.Println(number1)

	var number2 int32 = 1000000000
	fmt.Println(number2)

	// alias
	// int32 = RUNE
	var number3 rune = 1823456
	fmt.Println(number3)

	// uint8 = BYTE
	var number4 byte = 123
	fmt.Println(number4)

	var realNumber1 float32 = 8712.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 8231712.45
	fmt.Println(realNumber2)

	var str string = "Hello Go Lang"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	var boolean1 bool
	fmt.Println(boolean1)

	var error1 error = errors.New("Error message óÒ")
	fmt.Println(error1)
}
