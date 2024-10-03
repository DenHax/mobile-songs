package migrations

import (
	"database/sql"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(storageUrl string, migrationsPath string, log *slog.Logger) {
	if storageUrl == "" {
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	db, err := sql.Open("postgres", storageUrl)
	if err != nil {
		log.Error("failed to connect to the database", slog.String("error", err.Error()))
	} else {
		log.Debug("success connect to database")
	}

	defer db.Close()

	m, err := migrate.New("file://"+migrationsPath, storageUrl)

	if err != nil {
		log.Error("failed to create migrate instance", slog.String("error", err.Error()))
	} else {
		log.Debug("success find migrate files")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed to apply migrations", slog.String("error", err.Error()))
	}

	log.Info("migrations applied successfully")
}
