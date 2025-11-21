package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	var mapOfCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return mapOfCommands
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, elem := range getCommands() {
		fmt.Printf("%v: %v\n", elem.name, elem.description)
	}
	return nil
}

func executeCommand(command string) {
	elem, ok := getCommands()[command]
	if !ok {
		fmt.Print("Unknown command")
		return
	}

	elem.callback()
}
