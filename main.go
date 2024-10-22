package main

import (
	"github.com/gbelintani/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeClient: pokeClient,
	}

	startRepl(cfg)
}
