package main

import (
	"errors"
	"fmt"

	"github.com/Tubelight30/my_pokedex/internal/pokeapi"
)

func commandExplore(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No location area provided")
	}
	locationAreaName := args[0]
	resp, err := pokeapi.ExploreArea(locationAreaName)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	fmt.Printf("Pokemon in %s:\n", resp.Name)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
