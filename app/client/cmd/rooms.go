package cmd

import (
	"chat-server/chat"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/spf13/cobra"
)

// getRoomsCmd represents the get room command
var getRoomsCmd = &cobra.Command{
	Use:   "rooms",
	Short: "get a list of chat rooms",
	Long:  `get all the available rooms`,
	RunE:  runGetRoomsCmd(),
}

func runGetRoomsCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Println("get rooms")

		chatRoomsResponse, err := chatClient.call.GetRooms(context.TODO(), &chat.GetRoomsRequest{}, chatClient.callOpts...)
		if err != nil {
			return err
		}

		log.Printf("chatrooms: %+v\n", chatRoomsResponse.Rooms)
		return nil
	}
}

// addRoomsCmd represents the get rooms command
var addRoomCmd = &cobra.Command{
	Use:   "room",
	Short: "get a list of chat rooms",
	Long:  `get all the available rooms`,
	RunE:  runAddRoomCmd(),
}

func runAddRoomCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("add room called with args: %+v\n", args)

		if len(args) != 1 {
			return errors.New("must provide a name for the room")
		}

		createRoomResponse, err := chatClient.call.CreateRoom(context.TODO(), &chat.CreateRoomRequest{Name: args[0]}, chatClient.callOpts...)
		if err != nil {
			//status.Code(err)
			return err
		}

		log.Printf("chatrooms: %+v\n", createRoomResponse.Room)
		return nil
	}
}

// addRoomsCmd represents the get rooms command
var streamRoomCmd = &cobra.Command{
	Use:   "room",
	Short: "stream chats from the room",
	Long:  `stream chats from the room`,
	RunE:  runStreamRoomCmd(),
}

func runStreamRoomCmd() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("stream room called with args: %+v\n", args)

		if len(args) != 1 {
			return errors.New("must provide a name for the room")
		}

		room := args[0]
		stream, err := chatClient.call.StreamRoom(context.TODO(), &chat.StreamRoomRequest{Name: room}, chatClient.callOpts...)
		if err != nil {
			return err
		}

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				log.Printf("%s: ### server closing stream ###\n", room)
				break
			}
			if err != nil {
				return err
			}

			log.Printf("%s: %s\n", room, msg)
		}

		return nil
	}
}

func init() {
	getCmd.AddCommand(getRoomsCmd)
	addCmd.AddCommand(addRoomCmd)
	streamCmd.AddCommand(streamRoomCmd)
}
