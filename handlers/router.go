package handlers

import (
	"github.com/suryapandian/pokemon-translator/logger"
	"github.com/suryapandian/pokemon-translator/pokemon"
	"github.com/suryapandian/pokemon-translator/translator"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRouter(pokemonAPI pokemon.PokemonAPI, translatorAPI translator.TranslatorAPI) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	setPingRoutes(mux)
	newPokemonRouter(
		pokemonAPI,
		translatorAPI,
		logger.LogEntryWithRef(),
	).setPokemonRoutes(mux)

	return mux
}
