package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("This is for web requests")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close() // Close the response body when main() exits

	fmt.Printf("Response type: %T\n", response)

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	content := string(databytes)
	fmt.Println("Response body content:\n", content)
}
