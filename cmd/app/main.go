package main

import (
	"log"

	"github.com/alexKudryavtsev-web/default-service/config"
	"github.com/alexKudryavtsev-web/default-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("can't init config: %s", err)
	}

	app.Run(cfg)
}
