package pokeapi

import (
	"time"

	"github.com/Tubelight30/my_pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

var cache = pokecache.NewCache(5 * time.Minute)
