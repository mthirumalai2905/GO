package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This are files")
	content := "This needs to go in a file - LearnGo.in"

	file, err := os.Create("./mytxtgofile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // Close the file at the end of main function

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length is: ", length)

	readFile() // Call the readFile function to read from the file
}

func readFile() {
	fileName := "./mytxtgofile.txt" // Define the filename
	databyte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}
