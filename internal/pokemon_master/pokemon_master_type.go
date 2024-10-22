package pokemon_master

import "github.com/gbelintani/pokedex/internal/pokeapi"

type PokemonMaster struct {
	pokedex map[string]pokeapi.Pokemon
}
