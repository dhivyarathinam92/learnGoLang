package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	PerformGetRequest()
	PerformPostJsonReq()
	PerformPostFormReq()
}

func PerformPostFormReq() {
	const myurl = "http://localhost:9000/postform"
	//form data
	data := url.Values{}
	data.Add("firstname", "first")
	data.Add("lastname", "last")
	data.Add("email", "test@email.com")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJsonReq() {
	const myurl = "http://localhost:9000/post"
	// creating a json req data
	reqBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		}
	`)
	response, err := http.Post(myurl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformGetRequest() {
	const myurl = "http://localhost:9000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length is ", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}
