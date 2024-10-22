package main

import "fmt"

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid args %s", args)
	}

	res, err := config.pokeClient.GetLocationDetail(args[0])
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
