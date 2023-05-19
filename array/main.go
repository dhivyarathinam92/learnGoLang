package main

import "fmt"

func main() {
	fmt.Println("Array Class")
	var sampleList [4]string
	sampleList[0] = "ABCDE"
	sampleList[1] = "EFGHI"
	sampleList[3] = "OPQRS"

	fmt.Println("Sample list is : ", sampleList)
	fmt.Println("Sample list is : ", len(sampleList))

	var langList = [3]string{"Java", "GO", "Python"}
	fmt.Println("Sample list is : ", langList)
	fmt.Println("Sample list is : ", len(langList))
}
