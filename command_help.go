package main

import "fmt"

func commandHelp(args ...string) error {
	helpString :=  "Welcome to the Pokedex!\nUsage:\n\n"

	commands := getCommands()
	for _, command := range commands {
		helpString += command.name + " " + command.description + "\n"
	}

	fmt.Println(helpString)
	return nil
}