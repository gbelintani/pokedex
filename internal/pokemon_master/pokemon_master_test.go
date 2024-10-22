package pokemon_master

import (
	"testing"

	"github.com/gbelintani/pokedex/internal/pokeapi"
)

func TestAddPokedex(t *testing.T) {
	master := NewPokemonMaster()
	master.AddToPokedex(pokeapi.Pokemon{Name: "test"})

	_, ok := master.pokedex["test"]
	if !ok {
		t.Errorf("test should be in pokedex")
	}
}
