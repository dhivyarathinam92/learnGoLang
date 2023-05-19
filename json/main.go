package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lgoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}
	//pkg this data as json data
	finalJson, err := json.MarshalIndent(lgoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename":"ReactJs BootCamp",
		"Price": 299,
		"website":"www.youtube.com",
		"tags":["open-src","sr"]
	}
	`)
	var lCourse course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lCourse)
		fmt.Printf("%#v\n", lCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// add data to key value pair
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Println("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key  is %v and value is %v "+
			"and Type is : %T\n", k, v, v)
	}
}

func main() {
	EncodeJson()
	DecodeJson()
}
