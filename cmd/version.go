package cmd

import (
	"fmt"
	"onlyfans/internal/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("onlyfans", version.String())
	},
}
