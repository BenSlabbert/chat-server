package chat

import (
	"github.com/google/uuid"
	"log"
	"time"
)

func NewUser(name string) *User {
	return &User{Id: uuid.New().String(), Name: name}
}

func NewRoom(name string) *Room {
	return &Room{Name: name, Users: make(map[string]*User)}
}

type RoomMessenger struct {
	room        *Room
	messages    []string
	I           chan string
	subscribers map[string]chan string
}

func (rm *RoomMessenger) Close() error {
	close(rm.I)
	return nil
}

func (rm *RoomMessenger) write() {
	lastIndex := 0
	for {
		if len(rm.messages) == 0 {
			<-time.After(250 * time.Millisecond)
			continue
		}

		if len(rm.subscribers) == 0 {
			<-time.After(250 * time.Millisecond)
			continue
		}

		for idx, msg := range rm.messages {
			for id, subs := range rm.subscribers {
				select {
				case <-time.After(250 * time.Millisecond):
					log.Println("subscriber took too long to read from chan")
					close(subs)
					delete(rm.subscribers, id)
				case subs <- msg:
					log.Println("subscriber read from chan")
				}
			}
			lastIndex = idx
		}

		rm.messages = rm.messages[lastIndex+1:]
	}
}

func (rm *RoomMessenger) read() {
	for in := range rm.I {
		rm.messages = append(rm.messages, in)
	}
}

func (rm *RoomMessenger) Unsubscribe(id string) {
	log.Printf("unsubscribing: %s", id)
	delete(rm.subscribers, id)
}

func (rm *RoomMessenger) Subscribe() (string, <-chan string) {
	newChan := make(chan string)
	id := uuid.New().String()
	rm.subscribers[id] = newChan
	return id, newChan
}

func NewRoomMessenger(room *Room) *RoomMessenger {
	rm := &RoomMessenger{
		room:        room,
		messages:    []string{},
		I:           make(chan string),
		subscribers: make(map[string]chan string),
	}

	go rm.read()
	go rm.write()

	return rm
}
