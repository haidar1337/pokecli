package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/haidar1337/pokecli/internal/pokecache"
)

type Result struct {
	Name string
	Url string
}

type LocationArea struct {
	Count int 
	Next *string
	Previous *string
	Results []Result
}

type General struct {
	PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters []struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}


var cache = pokecache.NewCache(time.Minute * 5)

func FetchNextTwentyLocations(config *Config) ([]Result, error) {
	if config.Next == nil {
		return nil, errors.New("Invalid config")
	}

	locationArea := LocationArea{}
	var results []Result

	val, ok := cache.Get(*config.Next)
	if ok {
		fmt.Printf("Cache hit")
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return nil, err
		}
	} else {
		res, err := http.Get(*config.Next)
		if err != nil {
			return nil, errors.New("Unable to fetch next 20 locations")
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode >= 300 {
			errorMessage := fmt.Sprintf("Response failed with status code: %d", res.StatusCode)
			return nil, errors.New(errorMessage)
		}
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &locationArea)
		if err != nil {
			return nil, err
		}

		cache.Add(*config.Next, body)
	}

	results = locationArea.Results
	setConfig(locationArea.Next, locationArea.Previous)

	return results, nil
}

func FetchPreviousTwentyLocations(config *Config) ([]Result, error) {
	if config.Previous == nil {
		return nil, errors.New("Invalid config")
	}

	locationArea := LocationArea{}
	var results []Result

	val, ok := cache.Get(*config.Previous)
	if ok {
		fmt.Printf("Cache hit")
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return nil, err
		}
	} else {
		res, err := http.Get(*config.Previous)
		if err != nil {
			return nil, errors.New("Unable to fetch previous 20 locations")
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode >= 300 {
			errorMessage := fmt.Sprintf("Response failed with status code: %d", res.StatusCode)
			return nil, errors.New(errorMessage)
		}
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &locationArea)
		if err != nil {
			return nil, err
		}

		cache.Add(*config.Previous, body)
	}

	results = locationArea.Results
	setConfig(locationArea.Next, locationArea.Previous)

	return results, nil
}

func FetchPokemonsArea(s string) (PokemonEncounters, error) {
	url := baseUrl + "/location-area/" + s

	val, ok := cache.Get(url)
	if ok {
		general := General{}
		if err := json.Unmarshal(val, &general); err == nil {
			return general.PokemonEncounters, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		errorMessage := fmt.Sprintf("Response failed with status code: %d", res.StatusCode)
		return nil, errors.New(errorMessage)
	}
	if err != nil {
		return nil, err
	}
	
	pokemonEncounters := General{}
	if err := json.Unmarshal(body, &pokemonEncounters); err != nil {
		return nil, err
	}

	cache.Add(url, body)

	return pokemonEncounters.PokemonEncounters, nil
}