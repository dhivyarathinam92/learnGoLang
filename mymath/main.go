package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	/*var mynum1 int = 2
	var mynum2 float64 = 4*/

	//fmt.Println("the sum is:", mynum1+int(mynum2))

	//random number
	/*rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)*/

	//random from crypto
	myRandomNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNo)
}
