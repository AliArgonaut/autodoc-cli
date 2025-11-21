package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auto-doc",
	Short: "auto-doc is a tool to automatically generate documentation and beautiful markdown files",
	Long: `auto-doc is a CLI tool designed to initialize and manage automated documentation workflows.
	auto-doc supports two commands, similar to git. For example:

  auto-doc init      Initialize a new documentation project
  auto-doc generate  Generate docs from source files`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
}
