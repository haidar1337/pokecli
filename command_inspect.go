package main

import (
	"errors"
	"fmt"

	"github.com/haidar1337/pokecli/internal/api"
)

func commandInspect(args ...string) error {
	if len(args) != 1 {
		return errors.New("An inspect command requires an argument. Usage: inspect <pokemon_name>")
	}

	pokemon, ok := api.UserPokedex.Pokemons[args[0]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	
	pokemonInfo := fmt.Sprintf("Name: %s\nHeight: %v\nWeight: %v\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		// pokemonInfo += " " + stat.Stat.Name + ": " + string(stat.BaseStat) + "\n"
		str := fmt.Sprintf(" %s: %v\n", stat.Stat.Name, stat.BaseStat)
		pokemonInfo += str
	}
	pokemonInfo += "Types:\n"
	for _, t := range pokemon.Types {
		pokemonInfo += " - " + t.Type.Name
	}

	fmt.Println(pokemonInfo)

	return nil
}