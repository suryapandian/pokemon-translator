package translator

import (
	"testing"

	"github.com/suryapandian/pokemon-translator/logger"

	"github.com/stretchr/testify/assert"
)

func GetTranslation(t *testing.T) {
	a := assert.New(t)
	translatorAPI := TranslatorMock{}
	translation, err := NewTranslatorStore(logger.LogEntryWithRef()).GetTranslation("test data", &translatorAPI)
	a.Nil(err)
	a.NotEmpty(translation)
}
