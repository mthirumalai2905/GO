package main

import "fmt"

func main() {
	greeter()
	fmt.Println("This is about function")
	greeterTwo()

	result := add(5, 10)
	fmt.Println("Result is:", result)

	Biggie := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(Biggie)
}

func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Hello from golang")
}

func greeterTwo() {
	fmt.Println("Hello from golang but the second one")
}
