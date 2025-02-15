package app

import (
	"log"

	"github.com/alexKudryavtsev-web/default-service/config"
	"github.com/alexKudryavtsev-web/default-service/pkg/logger"
)

func Run(cfg *config.Config) {
	logger, err := logger.New(cfg.Log.Level, cfg.Log.Destination)
	if err != nil {
		log.Fatalf("can't init logger: %s", err)
	}
	logger.Info("logger init")
}
