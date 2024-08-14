package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)


type General struct {
	PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters []struct {
	Pokemon struct{
		Name string
		Url string
	} `json:"pokemon"`
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