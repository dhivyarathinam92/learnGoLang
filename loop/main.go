package main

import "fmt"

func main() {
	fmt.Println("Loops in Go Lang")
	// slice data structure
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)
	// iterate slice using loop
	for d := 0; d < len(days); d++ {
		fmt.Println("The formal for loop ", days[d])
	}
	// range is used to iterate for array / slice
	for i := range days {
		fmt.Println("The range for loop ", days[i])
	}
	// we can use it with key value pair for Map ( for each loop )
	for index, day := range days {
		fmt.Printf("index is %v and the value is %v\n", index, day)
	}

	rougueValue := 1
	// while loop
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lgc
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		} else if rougueValue == 9 {
			break
		}
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}
lgc:
	fmt.Println("Jumping at learn Go code")

}
