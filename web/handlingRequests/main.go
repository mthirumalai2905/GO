package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://chatgpt.com/"

func main() {
	fmt.Println("This is about handling web requests")

	// Make the web request
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Ensure the connection is closed

	// Read the response body
	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	// Create and write to the file
	file, err := os.Create("./newbie.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // Ensure the file is closed

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content has been written to newbie.txt")
}
