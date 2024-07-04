package main

import "fmt"

func main() {
	fmt.Println("This is about control flow statements")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Good User"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

	var logic = []string{}

	fmt.Println(logic)

	var name string
	fmt.Println(name)

}
