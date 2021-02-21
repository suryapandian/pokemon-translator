package handlers

import (
	"net/http"

	"github.com/suryapandian/pokemon-translator/pokemon"
	"github.com/suryapandian/pokemon-translator/translator"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type pokemonRouter struct {
	pokemonAPI    pokemon.PokemonAPI
	translatorAPI translator.TranslatorAPI
	translations  *translator.TranslatorStore
	pokemonStore  *pokemon.PokemonStore
	logger        *logrus.Entry
}

func newPokemonRouter(pokemonAPI pokemon.PokemonAPI, translatorAPI translator.TranslatorAPI, logger *logrus.Entry) *pokemonRouter {
	return &pokemonRouter{
		pokemonAPI:    pokemonAPI,
		translatorAPI: translatorAPI,
		pokemonStore:  pokemon.NewPokemonStore(logger),
		translations:  translator.NewTranslatorStore(logger),
		logger:        logger,
	}
}

func (p *pokemonRouter) setPokemonRoutes(router chi.Router) {
	router.Route("/pokemon", func(r chi.Router) {
		r.Get("/{pokemonName}", p.getTranslatedDescription)
	})
}

func (p *pokemonRouter) getTranslatedDescription(w http.ResponseWriter, r *http.Request) {
	pokemonName := chi.URLParam(r, "pokemonName")

	version := r.URL.Query().Get("version")
	p.logger = p.logger.WithField("pokemonName", pokemonName).WithField("version", version)
	p.logger.Infof("translations")

	pokemonDetails, err := p.pokemonStore.GetDetailsByName(pokemonName, version, p.pokemonAPI)
	if err != nil {
		p.logger.Errorf("error while fetching pokemon details %v", err)
		errorCode := http.StatusInternalServerError
		if err == pokemon.ErrPokemonNotFound {
			errorCode = http.StatusNotFound
		}
		writeJSONMessage(err.Error(), errorCode, w)
		return
	}

	translation, err := p.translations.GetTranslation(pokemonDetails.Description, p.translatorAPI)
	if err != nil {
		p.logger.Errorf("error while translating description %v", err)
		writeJSONMessage(err.Error(), http.StatusInternalServerError, w)
		return
	}

	var response struct {
		Name        string `json:"string"`
		Description string `json:"description"`
	}

	response.Name = pokemonName
	response.Description = translation
	writeJSONStruct(response, http.StatusOK, w)
}
