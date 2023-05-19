package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of Map ", languages)
	fmt.Println("Print data by providing a key ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of Map after deleting ", languages)

	// Iterate in Map
	for key, val := range languages {
		fmt.Println("The key Value Pair is", key, val)
	}

	for _, val := range languages {
		fmt.Println("The value is ", val)
	}
}
