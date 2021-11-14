package main

import "fmt"

type User struct {
	name    string
	age     uint8
	address Address
}

type Address struct {
	location string
	number   uint8
}

func main() {
	fmt.Println("Structs file")

	var user1 User
	fmt.Println(user1)

	user1.name = "Maria"
	user1.age = 25
	fmt.Println(user1)

	user2 := User{"Pedro", 22, Address{"Rua Exemplo", 17}}
	fmt.Println(user2)

	user3 := User{name: "João"}
	fmt.Println(user3)
}
