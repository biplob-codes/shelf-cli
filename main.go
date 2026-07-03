/*
Copyright © 2026 biplob-codes
*/
package main

import (
	"os"

	"github.com/biplob-codes/shelf-cli/cmd"
	"github.com/biplob-codes/shelf-cli/internal/db"
	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/biplob-codes/shelf-cli/internal/ui"
	_ "modernc.org/sqlite"
)

func main() {
	database, err := db.Connect("shelf.db")
	if err != nil {
		ui.PrintError("Database connection: %v", err)
		os.Exit(1)
	}
	defer database.Close()

	if err := db.Migrate(database); err != nil {
		ui.PrintError("Database migration: %v", err)
		os.Exit(1)
	}

	collectionRepo := store.NewCollectionRepository(database)
	linkRepo := store.NewLinkRepository(database, collectionRepo)

	cmd.RootCMD.AddCommand(cmd.CollectionCMD(collectionRepo))
	cmd.RootCMD.AddCommand(cmd.LinkCMD(linkRepo))

	if err := cmd.RootCMD.Execute(); err != nil {
		ui.PrintErrorMsg(err.Error())
		os.Exit(1)
	}
}