package main

import "fmt"

func main() {
	//fmt.Println("Hello, World!")

	// var a int = 10
	// var b int = 15
	// var sum int = a + b

	// fmt.Printf("The sum of two numbers are and type is %T \n ", sum)

	var num int

	fmt.Println("Enter a number:- ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("The entered number", num, "is a Even number")
	} else {
		fmt.Println("The entered number is a Odd Number")
	}

}
