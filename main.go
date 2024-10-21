package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]command

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
	}
}

func main() {
	for {
		fmt.Printf("Pokedex > ")
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
