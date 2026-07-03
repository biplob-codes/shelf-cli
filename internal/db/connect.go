package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func Connect(name string) (*sql.DB, error) {
	dbPath, err := getDataSource()
	if err != nil {
		return nil, fmt.Errorf("getting  datasource: %w", err)

	}
	dataSource := fmt.Sprintf("%s?_pragma=foreign_keys(1)", dbPath)
	db, err := sql.Open("sqlite", dataSource)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		return nil, fmt.Errorf("pinging database: %w", pingErr)
	}

	return db, nil
}
