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
	callback    func(*config, ...string) error
}

type config struct {
	pokeClient pokeapi.Client
	next       *string
	previous   *string
}

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("\nPokedex > ")
		reader.Scan()
		input := reader.Text()

		words := strings.Split(input, " ")
		if len(words) == 0 {
			continue
		}

		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := commands[words[0]]
		if !ok {
			fmt.Printf("Invalid Command %s", input)
			continue
		}
		err := command.callback(config, args...)
		if err != nil {
			fmt.Println(fmt.Errorf("error on command %s: %w", input, err))
			continue
		}
	}
}

func getCommands() map[string]command {
	return map[string]command{
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
}
