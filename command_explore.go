package main

import (
	"errors"
	"fmt"

	"github.com/haidar1337/pokecli/internal/api"
)

func commandExplore(args ...string) error {
	if len(args) != 1 {
		return errors.New("An explore command requires an argument. Usage: explore <area_name>")
	}
	pokemons, err := api.FetchPokemonsArea(args[0])
	if err != nil {
		return fmt.Errorf("An error has occurred, check the area name provided. %w", err)
	}

	for _, p := range pokemons {
		fmt.Println(p.Pokemon.Name)
	}

	return nil
}