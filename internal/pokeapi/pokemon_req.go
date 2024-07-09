package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// dat, ok
	res, err := http.Get(fullURL)

	dat, ok := cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!!")
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	fmt.Println("cache miss")
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		return Pokemon{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(fullURL, body)
	return pokemon, nil

}
