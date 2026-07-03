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
	Long: `Shelf is a CLI tool for saving, organizing, and retrieving links without
ever leaving your terminal.

Group related links into collections, tag them for quick filtering, and
list or search through them whenever you need. Everything is stored
locally in a SQLite database, so your links stay fast, private, and yours.

Core concepts:
  Collections   Named groups of links, e.g. "work" or "reading-list"
  Links         URLs saved under a collection, optionally tagged
  Tags          A single label per link, used to filter results

Examples:
  shelf collection create work
  shelf link add https://example.com -c work -t golang
  shelf link list -c work
  shelf link update 3 tech-articles
  shelf collection list`,
	SilenceUsage:  true,
	SilenceErrors: true,
}