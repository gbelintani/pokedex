package main

import "fmt"

func commandMap(config *config, args ...string) error {
	res, err := config.pokeClient.GetLocationAreas(config.next)
	if err != nil {
		return err
	}
	config.next = res.Next
	config.previous = res.Previous
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}

func commandMapb(config *config, args ...string) error {
	res, err := config.pokeClient.GetLocationAreas(config.previous)
	if err != nil {
		return err
	}
	config.next = res.Next
	config.previous = res.Previous
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}
