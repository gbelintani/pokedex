package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	args     []string
}

var commands map[string]command
var currentConfig config
var client pokeapi.Client

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
		"explore": {
			name:        "explore <location>",
			description: "Get pokemon in location area",
			callback:    commandExplore,
		},
	}

	currentConfig = config{}
	client = pokeapi.NewClient()
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		reader.Scan()
		input := reader.Text()

		words := strings.Split(input, " ")
		if len(words) == 0 {
			continue
		}

		if len(words) > 1 {
			currentConfig.args = words[1:]
		}

		command, ok := commands[words[0]]
		if !ok {
			fmt.Printf("Invalid Command %s", input)
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(fmt.Errorf("error on command %s: %w", input, err))
			continue
		}
	}
}

func commandExplore() error {
	if len(currentConfig.args) != 1 {
		return fmt.Errorf("invalid args %s", currentConfig.args)
	}

	res, err := client.GetLocationDetail(currentConfig.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", res.Name)
	fmt.Println("Found Pokemon:")
	for _, item := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", item.Pokemon.Name)
	}

	return nil
}

func commandMap() error {
	res, err := client.GetLocationAreas(currentConfig.next)
	if err != nil {
		return err
	}
	currentConfig.next = res.Next
	currentConfig.previous = res.Previous
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}

func commandMapb() error {
	res, err := client.GetLocationAreas(currentConfig.previous)
	if err != nil {
		return err
	}
	currentConfig.next = res.Next
	currentConfig.previous = res.Previous
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
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
