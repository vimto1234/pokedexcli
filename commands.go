package main

import (
	"fmt"
	"os"

	"github.com/vimto1234/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
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
		"map": {
			name:        "map",
			description: "Displays the next twenty locations in the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous twenty locations in the map",
			callback:    commandMapB,
		},
	}
	return mapOfCommands
}

func executeCommand(command string) {
	elem, ok := getCommands()[command]
	if !ok {
		fmt.Printf("Unknown command '%v'", command)
		return
	}

	err := elem.callback(&mainConfig)
	if err != nil {
		fmt.Print(err)
	}
}

func commandExit(c *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, elem := range getCommands() {
		fmt.Printf("%v: %v\n", elem.name, elem.description)
	}
	return nil
}

func commandMap(c *config) error {
	location, err := pokeapi.GetLocation(c.next, c.locationCache)
	if err != nil {
		return err
	}
	c.next = location.Next
	c.previous = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.previous == "" {
		fmt.Printf("you're on the first page")
		return nil
	}

	fmt.Print("test")

	location, err := pokeapi.GetLocation(c.previous, c.locationCache)
	if err != nil {
		return err
	}
	c.next = location.Next
	c.previous = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
