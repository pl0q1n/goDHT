package main

import (
	"flag"
	"log"
	"net"

	proto "github.com/pl0q1n/goDHT/proto"
	api "github.com/pl0q1n/goDHT/server/api"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "127.0.0.1:8081", "host")
	flag.Parse()

	listener, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNodeServer(grpcServer, &api.Server{})
	grpcServer.Serve(listener)
}