package main

import "fmt"

func mySum(n1, n2 int8) int8 {
	return n1 + n2
}

func myCalc(n1, n2 int8) (int8, int8) {
	sum1 := n1 + n2
	sub1 := n1 - n2
	return sum1, sub1

}
func main() {
	sum1 := mySum(10, 20)
	fmt.Println(sum1)

	var f = func(text string) string {
		fmt.Println(text)
		return text
	}

	fmt.Println(f("Test message.."))

	resultSum1, resultSub1 := myCalc(10, 8)
	fmt.Println(resultSum1, resultSub1)

	// The "_" ignore any return value from function
	resultSum2, _ := myCalc(10, 8)
	fmt.Println(resultSum2)
}
