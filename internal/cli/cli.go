// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cli provides the command-line interface used for interacting
// with the app management system. It handles parsing flags,
// dispatching commands, and coordinating with the underlying handling
// and storage layers.
package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/nullzeiger/linkliste/internal/handling"
	"github.com/nullzeiger/linkliste/internal/storage"
	"github.com/nullzeiger/linkliste/internal/tui"
	"github.com/nullzeiger/linkliste/internal/types"
	"github.com/spf13/cobra"
)

var (
	// Flags for the add command
	description string
	name        string
	link        string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "linkliste",
	Short: "A CLI tool for managing link entries",
	Long: `Linkliste is a command-line application for managing and organizing
your links with descriptions and names. It supports listing, adding,
deleting, and searching through your saved links.`,
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls", "all"},
	Short: "List all link entries",
	Long:  "Display all saved link entries with their details.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := storage.Create(); err != nil {
			fmt.Println("Error creating .links.json file:", err)
			return
		}

		entries, err := handling.All()
		if err != nil {
			fmt.Println("Error listing entries:", err)
			return
		}
		for _, e := range entries {
			fmt.Println(e)
		}
	},
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new link entry",
	Long:  "Create a new link entry with description, name, and URL.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := storage.Create(); err != nil {
			fmt.Println("Error creating .links.json file:", err)
			return
		}

		// Validate required fields
		if description == "" || name == "" || link == "" {
			fmt.Println("Error: --description, --name, and --link are all required.")
			cmd.Usage()
			os.Exit(1)
		}

		// Create new link entry
		newEntry := types.Link{
			Date:        time.Now().Local(),
			Description: description,
			Name:        name,
			Link:        link,
		}

		// Save the new entry
		if err := handling.Create(newEntry); err != nil {
			fmt.Println("Error adding entry:", err)
			return
		}

		fmt.Println("Entry added successfully.")
	},
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [index]",
	Aliases: []string{"rm", "remove"},
	Short: "Delete a link entry by index",
	Long:  "Remove a link entry from the list using its index number.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := storage.Create(); err != nil {
			fmt.Println("Error creating .links.json file:", err)
			return
		}

		var index int
		if _, err := fmt.Sscanf(args[0], "%d", &index); err != nil {
			fmt.Println("Error: index must be a number")
			return
		}

		ok, err := handling.Delete(index)
		if err != nil {
			fmt.Println("Error deleting entry:", err)
			return
		}
		if ok {
			fmt.Printf("Entry [%d] deleted successfully.\n", index)
		}
	},
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search entries by keyword",
	Long:  "Search through all link entries for matches containing the given keyword.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := storage.Create(); err != nil {
			fmt.Println("Error creating .links.json file:", err)
			return
		}

		keyword := args[0]
		matches, err := handling.Search(keyword)
		if err != nil {
			fmt.Println("Error searching:", err)
			return
		}

		// No results found
		if len(matches) == 0 {
			fmt.Println("No results found.")
			return
		}

		// Print matching entries
		for _, m := range matches {
			fmt.Println(m)
		}
	},
}

// tuiCmd represents the TUI command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the Terminal User Interface",
	Long:  "Start an interactive TUI for managing links.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := storage.Create(); err != nil {
			fmt.Println("Error creating .links.json file:", err)
			return
		}

		tui.RunTui()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add flags to the add command
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the link (required)")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the link (required)")
	addCmd.Flags().StringVarP(&link, "url", "u", "", "URL of the link (required)")
	
	// Mark flags as required
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("url")

	// Add all commands to root
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(tuiCmd)
}