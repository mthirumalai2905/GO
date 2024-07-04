package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about slices")

	// var fruitlist = []string{"apple", "tomato", "peach"}
	// fmt.Printf("Type of fruitlist is: %T \n", fruitlist)

	// fruitlist = append(fruitlist, "Mango", "Banana")
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:3])
	// fmt.Println(fruitlist)

	// highScores := make([]int, 4)

	// highScores[0] = 2347
	// highScores[1] = 2650
	// highScores[2] = 5659
	// highScores[3] = 5432

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)
	// fmt.Println(len(highScores))

	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// How to remove elements from the slices based on indexes

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
