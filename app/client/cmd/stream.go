package cmd

import (
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream [room]",
	Short: "Stream from  a room",
	Long:  `Stream from a room`,
}

func init() {
	rootCmd.AddCommand(streamCmd)
}
