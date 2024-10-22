package main

import (
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid args: %s", args)
	}

	pokemon, err := config.pokemonMaster.GetFromPokedex(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf("%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
