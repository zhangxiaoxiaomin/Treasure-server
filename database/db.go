package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Init initializes the SQLite database
func Init(dbPath string) error {
	var err error

	// Ensure directory exists
	dbDir := dbPath[:len(dbPath)-len("treasure.db")]
	if dbDir != "" {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return fmt.Errorf("failed to create database directory: %w", err)
		}
	}

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Enable WAL mode and foreign keys
	pragmas := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA foreign_keys=ON",
		"PRAGMA busy_timeout=5000",
	}
	for _, pragma := range pragmas {
		if _, err := DB.Exec(pragma); err != nil {
			return fmt.Errorf("failed to set pragma: %w", err)
		}
	}

	if err := createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

func createTables() error {
	schema := `
	CREATE TABLE IF NOT EXISTS collections (
		id TEXT PRIMARY KEY,
		title_cn TEXT NOT NULL DEFAULT '',
		title_en TEXT NOT NULL DEFAULT '',
		category TEXT NOT NULL DEFAULT 'other',
		image TEXT NOT NULL DEFAULT '',
		detail_images TEXT NOT NULL DEFAULT '[]',
		views INTEGER NOT NULL DEFAULT 0,
		likes INTEGER NOT NULL DEFAULT 0,
		comment_count INTEGER NOT NULL DEFAULT 0,
		badge_cn TEXT NOT NULL DEFAULT '',
		badge_en TEXT NOT NULL DEFAULT '',
		date_str_cn TEXT NOT NULL DEFAULT '',
		date_str_en TEXT NOT NULL DEFAULT '',
		description_cn TEXT NOT NULL DEFAULT '',
		description_en TEXT NOT NULL DEFAULT '',
		detail_desc_cn TEXT NOT NULL DEFAULT '',
		detail_desc_en TEXT NOT NULL DEFAULT '',
		comments TEXT NOT NULL DEFAULT '[]',
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := DB.Exec(schema)
	return err
}

// TimestampRFC3339 returns current timestamp in RFC3339 format
func TimestampRFC3339() string {
	return time.Now().Format(time.RFC3339)
}