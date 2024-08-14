package main

import (
	"errors"
	"fmt"

	"github.com/haidar1337/pokecli/internal/api"
)

func commandCatch(args ...string) error {
	if len(args) != 1 {
		return errors.New("A catch command requires an argument. Usage: catch <pokemon_name>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	val, err := api.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}