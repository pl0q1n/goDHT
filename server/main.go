package main

import (
	"flag"
	"log"
	"net"

	pb "goDHT/client_proto"
	api "goDHT/server/api"

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
	pb.RegisterNodeServer(grpcServer, &api.Server{})
	grpcServer.Serve(listener)
}
