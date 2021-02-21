package handlers

import (
	"net/http"

	"github.com/suryapandian/pokemon-translator/pokemon"
	"github.com/suryapandian/pokemon-translator/translator"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type pokemonRouter struct {
	pokemonService    *pokemon.PokemonService
	translatorService *translator.TranslatorService
	logger            *logrus.Entry
}

func newPokemonRouter(pokemonAPI pokemon.PokemonAPI, translatorAPI translator.TranslatorAPI, logger *logrus.Entry) *pokemonRouter {
	return &pokemonRouter{
		pokemonService:    pokemon.NewPokemonService(pokemonAPI, pokemon.NewPokemonStore(logger)),
		translatorService: translator.NewTranslatorService(translatorAPI, translator.NewTranslatorStore(logger)),
		logger:            logger,
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

	pokemonDetails, err := p.pokemonService.GetDetailsByName(pokemonName, version)
	if err != nil {
		p.logger.Errorf("error while fetching pokemon details %v", err)
		errorCode := http.StatusInternalServerError
		if err == pokemon.ErrPokemonNotFound {
			errorCode = http.StatusNotFound
		}
		writeJSONMessage(err.Error(), errorCode, w)
		return
	}

	translation, err := p.translatorService.GetTranslation(pokemonDetails.Description)
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
