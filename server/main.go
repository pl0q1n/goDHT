package main

import (
	"flag"
	"log"
	"net"

	"golang.org/x/net/context"

	pbClient "github.com/pl0q1n/goDHT/client_proto"
	pbNode "github.com/pl0q1n/goDHT/node_proto"
	api "github.com/pl0q1n/goDHT/server/api"

	"google.golang.org/grpc"
)

func main() {
	mode := flag.String("mode", "new", "mode")
	connectTo := flag.String("target", "127.0.0.1:8080", "target host")
	host := flag.String("host", "127.0.0.1:8081", "host")
	flag.Parse()
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	api.GlobalNode.SetId(host)
	grpcServer := grpc.NewServer()
	pbNode.RegisterNodeServer(grpcServer, &api.NodeServer{})
	pbClient.RegisterKeyValueServer(grpcServer, &api.ClientServer{})
	if *mode == "join" {
		join := &pbNode.JoinRequest{
			Id:   api.GlobalNode.GetId(),
			Host: *host,
		}
		conn, err := grpc.Dial(*connectTo, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		cl := pbNode.NewNodeClient(conn)
		resp, err := cl.ProcessJoin(context.Background(), join)
		if err != nil {
			log.Fatalf("Put request failed: %v", err)
		}
		if resp.Status != pbNode.JoinResponse_Success {
			log.Printf("You did it wrong: %d", resp.Status)
		}
		api.GlobalNode.HashTable = resp.InitStorage
		fingerTable := api.GetFingerTableFromProto(resp.FingerTable)
		api.GlobalNode.FingerTableConn.UpdateChan <- *fingerTable
	}
	go api.ProcessConnections(api.GlobalNode)
	log.Printf("connected with mode: %s", *mode)
	log.Printf("Start with Self id: %d and self Host: %s", api.GlobalNode.GetId(), api.GlobalNode.FingerTable.SelfEntry.Host)
	grpcServer.Serve(listener)

}
