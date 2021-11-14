package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 7
	array1[1] = 14

	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int{2, 4, 6, 8}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	// Insert an new element to Slice and return new slice
	slice1 = append(slice1, 12)
	fmt.Println(slice1)

	// Pointer to an array (1:3 => Pos 1, Pos 2)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[2] = "Position changed"
	fmt.Println(slice2)

	// Internal array
	fmt.Println("---------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	// Slices increase the length and capacity dynamic
	slice3 = append(slice3, 2)
	slice3 = append(slice3, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity
}
