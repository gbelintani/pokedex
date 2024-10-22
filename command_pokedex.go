package main

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {

	pokemons, err := config.pokemonMaster.GetAllFromPokedex()
	if err != nil {
		return err
	}

	fmt.Printf("Your Pokedex:\n")
	for _, p := range pokemons {
		fmt.Printf(" - %s\n", p.Name)
	}

	return nil
}
