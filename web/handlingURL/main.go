package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fsdhsdhbav582"

func main() {
	fmt.Println("This is about handling URLs")
	fmt.Println(myurl)

	//parsing
	res, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params are: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=thiru",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
