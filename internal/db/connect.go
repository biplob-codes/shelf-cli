package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func Connect(name string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s?_pragma=foreign_keys(1)", name)
	db, err := sql.Open("sqlite", dataSource)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}
	
	if pingErr := db.Ping(); pingErr != nil {
		return nil, fmt.Errorf("pinging database: %w", pingErr)
	}

	fmt.Println("Database connection established successfully.")

	return db, nil
}
