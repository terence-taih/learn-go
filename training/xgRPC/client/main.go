package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"

	api "training/xgRPC/server"
)

func main() {
	 conn, err := grpc.Dial(":7777", grpc.WithInsecure())

	 if err != nil {
	 	fmt.Println("Cannot open connection to server", err)
	 }

	 defer conn.Close()

	 c := api.NewPingClient(conn)
	 message, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Hi server, I'm client"})

	 if err != nil {

	 }

	 fmt.Println("Reponse from server:", message.Greeting)

}