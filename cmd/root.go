/*
Copyright © 2026 biplob-codes 
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Persistent flags go here — available to all subcommands.
	// Example: rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shelf.yaml)")
}