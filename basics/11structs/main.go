package main

import "fmt"

func main() {
	fmt.Println("This is about structs")

	//no inheritance in golang
	thiru := User{"Thirumalai", "thirutopg@gmail.com", true, 18}
	fmt.Println(thiru)
	fmt.Printf("thiru's details are: %+v\n", thiru)
	fmt.Printf("Name is %v and email is%v", thiru.Name, thiru.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
