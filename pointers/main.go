package main

import "fmt"

func main() {
	fmt.Println("Starting Pointers")
	var ptr *int // * -data type of pointer storing int values
	fmt.Println("Value of the pointer is ", ptr)

	myNum := 25
	var ptr1 = &myNum // & - reference of the memory
	fmt.Println("Value of the actual pointer is ", ptr1)
	fmt.Println("Value of the actual pointer is ", *ptr1)

	*ptr1 = *ptr1 * 5
	fmt.Println("New Value is ", myNum)
}
