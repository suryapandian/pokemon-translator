package pokemon

import (
	"testing"

	"github.com/suryapandian/pokemon-translator/logger"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	pokemon := &Pokemon{
		Name:    "charizard",
		Version: "ruby",
	}
	pokemon.setUniqueID()
	pokemonCache := NewPokemonStore(logger.LogEntryWithRef())
	a := assert.New(t)
	pokemonDetails, err := pokemonCache.Get(pokemon.UniqueID)
	a.Empty(pokemonDetails)
	a.Equal(ErrPokemonNotFound, err)

	pokemonCache.Save(pokemon)
	pokemonDetails, err = pokemonCache.Get(pokemon.UniqueID)
	a.Nil(err)
	a.NotEmpty(pokemonDetails)
}
