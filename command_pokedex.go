package main

import "fmt"

func commandPokedex(c *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokemonCaught {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
