package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/suryapandian/pokemon-translator/pokemon"
	"github.com/suryapandian/pokemon-translator/translator"

	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {
	var testCases = []struct {
		desc               string
		expectedStatusCode int
	}{
		{
			"sanity",
			http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			a := assert.New(t)
			r := httptest.NewRequest(http.MethodGet, "/ping", nil)
			w := httptest.NewRecorder()
			GetRouter(&pokemon.PokeMockAPI{}, &translator.TranslatorMock{}).ServeHTTP(w, r)
			response := w.Result()
			a.Equal(testCase.expectedStatusCode, response.StatusCode)
		})
	}

}
