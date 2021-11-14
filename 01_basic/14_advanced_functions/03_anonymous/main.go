package main

import "fmt"

func main() {
	fmt.Println("Functions: Anonymous")

	result := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Hello Go Lang")

	fmt.Println(result)
}
