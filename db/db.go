package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS lunchorder (
		iid         INTEGER PRIMARY KEY AUTOINCREMENT,
		srestaurant TEXT NOT NULL,
		jmetadata   TEXT,
		dtcreatedon TEXT DEFAULT (datetime('now')),
		bisdeleted  INTEGER DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS officetally (
		iid         INTEGER PRIMARY KEY AUTOINCREMENT,
		srestaurant TEXT NOT NULL,
		itally      INTEGER NOT NULL,
		dtcreatedon TEXT DEFAULT (datetime('now')),
		bisdeleted  INTEGER DEFAULT 0
	);
	`
	_, err := db.Exec(schema)
	return err
}
