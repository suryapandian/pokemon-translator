package pokemon

type PokemonService struct {
	ApiClient PokemonAPI
	Cache     PokemonCache
}

func NewPokemonService(apiClient PokemonAPI, cache PokemonCache) *PokemonService {
	return &PokemonService{ApiClient: apiClient, Cache: cache}
}

func (p *PokemonService) GetDetailsByName(name, version string) (*Pokemon, error) {
	pokemon := &Pokemon{Name: name, Version: version}
	pokemon.setUniqueID()

	if pokemonCache, err := p.Cache.Get(pokemon.UniqueID); err == nil {
		return pokemonCache, err
	}

	species, err := p.ApiClient.PokemonSpecies(name)
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
	p.Cache.Save(pokemon)
	return pokemon, err
}
