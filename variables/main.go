package main

import "fmt"

const Token string = "JDGUYEWU8989766" // Public ( Starting of the letter should be captial)

func main() {
	var username string = "SWIG"
	fmt.Println(username)
	fmt.Printf("Variable is of Type %T \n", username)

	var isBoolean bool = false
	fmt.Println(isBoolean)
	fmt.Printf("Variable is of Type %T \n", isBoolean)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of Type %T \n", smallVal)

	var smallFloat float32 = 4347.47666574
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type %T \n", smallFloat)

	var bigFloat float64 = 4347.476665747346376447
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of Type %T \n", bigFloat)

	var num int
	fmt.Println(num)
	fmt.Printf("Variable is of Type %T \n", num)

	// implicit type
	var website = "www.youtube.com"
	fmt.Println(website)
	fmt.Printf("Variable is of Type %T \n", website)

	//no var style := denotes declaration + assignment ( it can be used inside function not outside of any fn )
	noOfUser := 30000000
	fmt.Println(noOfUser)
	fmt.Printf("Variable is of Type %T \n", noOfUser)

	fmt.Println(Token)
	fmt.Printf("Variable is of Type %T \n", Token)
}
