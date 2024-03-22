package main

import "fmt"

func main() {
	fmt.Println("This is my function")
	greeter()

	result := adder(3, 7)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 3, 4, 5, 6, 6)
	fmt.Println("The  Result is: ", proResult)
	fmt.Println("The  Result is: ", myMessage)

}

/// ... triple dots are variadic function
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello"
}

func adder(a int, b int) int {
	return a + b
}

func greeter() {
	fmt.Println("Namstey from golang")
}
