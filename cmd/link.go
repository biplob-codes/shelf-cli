/*
Copyright © 2026 biplob-codes
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/spf13/cobra"
)

func getNumber(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

var longLinkMessage = `Save and organize your links — with tags and collections.

Examples:
  shelf link add https://example.com -t news -c reading-list
  shelf link list -c reading-list
  shelf link update 3 tech-articles
  shelf link delete 3

You can also use the shorter aliases:
  shelf lnk list
  shelf l list`

func LinkCMD(repo *store.Repository) *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "link",
		Aliases: []string{"lnk", "l"},
		Short:   "Manage your saved links",
		Long:    longLinkMessage,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	addCmd := &cobra.Command{
		Use:   "add [url]",
		Short: "Add a new link",
		Long: `Save a new link, optionally tagging it and assigning it to a collection.

Example:
  shelf link add https://example.com -t news -c reading-list`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			collection, err := cmd.Flags().GetString("collection")
			if err != nil {
				log.Fatalf("Error while getting flags: %v", err)
			}
			tag, err := cmd.Flags().GetString("tag")
			if err != nil {
				log.Fatalf("Error while getting flags: %v", err)
			}
			if err := repo.AddLink(args[0], tag, collection); err != nil {
				log.Fatalf("Add Link Command: %v", err)
			}
		},
	}
	addCmd.Flags().StringP("collection", "c", "", "Add links to your collection")
	addCmd.Flags().StringP("tag", "t", "no-tag", "Add tag to your links")

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List links in a collection",
		Long: `List all saved links, optionally filtered by collection.

Example:
  shelf link list -c reading-list`,
		Run: func(cmd *cobra.Command, args []string) {
			collection, err := cmd.Flags().GetString("collection")
			if err != nil {
				log.Fatalf("Error while getting flags: %v", err)
			}
			links, err := repo.GetLinks(collection)
			if err != nil {
				log.Fatalf("List Link Command: %v", err)
			}
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tURL\tTAGS")
			for _, link := range links {
				fmt.Fprintf(w, "%d\t%s\t%s\n", link.ID, link.URL, link.Tag)
			}
			w.Flush()
		},
	}
	listCmd.Flags().StringP("collection", "c", "", "Add links to your collection")

	updateCmd := &cobra.Command{
		Use:   "update [id] [tag]",
		Short: "Update a link's tag",
		Long: `Update the tag on an existing link by its ID.

Example:
  shelf link update 3 tech-articles`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			tag := args[1]
			if err := repo.UpdateLink(getNumber(id), tag); err != nil {
				log.Fatalf("Update link command: %v", err)
			}
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a link",
		Long: `Permanently delete a saved link by its ID.

Example:
  shelf link delete 3`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := getNumber(args[0])
			if err := repo.DeleteLink(id); err != nil {
				log.Fatalf("Delete link command: %v", err)
			}
		},
	}

	linkCmd.AddCommand(addCmd)
	linkCmd.AddCommand(listCmd)
	linkCmd.AddCommand(updateCmd)
	linkCmd.AddCommand(deleteCmd)
	return linkCmd
}
