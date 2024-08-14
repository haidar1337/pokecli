package main

import (
	"errors"
	"fmt"

	"github.com/haidar1337/pokecli/internal/api"
)


func commandPokedex(args ...string) error {
	pokedex := api.UserPokedex.Pokemons
	if pokedex == nil {
		return errors.New("you have not caught any Pokemon yet. use the catch command to catch a Pokemon")
	}

	pokemons := "Your Pokedex:\n"
	for key := range pokedex {
		str := fmt.Sprintf(" - %s\n", key)
		pokemons += str
	}

	fmt.Println(pokemons)

	return nil
}