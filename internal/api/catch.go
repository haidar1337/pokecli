package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	BaseExperience int `json:"base_experience"`
	Stats []struct{
		BaseStat int `json:"base_stat"`
		Stat struct{
			Name string
		} `json:"stat"`
	} `json:"stats"`
	Types []struct{
		Type struct{
			Name string
		} `json:"type"`
	} `json:"types"`
}

type Pokedex struct {
	Pokemons map[string]Pokemon
}

var UserPokedex Pokedex = Pokedex{}

func CatchPokemon(name string) (string, error) {
	url := baseUrl + "/pokemon/" + name

	val, ok := UserPokedex.Pokemons[name]
	if ok {
		return "you already caught " + val.Name, nil
	}

	if cacheVal, ok := cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(cacheVal, &pokemon)
		if err != nil {
			return "", err
		}

		caught := calculateChance(pokemon.BaseExperience)
		if !caught {
			return "", fmt.Errorf("%s escaped!", pokemon.Name)
		}

		addToPokedex(pokemon.Name, pokemon)

		return pokemon.Name + " was caught!", nil
	}
	
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode >= 300 {
		return "", err
	}
	if err != nil {
		return "", nil
	}
	cache.Add(url, body)

	pokemon := Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return "", err
	}

	caught := calculateChance(pokemon.BaseExperience)
	if !caught {
		return "", fmt.Errorf("%s escaped!", pokemon.Name)
	}

	addToPokedex(pokemon.Name, pokemon)

	return pokemon.Name + " was caught!", nil
}

func addToPokedex(name string, pokemon Pokemon) {
	if UserPokedex.Pokemons == nil {
		m := make(map[string]Pokemon)
		UserPokedex.Pokemons = m
	}

	UserPokedex.Pokemons[name] = pokemon
}

func calculateChance(baseExp int) bool {
	random := rand.Intn((int(baseExp) * 2) - baseExp + 30)
	if random < baseExp {
		return false
	}

	return true
}