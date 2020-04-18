package cmd

import (
	"chat-server/chat"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addUserCmd represents the user command
var addUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Add a user",
	Long:  `Adds a user to the server`,
	RunE:  runAddUserCmd(),
}

func runAddUserCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("user called with args: %+v\n", args)

		if len(args) != 1 {
			return errors.New("must provide a name for the user")
		}

		user, err := chatClient.CreateUser(context.TODO(), &chat.CreateUserRequest{Name: args[0]})
		if err != nil {
			return err
		}

		log.Printf("new user: %+v\n", user.User)
		return nil
	}
}

func init() {
	addCmd.AddCommand(addUserCmd)
}
