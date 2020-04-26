package cmd

import (
	"chat-server/chat"
	"context"
	"log"
	"time"

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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		err := chatClient.limiter.Wait(ctx)
		if err != nil {
			return err
		}

		pingResponse, err := chatClient.call.Ping(ctx, &chat.PingRequest{}, chatClient.callOpts...)
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
