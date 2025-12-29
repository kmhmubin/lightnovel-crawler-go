package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lightnovel crawler go",
	Short: "A light novel crawler written in Go",
	Long:  `lightnovel-crawler-go is a command-line tool to crawl and download light novels from various websites and convert them to EPUB.`,

	// This function will be called when no subcommands are provided
	Run: func(cmd *cobra.Command, args []string) {
		// Default message shown when run the program without any subcommands
		fmt.Println("Welcome to lightnovel Crawler Go! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
