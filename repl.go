package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func(...string) error
}


func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help":
			{
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},
			"exit":
			{
				name: "exit",
				description: "Exits the program",
				callback: commandExit,
			},
			"map":
			{
				name: "map",
				description: "Display the next 20 locations in the world of Pokemon",
				callback: commandMap,
			},
			"mapb":
			{
				name: "mapb",
				description: "Display the previous 20 locations in the world of Pokemon",
				callback: commandMapb,
			},
			"explore":
			{
				name: "explore",
				description: "Explore Pokemons in a certain area of the world of Pokemon. explore <area_name>",
				callback: commandExplore,
			},
			"catch":
			{
				name: "catch",
				description: "Throw a Pokeball to catch a Pokemon. catch <pokemon_name>",
				callback: commandCatch,
			},
			"inspect":
			{
				name: "inspect",
				description: "Inspect a Pokemon to show its information and stats. inspect <pokemon_name>",
				callback: commandInspect,
			},
			"pokedex":
			{
				name: "pokedex",
				description: "Display the list of Pokemons you have caught",
				callback: commandPokedex,
			},
	}

	return commands
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokecli > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		if err := command.callback(args...); err != nil {
			fmt.Println(err)
		}
	}
}