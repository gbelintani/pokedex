package main

import (
	"fmt"
	"os"

	"github.com/gbelintani/pokedex/internal/pokeapi"
)

type command struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next     *string
	previous *string
}

var commands map[string]command
var currentConfig config

func init() {
	commands = map[string]command{
		"help": {
			name:        "help",
			description: "Displays commands avaiable",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the application",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous location areas",
			callback:    commandMapb,
		},
	}

	currentConfig = config{}
}

func main() {
	for {
		fmt.Print("\nPokedex > ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(fmt.Errorf("error reading input: %w", err))
		}
		command, ok := commands[input]
		if !ok {
			fmt.Printf("Invalid Command %s", input)
			continue
		}
		err = command.callback()
		if err != nil {
			fmt.Println(fmt.Errorf("error on command %s: %w", input, err))
			continue
		}
	}
}

func printLocationAreas(r pokeapi.PokeApiListLocationResponse) {
	for _, i := range r.Results {
		fmt.Printf("%s\n", i.Name)
	}
}

func commandMap() error {
	res, err := pokeapi.GetLocationAreas(currentConfig.next)
	if err != nil {
		return err
	}
	currentConfig.next = res.Next
	currentConfig.previous = res.Previous
	printLocationAreas(res)
	return nil
}

func commandMapb() error {
	res, err := pokeapi.GetLocationAreas(currentConfig.previous)
	if err != nil {
		return err
	}
	currentConfig.next = res.Next
	currentConfig.previous = res.Previous
	printLocationAreas(res)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")

	for _, value := range commands {
		fmt.Printf("%s: %s \r\n", value.name, value.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
