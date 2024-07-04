package main

import "fmt"

func main() {
	fmt.Println("This is about pointers")

	//Pointers give us the assurance

	// var ptr *int //The default value of the pointer is <nil>

	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value", myNumber)

}
