package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid args: %s", args)
	}

	pokemon, err := config.pokeClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	captureChance := 0.5

	if rand.Float64() > captureChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.pokemonMaster.AddToPokedex(pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
