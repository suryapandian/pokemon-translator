package translator

import (
	"testing"

	"github.com/suryapandian/pokemon-translator/logger"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	a := assert.New(t)
	translation, err := NewTranslatorStore(logger.LogEntryWithRef()).Get("test data")
	a.Empty(translation)
	a.Equal(ErrTranslationNotFound, err)

	NewTranslatorStore(logger.LogEntryWithRef()).Save("test data", "test data translated")
	translation, err = NewTranslatorStore(logger.LogEntryWithRef()).Get("test data")
	a.Nil(err)
	a.NotEmpty(translation)
}
