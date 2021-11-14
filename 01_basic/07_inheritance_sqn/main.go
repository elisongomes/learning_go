package main

import "fmt"

type People struct {
	name     string
	lastname string
	age      uint8
	height   float32
}

type Student struct {
	People
	course  string
	college string
}

func main() {
	fmt.Println("Inheritance in Go?")

	people1 := People{"Juca", "Tigre", 21, 1.83}
	fmt.Println(people1)

	student1 := Student{people1, "Medicina", "Multivix"}
	fmt.Println(student1)
}
