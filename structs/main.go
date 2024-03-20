package main

import "fmt"

func main() {
	fmt.Println("This will be about structs")

	//no inheritance in golang;
	//no super or no parent

	thiru := User{
		"Thiru",
		"thiru@go.dev",
		true,
		18}

	fmt.Println(thiru)
	fmt.Printf("Thiru details are: %+v\n", thiru)
	fmt.Printf("Name is %v and email is %v\n", thiru.Name, thiru.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
