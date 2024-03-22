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
	thiru.GetStatus()
	thiru.NewMail()
	fmt.Println(thiru.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

//Copy of the function has been passed out
//if u want to change the value u nee to use the pointers
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
