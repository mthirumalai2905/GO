package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Let's make our own server in golang")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3001/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(content)
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
}
