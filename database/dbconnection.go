package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dbStr string, logger hclog.Logger) (*sqlx.DB, error) {
	logger.Debug("Database Connection starting")
	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {
		logger.Error("Failed to open a db connection", "error", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}

func Migrate(dbStr string, logger hclog.Logger) error {
	logger.Debug("migration starting")
	m, err := migrate.New(
		"file://migrations",
		dbStr,
	)
	if err != nil {
		logger.Error("Failed to Make Migrations", "error", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("Failed to Migrate")
		return err
	}

	return nil
}
