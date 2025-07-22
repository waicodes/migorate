package internal

import "database/sql"

func ensureMigrationTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            filename TEXT UNIQUE NOT NULL,
            hash TEXT NOT NULL,
            applied_at TIMESTAMPTZ DEFAULT now()
        )
    `)
	return err
}
