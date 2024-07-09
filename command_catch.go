package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Tubelight30/my_pokedex/internal/pokeapi"
)

func commandCatch(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	resp, err := pokeapi.GetPokemon(pokemonName)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	threshold := 50
	randNum := rand.Intn(resp.BaseExperience)
	fmt.Println(resp.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s", pokemonName)
	}

	c.pokemonCaught[pokemonName] = resp
	fmt.Printf("%s was caught\n", pokemonName)
	return nil

	// fmt.Printf("Pokemon in %s:\n", resp.Name)
	// for _, pokemon := range resp. {
	// 	fmt.Println(pokemon.Pokemon.Name)
	// }

}
