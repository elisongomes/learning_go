package main

import "fmt"

func myCalc(n1, n2 int) (retSum, retSub int) {
	retSum = n1 + n2
	retSub = n1 - n2

	// No need specify the variables, because it's named on function
	return
}

func main()  {
	fmt.Println("Functions: Named return")
	sum1, sub1 := myCalc(10, 20)
	fmt.Println(sum1, sub1)
}
