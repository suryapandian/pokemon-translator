package translator

import (
	"testing"

	"github.com/suryapandian/pokemon-translator/logger"

	"github.com/stretchr/testify/assert"
)

func GetTranslation(t *testing.T) {
	a := assert.New(t)
	translatorCache := NewTranslatorStore(logger.LogEntryWithRef())
	translation, err := NewTranslatorService(&TranslatorMock{}, translatorCache).GetTranslation("test data")
	a.Nil(err)
	a.NotEmpty(translation)
}
