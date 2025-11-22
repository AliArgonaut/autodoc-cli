package cmd

import (
	"cli/services"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize auto-doc project and create auto-doc-config.json",
	Long:  `Sets up the folder structure, captures content, and generates configuration file with sensible defaults for a new documentation project.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.AutodocInitService()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
