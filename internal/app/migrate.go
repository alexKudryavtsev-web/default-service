package app

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	dbURL := os.Getenv("PG_URL")
	if dbURL == "" {
		log.Fatal("PG_URL environment variable not set")
	}

	migrationsPath := "file://./migrations"

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			log.Printf("Source error: %v", srcErr)
		}
		if dbErr != nil {
			log.Printf("Database error: %v", dbErr)
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully!")

	version, dirty, err := m.Version()
	if err != nil {
		log.Printf("Failed to get migration version: %v", err)
	} else {
		fmt.Printf("Current migration version: %d, dirty: %t\n", version, dirty)
	}

	if dirty {
		log.Println("WARNING: Database is in a 'dirty' state.  Manual intervention may be required.")
	}
}
