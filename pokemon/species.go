package pokemon

func (p *PokemonStore) GetDetailsByName(name, version string, apiClient PokemonAPI) (*Pokemon, error) {
	pokemon := &Pokemon{Name: name, Version: version}
	pokemon.setUniqueID()

	if pokemonCache, err := p.Get(pokemon.UniqueID); err == nil {
		return pokemonCache, err
	}

	species, err := apiClient.PokemonSpecies(name)
	if err != nil {
		return nil, ErrPokemonNotFound
	}

	for _, des := range species.FlavorTextEntries {
		if version != "" && version != des.Version.Name {
			continue
		}

		if des.Language.Name == "en" {
			pokemon.Description = des.FlavorText
			break
		}
	}
	p.Save(pokemon)
	return pokemon, err
}
