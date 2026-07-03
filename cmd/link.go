/*
Copyright © 2026 biplob-codes
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/biplob-codes/shelf-cli/internal/ui"
	"github.com/spf13/cobra"
)

var longLinkMessage = `Save and organize your links — with a tag and a collection.

Examples:
  shelf link add https://example.com -t news -c reading-list
  shelf link list -c reading-list
  shelf link update 3 tech-articles
  shelf link delete 3

Aliases: shelf lnk, shelf l`

func LinkCMD(repo *store.Repository) *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "link",
		Aliases: []string{"lnk", "l"},
		Short:   "Manage your saved links",
		Long:    longLinkMessage,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	addCmd := &cobra.Command{
		Use:   "add [url]",
		Short: "Save a new link",
		Long: `Save a new link, optionally tagging it and assigning it to a collection.

Example:
  shelf link add https://example.com -t news -c reading-list`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url := args[0]
			collection, err := cmd.Flags().GetString("collection")
			if err != nil {
				return fmt.Errorf("reading --collection flag: %w", err)
			}
			tag, err := cmd.Flags().GetString("tag")
			if err != nil {
				return fmt.Errorf("reading --tag flag: %w", err)
			}
			if err := repo.AddLink(url, tag, collection); err != nil {
				return fmt.Errorf("adding link %q: %w", url, err)
			}
			ui.PrintSuccess("Saved %s", url)
			return nil
		},
	}
	addCmd.Flags().StringP("collection", "c", "", "Assign the link to a collection")
	addCmd.Flags().StringP("tag", "t", "no-tag", "Tag the link for easier filtering")

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List saved links",
		Long: `List all saved links, optionally filtered by collection.

Example:
  shelf link list -c reading-list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			collection, err := cmd.Flags().GetString("collection")
			if err != nil {
				return fmt.Errorf("reading --collection flag: %w", err)
			}
			links, err := repo.GetLinks(collection)
			if err != nil {
				return fmt.Errorf("listing links: %w", err)
			}
			if len(links) == 0 {
				ui.PrintInfo("No links yet. Add one with 'shelf link add <url>'.")
				return nil
			}
			fmt.Println(ui.RenderLinksTable(links))
			return nil
		},
	}
	listCmd.Flags().StringP("collection", "c", "", "Filter links by collection")

	updateCmd := &cobra.Command{
		Use:   "update [id] [tag]",
		Short: "Update a link's tag",
		Long: `Update the tag on an existing link by its ID.

Example:
  shelf link update 3 tech-articles`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			tag := args[1]
			if err := repo.UpdateLink(id, tag); err != nil {
				return fmt.Errorf("updating link %d: %w", id, err)
			}
			ui.PrintSuccess("Updated link %d with tag %q", id, tag)
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a saved link",
		Long: `Permanently delete a saved link by its ID.

Example:
  shelf link delete 3`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			if err := repo.DeleteLink(id); err != nil {
				return fmt.Errorf("deleting link %d: %w", id, err)
			}
			ui.PrintSuccess("Deleted link %d", id)
			return nil
		},
	}

	linkCmd.AddCommand(addCmd, listCmd, updateCmd, deleteCmd)
	return linkCmd
}

func parseID(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid id %q: must be a number", s)
	}
	return n, nil
}