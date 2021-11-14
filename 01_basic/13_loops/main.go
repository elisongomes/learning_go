package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	//Not have while, do while..

	i := 0
	for i < 5 {
		i++
		fmt.Printf("Increasing I %d\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 5; j++ {
		fmt.Printf("Increasing J %d\n", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Juca", "Maria", "joão"}

	for key, value := range names {
		fmt.Println(key, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	for key, value := range "WORD" {
		fmt.Println(key, string(value))
	}

	user1 := map[string]string{
		"name":    "Juca",
		"surname": "Tigre",
	}
	for key, value := range user1 {
		fmt.Println(key, value)
	}

	r := 0
	for {
		fmt.Printf("Infinity loop\n")
		time.Sleep(time.Second)
		r++
		if r > 5 {
			break
		}
	}
}
