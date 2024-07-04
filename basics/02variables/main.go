package main

import "fmt"

//capital letter in go is considered to be public
const LoginToken string = "gibberishvalues"

func main() {
	var username string = "Thirumalai"
	fmt.Println(username + " is topg")
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.56485656
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)

	// implicity type of declaring variables
	var website = "learntodance"
	fmt.Printf("The type of variable is: %T \n", website)
	fmt.Println(website)

	//no var style
	// walrus operator
	numberOfUser := 3000
	fmt.Println(numberOfUser)
}
