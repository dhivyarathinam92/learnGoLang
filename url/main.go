package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lgo.dev:3000/learn?coursename=reactjs&paymentid=jahsfgdjadfg"

func main() {
	fmt.Println("Handling Url in GoLang")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query param are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lgo.dev",
		Path:    "/tutcss",
		RawPath: "user=sample",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
