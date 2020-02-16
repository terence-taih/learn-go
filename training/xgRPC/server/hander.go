package api

import (
	"context"
	"fmt"
)

type Server struct {

}

func (s *Server) SayHello(ctx context.Context,in *PingMessage) (*PingMessage, error){
	fmt.Println("Receive msg from client: " , in.Greeting)
	return &PingMessage{Greeting: "Nice to meet you"}, nil
}

