package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/suryapandian/pokemon-translator/config"
	"github.com/suryapandian/pokemon-translator/handlers"
	"github.com/suryapandian/pokemon-translator/logger"
	"github.com/suryapandian/pokemon-translator/pokemon"
	"github.com/suryapandian/pokemon-translator/translator"

	"github.com/sirupsen/logrus"
)

func main() {
	logger.SetupLog(config.LOG_LEVEL)
	server := &http.Server{
		Addr: ":" + config.PORT,
		Handler: handlers.GetRouter(
			&pokemon.PokeAPI{},
			translator.NewTranslator(config.SHAKESPEARE_TRANSLATOR_URL),
		),
	}

	go func(server *http.Server) {
		logrus.Infof("Listening on port %s", config.PORT)
		if err := server.ListenAndServe(); err != nil {
			logrus.WithError(err).Fatal("Failed to start server!")
		}
	}(server)

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	logrus.Infof("gracefully shutting down server")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.WithError(err).Fatal("error shutting server down gracefully")
	}
}
