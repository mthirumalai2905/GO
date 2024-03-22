package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNummber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNummber)

	switch diceNummber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move two spots")
	case 3:
		fmt.Println("you can move three spots")
		fallthrough //this will print the current 3 and 4 also
	case 4:
		fmt.Println("You can move four spots")
	case 5:
		fmt.Println("You can move five spots")
	case 6:
		fmt.Println("You can move six spots")

	default:
		fmt.Println("What was that")
	}
}
