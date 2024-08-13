package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
	}

	return commands
}


func read(scanner *bufio.Scanner, ch chan cliCommand) {
	for {
		commands := getCommands()
	
		scanner.Scan()
		input := scanner.Text()
		if command := commands[input]; command.name != "" {
			ch <- command
		}
	}
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan cliCommand)

	fmt.Print("Pokecli > ")

	go read(scanner, ch)

	for {
		if command := <-ch; command.name != "" {
			command.callback()
			fmt.Print("Pokecli > ")
		}
	}
}