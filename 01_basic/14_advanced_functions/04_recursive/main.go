package main

import "fmt"

// Addeded memoization
func fibonacci(position uint, memo map[uint]uint) uint {
	if position < 2 {
		return position
	}
	if _, memorize := memo[position]; !memorize {
		memo[position] = fibonacci(position-2, memo) + fibonacci(position-1, memo)
	}
	return memo[position]
}

func main() {
	fmt.Println("Functions: Recursive")

	// 1 1 2 3 5 8 13
	position := uint(50)
	fmt.Println(fibonacci(position, map[uint]uint{}))

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i, map[uint]uint{}))
	}
}
