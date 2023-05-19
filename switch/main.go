package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in GoLang")
	rand.Seed(time.Now().UnixNano())
	diceNo := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNo)

	switch diceNo {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough // it will run the next case as well
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll dice again")
		fallthrough
	default:
		fmt.Println("What was that")
	}
}
