package pokemon

import "github.com/mtslzr/pokeapi-go/structs"

type PokeMockAPI struct{}

func (p *PokeMockAPI) PokemonSpecies(name string) (structs.PokemonSpecies, error) {
	if name == "invalidPokemon" {
		return structs.PokemonSpecies{}, ErrPokemonNotFound
	}

	species := structs.PokemonSpecies{
		FlavorTextEntries: []struct {
			FlavorText string `json:"flavor_text"`
			Language   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		}{
			{
				FlavorText: "description1",
				Language: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "en",
				},
				Version: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "version1",
				},
			},
			{
				FlavorText: "description2",
				Language: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "en",
				},
				Version: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "version2",
				},
			},
		},
	}
	return species, nil
}
