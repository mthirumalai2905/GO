package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("In this we are going to make POST request")
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3001/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
	"coursename": "Let's go with golang",
	"price": 0,
	"platform":"learncodeonline.in"
	}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3001/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "thirumalai")
	data.Add("lastname", "pirate")
	data.Add("email", "thiru@go.dev")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

}
