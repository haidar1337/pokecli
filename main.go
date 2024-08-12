package main

import "fmt"

func main() {
	for true {
		
	}
}


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
	}

	return commands
}

func commandHelp() error {
	helpString :=  "Welcome to the Pokedex!\nUsage:\n\n"

	commands := getCommands()
	for _, command := range commands {
		helpString += command.name + " " + command.description + "\n"
	}

	fmt.Println(helpString)
}

func commandExit() error {
	
}