package main

import "fmt"

// entrypoint for go program
func main() {
	fmt.Println("Welcome to Functions in go lang")
	greeter()
	result := add(3, 5)
	fmt.Println("The result is : ", result)
	proResult, myMsg := proAdder(2, 9, 5, 10, 8, 3, 6, 9)
	fmt.Println("The proResult is : ", proResult)
	fmt.Println("The proResult msg is : ", myMsg)
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	// Assign the result of `add` to a variable.
	plusTwo := addFP(2)
	// Call the function stored in `plusTwo`.
	fmt.Println("First Class Functional Go Programming Addition Logic : ", plusTwo(3)) // Output: 5

	// Define an anonymous function and pass it as an argument to `fmt.Println`.
	fmt.Println(func(x int) int {
		return x * x
	}(2)) // Output: 4

	//Higher-order functions
	// Define a slice of integers.
	numbers := []int{1, 2, 3, 4, 5}
	// Use the `map` function to apply the `square` function to each element of the slice.
	//squares := map(square,numbers)
	fmt.Println(numbers) // Output: [1 4 9 16 25]
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// it can take multiple integer argument and those arg is treated as slice
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro Result function"
}

func greeter() {
	fmt.Println("You Learnt it")
}

// The function `add` takes an integer and returns a function that takes another integer and returns an integer.
func addFP(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// The `square` function returns the square of an integer.
func square(x int) int {
	return x * x
}
