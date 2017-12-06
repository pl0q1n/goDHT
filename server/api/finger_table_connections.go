package api

import (
	"log"
	"time"

	pb "github.com/pl0q1n/goDHT/node_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Connection struct {
	Host       string
	Anime      chan struct{}
	Client     pb.NodeClient
	UpdateChan chan FingerTable
	ClientConn *grpc.ClientConn
}

type FingerTableConnection struct {
	Connections [64]Connection
	UpdateChan  chan FingerTable
}

func (conn *Connection) run() {
	for {
		fingerTableProto, err := conn.Client.ProcessFingerTable(context.Background(),
			&pb.FingerTableRequest{})
		if err != nil {
			log.Println("TODO: (Failover) notify fingertable about connection failing to node")
		}
		fingerTable := GetFingerTableFromProto(fingerTableProto)
		conn.UpdateChan <- *fingerTable
		select {
		case <-conn.Anime:
			conn.ClientConn.Close()
			return
		case <-time.After(5 * time.Second):
			continue
		}

	}
}

func (conn *FingerTableConnection) ProcessUpdate(update *Update) {
	for ind, elem := range update.updates {
		if elem == "" {
			continue
		}
		if conn.Connections[ind].Host == "" {
			connect, err := grpc.Dial(elem, grpc.WithInsecure())
			if err != nil {
				log.Printf("did not connect: %v", err)
				continue
			}
			conn.Connections[ind] = Connection{
				Host:       elem,
				Anime:      make(chan struct{}),
				Client:     pb.NewNodeClient(connect),
				UpdateChan: conn.UpdateChan,
				ClientConn: connect,
			}

			go conn.Connections[ind].run()
		}
	}
}
