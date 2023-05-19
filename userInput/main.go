package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the food rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating,", input)
	fmt.Printf("Variable is of Type %T", input)
}
