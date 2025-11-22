package cmd

import (
	"cli/services"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Send a snapshot of your workspace to our agent backend and document your code",
	Long:  "sends workspace meta-data to agents layer to help create a beautiful readme and matching documentation",
	Run: func(cmd *cobra.Command, args []string) {
		services.GenerateService()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
