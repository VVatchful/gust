package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gust",
	Short: "The Gust interpreter",
	Long:  ``,
	// NOTE TO SELF: add REPL here
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {}
