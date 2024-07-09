package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ExploreArea(location string) (LocationArea, error) {
	endpoint := "/location-area/" + location
	fullURL := baseURL + endpoint

	// dat, ok
	res, err := http.Get(fullURL)

	dat, ok := cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!!")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Println("cache miss")
	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		return LocationArea{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	cache.Add(fullURL, body)
	return locationArea, nil

}
