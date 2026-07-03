/*
Copyright © 2026 biplob-codes
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/spf13/cobra"
)

var longCollectionMessage=`var longCollectionMessage = Group related links into collections — like folders for your saved links.

Examples:
  shelf collection create horror-movies
  shelf collection list
  shelf collection update horror-movies -n classic-horror
  shelf collection delete horror-movies

You can also use the shorter aliases:
  shelf col list
  shelf c list`
func CollectionCMD(repo *store.Repository)*cobra.Command{
 collectionCmd := &cobra.Command{
	Use:     "collection",
	Aliases: []string{"col", "c"},
	Short:   "Manage your link collections",
	Long: longCollectionMessage,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
 createCmd := &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new collection",
	 Long: `Create a new collection to organize your saved links.

Example:
  shelf collection create reading-list`,
   Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
     if err:=repo.CreateCollection(args[0]); err!= nil{
			log.Fatalf("Create Collection Command: %v",err)
		}
	},
}
listCmd := &cobra.Command{
	Use:   "list",
	Short: "List all collections",
    Long: `List all your collections with link counts.

Example:
  shelf collection list`,
  Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		 collections,err:=repo.ReadCollections()
		 if err!=nil{
			log.Fatalf("List Collection Command: %v",err)
		 }
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTITLE\tLINKS")
		for _, collection := range collections {
			fmt.Fprintf(w, "%d\t%s\t%d\n", collection.ID, collection.Title, collection.NumberOfLinks)
		}
		w.Flush()
	},
}
updateCmd := &cobra.Command{
	Use:   "update [old-name] [new-name]",
	Short: "Update a collection",
	Long: `Update an existing collection's name.

Example:
  shelf collection update reading-list tech-articles`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldTitle:=args[0]
		newTitle:=args[1]
		if err:=repo.UpdateCollection(oldTitle,newTitle);err!=nil{
			log.Fatalf("Update collection command: %v",err)
		}
	},
}

deleteCmd := &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a collection",
	Long: `Delete a collection permanently.

Example:
  shelf collection delete reading-list`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title:=args[0]
		if err:=repo.DeleteCollection(title);err!=nil{
			log.Fatalf("Delete collection command: %v",err)
		}
	},
}
collectionCmd.AddCommand(createCmd)
collectionCmd.AddCommand(listCmd)
collectionCmd.AddCommand(updateCmd)
collectionCmd.AddCommand(deleteCmd)
return collectionCmd
}