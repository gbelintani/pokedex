package pokemon_master

import (
	"fmt"

	"github.com/gbelintani/pokedex/internal/pokeapi"
)

func NewPokemonMaster() PokemonMaster {
	return PokemonMaster{
		pokedex: map[string]pokeapi.Pokemon{},
	}
}

func (u *PokemonMaster) AddToPokedex(pokemon pokeapi.Pokemon) error {
	u.pokedex[pokemon.Name] = pokemon
	fmt.Printf("New Pokedex Entry: %s! Total: %d\n", pokemon.Name, len(u.pokedex))
	return nil
}

func (u *PokemonMaster) GetFromPokedex(pokemon string) (pokeapi.Pokemon, error) {
	poke, ok := u.pokedex[pokemon]
	if !ok {
		return pokeapi.Pokemon{}, fmt.Errorf("you haven't caught %s yet", pokemon)
	}

	return poke, nil
}

func (u *PokemonMaster) GetAllFromPokedex() ([]pokeapi.Pokemon, error) {
	pokes := make([]pokeapi.Pokemon, 0, len(u.pokedex))
	for k := range u.pokedex {
		pokes = append(pokes, u.pokedex[k])
	}
	return pokes, nil

}
