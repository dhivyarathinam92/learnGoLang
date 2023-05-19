package main

import "fmt"

func main() {
	// Equivalent to class but in GO , dont have inheritance ie ( no parent / super class )
	fmt.Println("Structs in GoLang")
	sampleUser := User{"Sample", "sample@mail.com", true, 32}
	fmt.Println(sampleUser)
	fmt.Printf("User details are : %+v\n", sampleUser)
	fmt.Printf("Name is %v and Email is %v", sampleUser.Name, sampleUser.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
