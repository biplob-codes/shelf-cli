package db

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {
	createLinkTable := `
    CREATE TABLE IF NOT EXISTS collection(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL UNIQUE,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS link(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		tag TEXT,
		collection_id INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (collection_id) REFERENCES collection(id) ON DELETE CASCADE
		);
	
		`

	_, err := db.Exec(createLinkTable)
	if err != nil {
		return fmt.Errorf("Opening database : %w", err)
	}
	fmt.Println("Database migrated successfully")
	return nil
}
