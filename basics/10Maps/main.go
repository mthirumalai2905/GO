package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Maps")

	languages := make(map[string]string)

	languages["JS"] = "javasript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println(languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// loops are interesting in golang

	for _, value := range languages {
		fmt.Printf("For , value is %v\n", value)
	}

}
