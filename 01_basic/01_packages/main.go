package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliary"
)

func main() {
	fmt.Println("Typing from main file")
	auxiliary.Type1()

	erro := checkmail.ValidateFormat("email@@site.com")
	fmt.Println(erro)
}

// Import package: go get github.com/badoux/checkmail
// Clear packages not used: go mod tidy