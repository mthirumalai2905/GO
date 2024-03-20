package main

import "fmt"

func main() {
	fmt.Println("These are pointers")

	// var ptr *int

	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber //Reference means ambersand

	fmt.Println("Value of actual pointer is ", ptr) // gives the memory address
	fmt.Println("Value of actual pointer is", *ptr) // Gives the actual values

	// *ptr = *ptr * 2
	// fmt.Println("New value is: ", myNumber)
	// fmt.Println("Value of actual pointer is", *ptr)

}
