package cmd

import (
	"fmt"
	"os"

	"onlyfans/internal/logger"
	"onlyfans/internal/version"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var verbose bool
var debug bool

var rootCmd = &cobra.Command{
	Use:   "onlyfans",
	Short: "onlyfans CLI collects data from a logged-in OnlyFans account",
	Long:  "A CLI tool to collect data from a logged-in OnlyFans account. This is a scaffolded Cobra application.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")

	cobra.OnInitialize(initLogger, func() {
		logger.Logger().Info("Application started", zap.String("version", version.String()))
	})

	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(versionCmd)
}

func initLogger() {
	if debug {
		err := logger.Init()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
			os.Exit(1)
		}
	}
}
