package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	for _, ele := range getCommands() {
		fmt.Printf("%s: %s\n", ele.name, ele.description)
	}
	return nil
}
