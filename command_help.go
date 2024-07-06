package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	// fmt.Println("help: Displays a help message")
	// fmt.Println("exit: Exit the Pokedex")
	for _, ele := range getCommands() {
		fmt.Printf("%s: %s\n", ele.name, ele.description)
	}
	return nil
}