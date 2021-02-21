package translator

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

type TranslatorStore struct {
	Cache  *sync.Map
	logger *logrus.Entry
}

var cache sync.Map

func NewTranslatorStore(logger *logrus.Entry) *TranslatorStore {
	return &TranslatorStore{
		Cache:  &cache,
		logger: logger,
	}
}

var ErrTranslationNotFound = errors.New("translation not found")

func (t *TranslatorStore) Get(text string) (string, error) {
	translation, ok := t.Cache.Load(text)
	if !ok {
		return "", ErrTranslationNotFound
	}
	t.logger.WithField("text", text).Infof("translation retrieved from cache")
	return translation.(string), nil

}

func (t *TranslatorStore) Save(text, translation string) {
	t.logger.WithField("text", text).Infof("translation saved in cache")
	t.Cache.Store(text, translation)
}
