package cmd

import (
	"cli/services"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "tests health and connectivity of auto-doc ai agent backends",
	Long:  `this pings the backend with a get request to confirm connectivity. Reccomended for troubleshooting`,
	Run: func(cmd *cobra.Command, args []string) {
		services.HealthCheck() //calls services layer
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
