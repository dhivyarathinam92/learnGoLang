package main

import "fmt"

func main() {
	// defer keyword follows LIFO ie reverse chronological order
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hey")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
