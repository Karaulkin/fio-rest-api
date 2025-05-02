package repository

import (
	"database/sql"
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/config"
	"github.com/pressly/goose/v3"
	"log/slog"
	"path/filepath"
)

type DB struct {
	*sql.DB
}

// NewDB создает новое соединение с базой данных
func NewDB(cfg *config.Config) (*DB, error) {
	connStr := "postgres://" + cfg.Database.Username + ":" + cfg.Database.Password + "@" + cfg.Database.Host + ":" + cfg.Database.Port

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

// RunMigrations запускает миграции базы данных
func RunMigrations(db *sql.DB, migrationsDir string, log *slog.Logger) error {
	absPath, err := filepath.Abs(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for migrations: %w", err)
	}

	log.Info("Running migrations from %s", absPath)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	// Применить все миграции из папки migrations
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}
	return nil

}
