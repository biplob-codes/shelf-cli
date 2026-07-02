/*
Copyright © 2026 biplob-codes
*/
package main

import (
	"log"

	"github.com/biplob-codes/shelf-cli/cmd"
	"github.com/biplob-codes/shelf-cli/internal/db"
	_ "modernc.org/sqlite"
)

func main() {
	database,err:=db.Connect("shelf.db")
	if err !=nil{
		log.Fatalf("Database connection : %v",err)
	}
	if err:=db.Migrate(database);err!=nil{
		log.Fatalf("Database migration : %v",err)
	}
	
	cmd.Execute()
	defer database.Close()
}
