package main

import (
	"fmt"
	"os"
	"google.golang.org/grpc"
	"net"
	api "training/xgRPC/server"
)

func main(){
	tcpListenner, err := net.Listen("tcp", ":7777")

	if err != nil {
		fmt.Println("Cannot open tcp connection ", err)
		os.Exit(1)
	}

	server := api.Server{}
	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &server)

	if err := grpcServer.Serve(tcpListenner); err != nil {
		fmt.Println("Fail to start server", err)
	}
}