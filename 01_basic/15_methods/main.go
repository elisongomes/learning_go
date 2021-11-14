package main

import "fmt"

type User struct {
	name string
	age  uint8
	active bool
}

func (u User) save() {
	fmt.Printf("Saving user data for: %s\n", u.name)
}

func (u User) isAdult() bool {
	return u.age >= 18
}

func (u *User) activate() {
	u.active = true;
}

func main() {
	fmt.Println("Methods")
	user1 := User{"Juca Tigre", 21, false}
	fmt.Println(user1)
	user1.save()
	fmt.Println(user1.isAdult())

	user1.activate()
	fmt.Println(user1)
}
