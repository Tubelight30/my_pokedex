package main

import "github.com/Tubelight30/my_pokedex/internal/pokeapi"

type config struct {
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	pokemonCaught       map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
		pokemonCaught:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)

	// ticker := time.NewTicker(500 * time.Millisecond)
	// for {

	// 	t := <-ticker.C
	// 	fmt.Println(t)
	// }

}
