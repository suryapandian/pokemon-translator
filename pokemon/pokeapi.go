package pokemon

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type PokemonAPI interface {
	PokemonSpecies(name string) (structs.PokemonSpecies, error)
}

type PokeAPI struct{}

func (p *PokeAPI) PokemonSpecies(name string) (structs.PokemonSpecies, error) {
	return pokeapi.PokemonSpecies(name)
}
