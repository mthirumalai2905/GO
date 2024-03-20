package main

import (
	"fmt"
)

func main() {
	fmt.Println("These are maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "rb")
	fmt.Println(languages)

	// loops are intresting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)

	}

}
