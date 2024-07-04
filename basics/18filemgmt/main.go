package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is about file management in golang")

	content := "This needs to go in a file - Mugiwara no ichi"

	file, err := os.Create("./myluffygo.txt")

	if err != nil {
		panic(err)
	}

	io.WriteString(file, content)
	defer file.Close()

	readFile("./myluffygo.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
