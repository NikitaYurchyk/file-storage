package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Path        string
	LogLevel    logger.LogLevel
	AutoMigrate bool
}

type DB struct {
	*gorm.DB
}

func New(cfg *Config) (*DB, error) {
	if cfg.Path == "" {
		cfg.Path = "data/app.db"
	}

	if dir := filepath.Dir(cfg.Path); dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create database directory: %w", err)
		}
	}

	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{
		Logger: logger.Default.LogMode(cfg.LogLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if cfg.AutoMigrate {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, fmt.Errorf("failed to get sql.DB: %w", err)
		}
		if err := RunMigrations(sqlDB); err != nil {
			return nil, fmt.Errorf("failed to run migrations: %w", err)
		}
	}

	return &DB{db}, nil
}

// SqlDB returns the underlying *sql.DB
func (d *DB) SqlDB() (*sql.DB, error) {
	return d.DB.DB()
}
