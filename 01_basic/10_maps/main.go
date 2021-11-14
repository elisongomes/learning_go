package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Key => Value
	user1 := map[string]string{
		"first":     "Pedro",
		"last_name": "Silva",
	}
	fmt.Println(user1)

	user2 := map[string]map[string]string{
		"name": {
			"first": "Juca",
			"last":  "Tigre",
		},
		"course": {
			"name":    "Medicina",
			"college": "Multivix",
		},
	}
	fmt.Println(user2)

	delete(user2, "course")
	fmt.Println(user2)

	user2["family"] = map[string]string{
		"father":  "João",
		"monther": "Maria",
	}
	fmt.Println(user2)
}
