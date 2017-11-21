package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	proto "../DHT_proto"
	"log"
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
	cl := proto.NewNodeClient(conn)
	cl.ProcessPut(context.Background(), &proto.PutRequest{Value: "1337"})
	//log.Printf("Greeting: %d, %d", r.Key, r.Status)


}