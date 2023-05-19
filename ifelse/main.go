package main

import "fmt"

func main() {
	fmt.Println("Control Flow Statements in Golang")

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("No is even")
	} else {
		fmt.Println("No is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("No is less than 10")
	} else {
		fmt.Println("No is greater than 10")
	}
}
