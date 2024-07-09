package main

import (
	"errors"
	"fmt"

	"github.com/Tubelight30/my_pokedex/internal/pokeapi"
)

func commandMap(c *config, args ...string) error {

	resp, err := pokeapi.ListLocationAreas(c.nextLocationAreaURL)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	// fmt.Println(resp)
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.prevLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := pokeapi.ListLocationAreas(c.prevLocationAreaURL)
	if err != nil {
		return err
	}
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous
	return nil
}
