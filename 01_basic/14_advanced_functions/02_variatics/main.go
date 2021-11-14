package main

import "fmt"

func mySum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Only ONE variatic parameter, and on LAST position
func typeText(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println("Functions: Variatics")

	totalSum := mySum(1, 2, 3, 4, 50, 70, 90)
	fmt.Println(totalSum)

	typeText("Hello Go", 1, 2, 4, 5, 6, 7, 9)
}
