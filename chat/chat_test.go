package chat

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRoomMessenger(t *testing.T) {
	messenger := NewRoomMessenger(&Room{Name: "test"})
	defer messenger.Close()

	go func() {
		for i := 0; i < 10; i++ {
			messenger.I <- fmt.Sprintf("message: %d", i)
		}
	}()

	id, data := messenger.Subscribe()

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("messages from instance: %s\n", <-data)
	}

	fmt.Println("finished reading")
	time.Sleep(1 * time.Second)

	messenger.Unsubscribe(id)

	fmt.Println("done")
}
