package config

import (
	"os"
)

var (
	PORT                       string
	LOG_LEVEL                  string
	SHAKESPEARE_TRANSLATOR_URL string
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	if LOG_LEVEL == "" {
		LOG_LEVEL = "INFO"
	}
	SHAKESPEARE_TRANSLATOR_URL = os.Getenv("SHAKESPEARE_TRANSLATOR_URL")
	if SHAKESPEARE_TRANSLATOR_URL == "" {
		SHAKESPEARE_TRANSLATOR_URL = "https://api.funtranslations.com/translate/shakespeare"
	}
}
