package pokemon

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

type PokemonCache interface {
	Get(uniqueID string) (*Pokemon, error)
	Save(pokemon *Pokemon)
}

type PokemonStore struct {
	Store  *sync.Map
	logger *logrus.Entry
}

var pokemonStore sync.Map

func NewPokemonStore(logger *logrus.Entry) *PokemonStore {
	return &PokemonStore{
		Store:  &pokemonStore,
		logger: logger,
	}
}

type Pokemon struct {
	UniqueID    string
	Name        string
	Version     string
	Description string
}

func (pokemon *Pokemon) setUniqueID() {
	pokemon.UniqueID = pokemon.Name + pokemon.Version
}

var ErrPokemonNotFound = errors.New("pokemon not found")

func (p *PokemonStore) Get(uniqueID string) (*Pokemon, error) {
	pokemon, ok := p.Store.Load(uniqueID)
	if !ok {
		return nil, ErrPokemonNotFound
	}
	p.logger.WithField("uniqueId", uniqueID).Infof("pokemon details retrieved from cache")
	return pokemon.(*Pokemon), nil

}

func (p *PokemonStore) Save(pokemon *Pokemon) {
	if pokemon.UniqueID == "" {
		pokemon.setUniqueID()
	}
	p.logger.WithField("uniqueId", pokemon.UniqueID).Infof("pokemon details saved in cache")
	p.Store.Store(pokemon.UniqueID, pokemon)
}
