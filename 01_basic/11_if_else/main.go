package main

import "fmt"

func mySum(a, b int) int {
	return a + b
}
func main()  {
	fmt.Println("Control structures: If/Else")

	// If init; Variable only in IF scope
	if sum1 := mySum(2, 3); (sum1 % 2) == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
