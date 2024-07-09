package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := c.pokemonCaught[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf(" -%s\n", types.Type.Name)

	}

	return nil
}
