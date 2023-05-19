package main

import "fmt"

func main() {
	// Equivalent to class but in GO , dont have inheritance ie ( no parent / super class )
	fmt.Println("Structs in GoLang")
	sampleUser := User{"Sample", "sample@mail.com", true, 32}
	fmt.Println(sampleUser)
	fmt.Printf("User details are : %+v\n", sampleUser)
	fmt.Printf("Name is %v and Email is %v\n", sampleUser.Name, sampleUser.Email)
	sampleUser.GetStatus()
	sampleUser.GetNewMail()
	fmt.Printf("Email is %v\n", sampleUser.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// function which take struct as input known as methods
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) GetNewMail() {
	u.Email = "test@mail.com"
	fmt.Println("Email of the user is :", u.Email)
}
