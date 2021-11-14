package main

import "fmt"

func main() {
	var variable1 string = "Variable 1"

	// The type of variable is defined by value = Type inference
	variable2 := "Variable 2"

	fmt.Println(variable1)
	fmt.Println(variable2)

	// Multiple variables declaration
	var (
		variable3 string = "Variable 3"
		variable4 string = "Variable 4"
	)

	fmt.Println(variable3)
	fmt.Println(variable4)

	variable5, variable6 := "Variable 5", "Variable 6"
	fmt.Println(variable5, variable6)

	const constant1 string = "Constant 1"
	fmt.Println(constant1)

	//Trocar valor das variáveis
	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}
