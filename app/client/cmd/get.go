package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [rooms]",
	Short: "get various resources",
	Long:  `can retrieve various resources from the server`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
