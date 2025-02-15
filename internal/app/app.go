package app

import (
	"context"
	"fmt"
	"log"

	"github.com/alexKudryavtsev-web/default-service/config"
	"github.com/alexKudryavtsev-web/default-service/pkg/logger"
	"github.com/alexKudryavtsev-web/default-service/pkg/postgres"
)

func Run(cfg *config.Config) {
	logger, err := logger.New(cfg.Log.Level, cfg.Log.Destination)
	if err != nil {
		log.Fatalf("can't init logger: %s", err)
	}
	logger.Info("logger init")


	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatalf("can't init postgres: %s", err)
	}

	// tmp, test select
	query, args, err := pg.Builder.Select("version()").ToSql()
  if err != nil {
   log.Fatalf("Failed to build SQL: %v", err)
  }

  var version string
  err = pg.Pool.QueryRow(context.Background(), query, args...).Scan(&version)
  if err != nil {
   log.Fatalf("Failed to execute query: %v", err)
  }

  fmt.Println("PostgreSQL Version:", version)
}
