package main

import "fmt"

func panicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered execution")
	}
}

func studentApproved(n1, n2 float32) bool {
	defer panicRecover()
	median := (n1 + n2) / 2

	if median > 6 {
		return true
	} else if median < 6 {
		return false
	}

	panic("The median is equal 6!")

}

func main() {
	fmt.Println("Functions: Panic and Recover")
	fmt.Println(studentApproved(6, 6))
}
