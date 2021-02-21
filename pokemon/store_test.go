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

	a := assert.New(t)
	pokemonDetails, err := NewPokemonStore(logger.LogEntryWithRef()).Get(pokemon.UniqueID)
	a.Empty(pokemonDetails)
	a.Equal(ErrPokemonNotFound, err)

	NewPokemonStore(logger.LogEntryWithRef()).Save(pokemon)
	pokemonDetails, err = NewPokemonStore(logger.LogEntryWithRef()).Get(pokemon.UniqueID)
	a.Nil(err)
	a.NotEmpty(pokemonDetails)
}
