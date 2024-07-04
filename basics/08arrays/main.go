package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Golang")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))

	var veggies = [3]string{"potato", "beans", "brocolli"}
	fmt.Println("Vegy list is:", veggies)
}
