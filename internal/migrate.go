package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Migration struct {
	Name string
	Path string
	Hash string
}

func Run(db *sql.DB, migrationDir string) error {
	if err := ensureMigrationTable(db); err != nil {
		return err
	}

	files, err := readSQLFiles(migrationDir)
	if err != nil {
		return err
	}

	applied, err := getAppliedMigrations(db)
	if err != nil {
		return err
	}

	for _, f := range files {
		if hash, ok := applied[f.Name]; ok {
			if hash != f.Hash {
				return fmt.Errorf("hash mismatch for '%s': file was modified", f.Name)
			}
			log.Printf("Skipping: %s", f.Name)
			continue
		}

		log.Printf("Applying: %s", f.Name)
		if err := applySQL(db, f.Path); err != nil {
			return fmt.Errorf("failed to apply %s: %w", f.Name, err)
		}

		if err := recordMigration(db, f.Name, f.Hash); err != nil {
			return err
		}

		log.Printf("Applied: %s", f.Name)
	}

	log.Println("All migrations applied.")
	return nil
}

func getAppliedMigrations(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query("SELECT filename, hash FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[string]string)
	for rows.Next() {
		var name, hash string
		if err := rows.Scan(&name, &hash); err != nil {
			return nil, err
		}
		m[name] = hash
	}

	return m, nil
}

func applySQL(db *sql.DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	return err
}

func recordMigration(db *sql.DB, filename, hash string) error {
	_, err := db.Exec(`INSERT INTO migrations (filename, hash) VALUES ($1, $2)`, filename, hash)
	return err
}
