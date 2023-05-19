package main

import (
	"fmt"
	"sort"
)

func main() {

	var langList = []string{"Java", "GO", "Python"}
	fmt.Println(" Actual slice : ", langList)

	langList = append(langList, "C", "C++")
	fmt.Println("After appending1 : ", langList)

	langList = append(langList[1:])
	fmt.Println("After appending2 : ", langList)

	langList = append(langList[1:3]) // last value is non inclusive
	fmt.Println("After appending3 : ", langList)

	langList = append(langList[:3]) // last value is non inclusive
	fmt.Println("After appending4 : ", langList)

	highScores := make([]int, 4) // defined the size of slice ie 4
	highScores[0] = 123
	highScores[1] = 678
	highScores[2] = 387
	highScores[3] = 971
	fmt.Println("Slice with mem mgnt : ", highScores)

	highScores = append(highScores, 555, 289, 777) // relocate the memory according to the values
	fmt.Println("Slice with append : ", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted int with increasing order : ", highScores)
	fmt.Println("Boolean wether sorted or not : ", sort.IntsAreSorted(highScores))

	// remove value from slice based on index
	var courses = []string{"reactJs", "javaScript", "swift", "typeScript", "flutter"}
	fmt.Println("The actual course slice is : ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removing course slice is : ", courses)

}
