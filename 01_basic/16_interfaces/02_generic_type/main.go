package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Printf("Interfaces: generic")
	generic("String")
	generic(1)
	generic(true)

	map1 := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     10,
		true:         float64(12),
	}

	fmt.Println(map1)

}
