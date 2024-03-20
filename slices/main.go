package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("These are Slices")

	var fruitlist = []string{"Apple", "tomato", "peach"}

	fmt.Printf("The type of fruitlist is %T \n", fruitlist)

	fmt.Println(fruitlist)
	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 224
	highScores[2] = 254
	highScores[3] = 244
	// highScores[4] = 2434 Index out of bounds

	highScores = append(highScores, 55, 666, 45)
	fmt.Println(highScores)
	fmt.Printf("The type is %T \n", highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from the slices based on the index

	var courses = []string{"reactjs", "javascript", "swift", "python", "Ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
