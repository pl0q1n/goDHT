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

func ProcessConnections(node *Node) {
	for {
		fingerTable := <-node.FingerTableConn.UpdateChan
		for _, elem := range fingerTable.Entries {
			if elem.Host == "" {
				continue
			}
			update := node.FingerTable.Add(&elem)
			node.FingerTableConn.ProcessUpdate(&update)
		}
		update := node.FingerTable.Add(&fingerTable.SelfEntry)
		node.FingerTableConn.ProcessUpdate(&update)
	}
}

func (conn *Connection) run() {
	for {
		fingerTableProto, err := conn.Client.ProcessFingerTable(context.Background(),
			&pb.FingerTableRequest{})
		if err != nil {
			log.Println("TODO: (Failover) notify fingertable about connection failing to node")
			return
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
		if elem != conn.Connections[ind].Host && conn.Connections[ind].Host != "" {
			var empty struct{}
			conn.Connections[ind].Anime <- empty
			log.Printf("Connection closed: %s", conn.Connections[ind].Host)
		}
		if conn.Connections[ind].Host == "" || elem != conn.Connections[ind].Host {
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
			log.Printf("Connection opened: %s", elem)
			go conn.Connections[ind].run()
		}
	}
}
