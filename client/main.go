package main

import (
	"log"

	pb "github.com/pl0q1n/goDHT/client_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8081"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := pb.NewNodeClient(conn)
	r, err := cl.ProcessPut(context.Background(), &pb.PutRequest{Value: "1337"})
	if err != nil {
		log.Fatalf("Put request failed: %v", err)
	}
	log.Printf("Greeting: %d, %d", r.Key, r.Status)
}
