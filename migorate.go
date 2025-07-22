package migorate

import (
	"database/sql"

	"github.com/waicodes/migorate/internal"
)

func Run(db *sql.DB, migrationDir string) error {
	return internal.Run(db, migrationDir)
}
