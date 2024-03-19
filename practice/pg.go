package main

import "fmt"

func main() {
	// Knowing how to separate arguments in function calls fix
	// the following line so that it prints "go 2 true"
	var number int = 2
	var isLoggedIn bool = true
	fmt.Println("go", number, isLoggedIn)
}
