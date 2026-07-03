/*
Copyright © 2026 biplob-codes
*/
package cmd

import (
	"fmt"

	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/biplob-codes/shelf-cli/internal/ui"
	"github.com/spf13/cobra"
)

func CollectionCMD(repo *store.Repository) *cobra.Command {
	collectionCmd := &cobra.Command{
		Use:     "collection",
		Aliases: []string{"c"},
		Short:   "Manage your link collections",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	createCmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new collection",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if err := repo.CreateCollection(name); err != nil {
				return fmt.Errorf("creating collection %q: %w", name, err)
			}
			ui.PrintSuccess("Created collection %q", name)
			return nil
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all collections",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			collections, err := repo.ReadCollections()
			if err != nil {
				return fmt.Errorf("listing collections: %w", err)
			}
			if len(collections) == 0 {
				ui.PrintInfo("No collections yet. Create one with 'shelf collection create <name>'.")
				return nil
			}
			fmt.Println(ui.RenderCollectionsTable(collections))
			return nil
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update [old-name] [new-name]",
		Short: "Rename a collection",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			oldName, newName := args[0], args[1]
			if err := repo.UpdateCollection(oldName, newName); err != nil {
				return fmt.Errorf("renaming collection %q to %q: %w", oldName, newName, err)
			}
			ui.PrintSuccess("Renamed collection %q to %q", oldName, newName)
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a collection",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if err := repo.DeleteCollection(name); err != nil {
				return fmt.Errorf("deleting collection %q: %w", name, err)
			}
			ui.PrintSuccess("Deleted collection %q", name)
			return nil
		},
	}

	collectionCmd.AddCommand(createCmd, listCmd, updateCmd, deleteCmd)
	return collectionCmd
}