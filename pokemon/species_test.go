package pokemon

import (
	"testing"

	"github.com/suryapandian/pokemon-translator/logger"

	"github.com/stretchr/testify/assert"
)

func TestGetDescriptionByName(t *testing.T) {
	var testCases = []struct {
		version         string
		description     string
		testDescription string
	}{
		{
			version:         "",
			description:     "description1",
			testDescription: "get description without version",
		},
		{
			version:         "version1",
			description:     "description1",
			testDescription: "get description with version",
		},
		{
			version:         "version1",
			description:     "description1",
			testDescription: "get description from cache",
		},
		{
			version:         "version2",
			description:     "description2",
			testDescription: "get description with different version",
		},
	}
	a := assert.New(t)
	pokemonCache := NewPokemonStore(logger.LogEntryWithRef())
	pokemonService := NewPokemonService(&PokeMockAPI{}, pokemonCache)
	for _, testCase := range testCases {
		t.Run(testCase.testDescription, func(t *testing.T) {
			pokemonDetails, err := pokemonService.GetDetailsByName("pokemon", testCase.version)
			a.Nil(err, testCase)
			a.Equal(testCase.description, pokemonDetails.Description)
		})
	}
}
