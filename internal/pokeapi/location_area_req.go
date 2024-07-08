package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area?offset=0&limit=20"
	var fullURL string
	if pageURL != nil {
		fullURL = *pageURL
	} else {
		fullURL = baseURL + endpoint
	}
	// fmt.Println(fullURL)

	dat, ok := cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil
	}
	fmt.Println("cache miss")

	res, err := http.Get(fullURL)

	if err != nil {
		return LocationAreasResp{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		// log.Fatal(err)
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}

	//now we have to unmarshal the json body to struct
	err = json.Unmarshal(body, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	cache.Add(fullURL, body)
	return locationAreasResp, nil
}
