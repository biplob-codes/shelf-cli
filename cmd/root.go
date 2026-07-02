/*
Copyright © 2026 biplob-codes 
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Use:   "shelf",
	Short: "Organize and manage your links from the terminal",
	Long: `Shelf is a CLI tool for saving, organizing, and retrieving links.

Group your links into collections, tag them for easy filtering,
and quickly open or search through them without leaving your terminal.

Examples:
  shelf add https://example.com --collection work --tags golang,cli
  shelf list --collection work
  shelf open <link-id>`,
}

