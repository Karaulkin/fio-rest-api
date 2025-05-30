package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log/slog"
	"path/filepath"
)

type DB struct {
	*sql.DB
}

// NewDB создает новое соединение с базой данных
func NewDB(cfg *config.Config) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)
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

	log.Info("running migrations")
	log.Debug("running migrations from", absPath)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	// Применить все миграции из папки migrations
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}
	return nil

}
