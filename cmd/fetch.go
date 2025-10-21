package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data from the logged-in OnlyFans account",
	Long:  "Fetch various data (profile, posts, messages) from a logged-in OnlyFans account. Requires authentication setup.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetch called - stub implementation. Add authentication and scraping logic.")
	},
}
