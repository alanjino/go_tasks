package main

import (
	"context"
	"log"
	"net"

	chat "github.com/go_tasks/tasks-go/grpc/chat"

	"google.golang.org/grpc"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, msg *chat.MessageReq) (*chat.MessageRes, error) {
	log.Printf("got the message from client Msg:%s", msg.Body)
	return &chat.MessageRes{Body: "Msg received"}, nil
}

func (s *Server) SayBye(ctx context.Context, msg *chat.MessageReqBye) (*chat.MessageResBye, error) {
	log.Printf("got the message from client Msg:%s", msg.Body)
	return &chat.MessageResBye{Body: "Msg received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to open port :50052 for listen, error: %v", err)
	}

	s := Server{}
	gServer := grpc.NewServer()
	chat.RegisterChatServiceServer(gServer, &s)

	if err := gServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serv GRPC Server over port 50052, %v", err)
	}
}
