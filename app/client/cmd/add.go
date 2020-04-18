package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [user]",
	Short: "Add a resource",
	Long:  `Add user to the server`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
