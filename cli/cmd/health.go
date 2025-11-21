package cmd

import (
	"cli/services"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "tests health and connectivity of auto-doc ai agent backends",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		services.ShowIntegrated() //calls services layer
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
