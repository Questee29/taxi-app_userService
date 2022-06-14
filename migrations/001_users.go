package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUsers, downUsers)
}

func upUsers(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS users(
		"id" SERIAL PRIMARY KEY,
		"name" text,
		"phone" text,
		"email" text,
		"password" text);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downUsers(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS users;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
