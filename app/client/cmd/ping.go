package cmd

import (
	"chat-server/chat"
	"context"
	"log"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the server",
	Long:  `Ping the server`,
	RunE:  runPingCmd(),
}

func runPingCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		pingResponse, err := chatClient.Ping(context.TODO(), &chat.PingRequest{})
		if err != nil {
			return err
		}
		log.Printf("pingResponse: %s\n", pingResponse.Pong)

		return nil
	}
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
