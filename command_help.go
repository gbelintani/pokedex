package main

import "fmt"

func commandHelp(config *config, args ...string) error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")

	for _, value := range getCommands() {
		fmt.Printf("%s: %s \r\n", value.name, value.description)
	}
	return nil
}
