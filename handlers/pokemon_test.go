package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/suryapandian/pokemon-translator/pokemon"
	"github.com/suryapandian/pokemon-translator/translator"

	"github.com/stretchr/testify/assert"
)

func TestGetTranslatedDescription(t *testing.T) {
	var testCases = []struct {
		desc               string
		path               string
		expectedStatusCode int
	}{
		{
			"test API",
			"/pokemon/testpokemon",
			http.StatusOK,
		},
		{
			"test API with version",
			"/pokemon/testpokemon?version=2",
			http.StatusOK,
		},
		{
			"invalid pokemon character",
			"/pokemon/invalidPokemon",
			http.StatusNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			a := assert.New(t)
			r := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			w := httptest.NewRecorder()
			GetRouter(&pokemon.PokeMockAPI{}, &translator.TranslatorMock{}).ServeHTTP(w, r)
			response := w.Result()
			a.Equal(testCase.expectedStatusCode, response.StatusCode)
		})
	}

}
