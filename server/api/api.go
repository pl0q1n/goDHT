package api

import (
	"encoding/binary"
	"log"

	"context"
	"crypto/sha512"
	"fmt"

	pbClient "github.com/pl0q1n/goDHT/client_proto"
	pbNode "github.com/pl0q1n/goDHT/node_proto"
)

// Node...Brief implimentation of node structure for DHT
type Node struct {
	HashTable       map[uint64][]byte
	FingerTable     FingerTable
	FingerTableConn FingerTableConnection
	start           uint64
	end             uint64
	id              uint64
}

var GlobalNode *Node = &Node{
	HashTable: make(map[uint64][]byte),
	FingerTableConn: FingerTableConnection{
		UpdateChan: make(chan FingerTable, 1),
	},
	FingerTable: FingerTable{
		PreviousEntry: Entry{
			Host: "",
			Hash: 0,
		},
		SelfEntry: Entry{
			Host: "",
			Hash: 0,
		},
	},
}

func SHAToUint64(hash [64]byte) uint64 {
	return binary.BigEndian.Uint64(hash[:8])
}

func (node *Node) GetId() uint64 {
	return node.FingerTable.SelfEntry.Hash
}

func (node *Node) SetId(host *string) {
	hashSum := sha512.Sum512([]byte(*host))
	node.FingerTable.SelfEntry.Host = *host
	node.FingerTable.SelfEntry.Hash = SHAToUint64(hashSum)
	log.Printf("Log from SetId. HOST: %S, ID: %D", *host, node.FingerTable.SelfEntry.Hash)
}

func (node *Node) GetRange(start uint64, end uint64) map[uint64][]byte {
	rangeMap := make(map[uint64][]byte)
	for key, value := range node.HashTable {
		if key >= start && key < end {
			rangeMap[key] = value
		}
	}
	return rangeMap
}

func (node *Node) ProcessGet(request *pbClient.GetRequest) *pbClient.GetResponse {
	response := &pbClient.GetResponse{}
	key := SHAToUint64(sha512.Sum512(request.Key))
	value, ok := node.HashTable[key]
	if !ok {
		response.Status = 1 // I don't get how to take "value-name" of enum from pb
	} else {
		response.Status = 0
	}
	response.Value = value
	return response
}

func (node *Node) ProcessDelete(request *pbClient.DeleteRequest) *pbClient.DeleteResponse {
	response := &pbClient.DeleteResponse{}
	key := SHAToUint64(sha512.Sum512(request.Key))
	_, ok := node.HashTable[key]
	if ok {
		response.Status = 0
		delete(node.HashTable, key)
	} else {
		response.Status = 1
	}
	return response
}

func (node *Node) ProcessPut(request *pbClient.PutRequest) *pbClient.PutResponse {
	response := &pbClient.PutResponse{}

	// check that Key is null
	if len(request.Key) == 0 {
		response.Status = 3
		return response
	}
	key := SHAToUint64(sha512.Sum512(request.Key))
	_, exist := node.HashTable[key]
	if exist {
		response.Status = 1
	} else {
		response.Status = 0
		// temp if for server_tests. Should create mock or something to avoid this runtime check
		if node.HashTable == nil {
			node.HashTable = make(map[uint64][]byte)
		}
		node.HashTable[key] = request.Value
		//temp print, just to know that everything is alright with client's PUT
		fmt.Printf("added to node with next args: key: %d, value: %s \n", key, request.Value)
	}
	return response
}

type ClientServer struct{}

// These methods are not thread-safe (TODO)
// I'm not sure about error handling here (nothing to handle)
func (s *ClientServer) ProcessGet(ctx context.Context, in *pbClient.GetRequest) (*pbClient.GetResponse, error) {
	log.Println("starting Process GET")
	key := SHAToUint64(sha512.Sum512(in.Key))
	host, ind := GlobalNode.FingerTable.Route(key)
	if host != GlobalNode.FingerTable.SelfEntry.Host {
		log.Printf("Route Get to node with host: %s", host)
		cl := pbClient.NewKeyValueClient(GlobalNode.FingerTableConn.Connections[ind].ClientConn)
		response, err := cl.ProcessGet(context.Background(), in)
		if err != nil {
			log.Fatalln("something wrong with get routing")
		}
		return response, nil
	}
	return GlobalNode.ProcessGet(in), nil
}

func (s *ClientServer) ProcessPut(ctx context.Context, in *pbClient.PutRequest) (*pbClient.PutResponse, error) {
	log.Println("starting Process PUT")
	key := SHAToUint64(sha512.Sum512(in.Key))
	host, ind := GlobalNode.FingerTable.Route(key)
	if host != GlobalNode.FingerTable.SelfEntry.Host {
		log.Printf("Route Put to node with host: %s", host)
		cl := pbClient.NewKeyValueClient(GlobalNode.FingerTableConn.Connections[ind].ClientConn)
		response, err := cl.ProcessPut(context.Background(), in)
		if err != nil {
			log.Fatalln("something wrong with put routing")
		}
		return response, nil
	}
	return GlobalNode.ProcessPut(in), nil
}

func (s *ClientServer) ProcessDelete(ctx context.Context, in *pbClient.DeleteRequest) (*pbClient.DeleteResponse, error) {
	log.Println("starting Process DELETE")
	key := SHAToUint64(sha512.Sum512(in.Key))
	host, ind := GlobalNode.FingerTable.Route(key)
	if host != GlobalNode.FingerTable.SelfEntry.Host {
		log.Printf("Route Delete to node with host: %s", host)
		cl := pbClient.NewKeyValueClient(GlobalNode.FingerTableConn.Connections[ind].ClientConn)
		response, err := cl.ProcessDelete(context.Background(), in)
		if err != nil {
			log.Fatalln("something wrong with delete routing")
		}
		return response, nil
	}
	return GlobalNode.ProcessDelete(in), nil
}

type NodeServer struct {
}

func (s *NodeServer) ProcessJoin(ctx context.Context, in *pbNode.JoinRequest) (*pbNode.JoinResponse, error) {
	log.Println("starting Process JOIN")
	var entry *Entry = &Entry{
		Hash: in.Id,
		Host: in.Host,
	}

	host, ind := GlobalNode.FingerTable.Route(in.Id)
	log.Printf("in ID here: %d", in.Id)
	log.Printf("Route result: %s", host)
	if host != GlobalNode.FingerTable.SelfEntry.Host {
		log.Printf("Route join to node with host: %s", host)
		response, err := GlobalNode.FingerTableConn.Connections[ind].Client.ProcessJoin(context.Background(), in)
		if err != nil {
			log.Fatalln("something wrong with join routing")
		}
		return response, nil
	}

	tempFingerTable := FingerTable{}
	tempFingerTable.Entries[0] = *entry
	GlobalNode.FingerTableConn.UpdateChan <- tempFingerTable
	protoFingerTable := GlobalNode.FingerTable.GetProtoFingerTable()

	response := &pbNode.JoinResponse{}
	response.FingerTable = protoFingerTable
	response.InitStorage = GlobalNode.GetRange(GlobalNode.FingerTable.PreviousEntry.Hash, in.Id)
	response.Status = pbNode.JoinResponse_Success

	return response, nil
}

func (s *NodeServer) ProcessFingerTable(ctx context.Context, in *pbNode.FingerTableRequest) (*pbNode.FingerTable, error) {
	response := GlobalNode.FingerTable.GetProtoFingerTable()
	return response, nil
}
