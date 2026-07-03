/*
Copyright © 2026 biplob-codes
*/
package main

import (
	"log"

	"github.com/biplob-codes/shelf-cli/cmd"
	"github.com/biplob-codes/shelf-cli/internal/db"
	"github.com/biplob-codes/shelf-cli/internal/store"
	_ "modernc.org/sqlite"
)

func main() {
	database, err := db.Connect("shelf.db")
	if err != nil {
		log.Fatalf("Database connection : %v", err)
	}
	defer database.Close()
	if err := db.Migrate(database); err != nil {
		log.Fatalf("Database migration : %v", err)
	}
	repo := store.NewLinkRepository(database)
	cmd.RootCMD.AddCommand(cmd.CollectionCMD(repo))
	cmd.RootCMD.AddCommand(cmd.LinkCMD(repo))
	if err := cmd.RootCMD.Execute(); err != nil {
		log.Fatal("Root CMD: ", err)
	}

}
