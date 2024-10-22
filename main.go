package main

import (
	"github.com/gbelintani/pokedex/internal/pokeapi"
	"github.com/gbelintani/pokedex/internal/pokemon_master"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeClient:    pokeClient,
		pokemonMaster: pokemon_master.NewPokemonMaster(),
	}

	startRepl(cfg)
}
