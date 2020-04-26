package main

import (
	"chat-server/chat"
	"context"
	"fmt"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	listener = netutil.LimitListener(listener, 10)
	chatServer := NewChatServer(10)
	srv := grpc.NewServer(chatServer.srvOpts...)
	chat.RegisterChatServer(srv, chatServer)
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

type ChatServer struct {
	users map[string]*chat.User

	roomsMtx sync.Mutex
	rooms    map[string]*chat.Room

	messengers map[string]*chat.RoomMessenger

	srvOpts []grpc.ServerOption
}

func (c *ChatServer) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("unaryInterceptor before handler: %s\n", info.FullMethod)
	resp, err = handler(ctx, req)
	log.Printf("unaryInterceptor after handler: %s\n", info.FullMethod)
	return resp, err
}

func (c *ChatServer) streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("streamInterceptor before handler: %s\n", info.FullMethod)
	err := handler(srv, ss)
	log.Printf("streamInterceptor after handler: %s\n", info.FullMethod)
	return err
}

func (c *ChatServer) rateLimiter(ctx context.Context, info *tap.Info) (context.Context, error) {
	log.Printf("rateLimiter tap: %s\n", info.FullMethodName)
	// is authenticated user? check in ctx
	return ctx, nil
}

func (c *ChatServer) StreamRoom(request *chat.StreamRoomRequest, stream chat.Chat_StreamRoomServer) error {
	room, ok := c.rooms[request.Name]
	if !ok {
		return status.Error(codes.NotFound, "room not found")
	}

	log.Printf("found room: %+v\n", room)

	id, msgStream := c.messengers[room.Name].Subscribe()
	defer c.messengers[room.Name].Unsubscribe(id)

	for {
		select {
		case msg := <-msgStream:
			log.Printf("msg: %s\n", msg)
			err := stream.Send(&chat.StreamRoomResponse{Message: msg})
			if err != nil {
				log.Printf("err sending to client: %v\n", err)
				return err
			}
		case <-stream.Context().Done():
			log.Printf("client cancelled stream")
			return nil
		}
	}
}

func (c *ChatServer) CreateRoom(ctx context.Context, request *chat.CreateRoomRequest) (*chat.CreateRoomResponse, error) {
	c.roomsMtx.Lock()
	defer c.roomsMtx.Unlock()

	room := chat.NewRoom(request.Name)
	c.rooms[room.Name] = room

	messenger := chat.NewRoomMessenger(room)
	c.messengers[room.Name] = messenger

	// todo allow clients to publish to server
	go func() {
		for {
			for _, rm := range c.messengers {
				rm.I <- fmt.Sprintf("message: %s", time.Now().String())
			}

			time.Sleep(500 * time.Millisecond)
		}
	}()

	return &chat.CreateRoomResponse{
		Room: room,
	}, nil
}

func (c *ChatServer) Ping(ctx context.Context, request *chat.PingRequest) (*chat.PingResponse, error) {
	log.Println("Ping")
	if rand.Intn(10) > 5 {
		return nil, status.Errorf(codes.Unavailable, "maybeFailRequest: failing it")
	}

	return &chat.PingResponse{
		Pong: "pong",
	}, nil
}

func (c *ChatServer) CreateUser(ctx context.Context, request *chat.CreateUserRequest) (*chat.CreateUserResponse, error) {
	return &chat.CreateUserResponse{
		User: chat.NewUser(request.Name),
	}, nil
}

func (c *ChatServer) GetRooms(ctx context.Context, request *chat.GetRoomsRequest) (*chat.GetRoomsResponse, error) {
	return &chat.GetRoomsResponse{
		Rooms: c.rooms,
	}, nil
}

func NewChatServer(maxStreams uint32) *ChatServer {
	cs := new(ChatServer)
	cs.Init(maxStreams)
	return cs
}

func (c *ChatServer) Init(maxStreams uint32) {
	c.rooms = make(map[string]*chat.Room)
	c.users = make(map[string]*chat.User)
	c.messengers = make(map[string]*chat.RoomMessenger)
	c.srvOpts = append(c.srvOpts, grpc.StreamInterceptor(c.streamInterceptor), grpc.UnaryInterceptor(c.unaryInterceptor), grpc.InTapHandle(c.rateLimiter), grpc.MaxConcurrentStreams(maxStreams))
}
