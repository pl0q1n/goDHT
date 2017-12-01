package api

import (
	"encoding/binary"

	"context"
	"crypto/sha512"
	"fmt"

	pbClient "github.com/pl0q1n/goDHT/client_proto"
	pbNode "github.com/pl0q1n/goDHT/node_proto"
)

// Node...Brief implimentation of node structure for DHT
type Node struct {
	hashTable   map[uint64][]byte
	fingerTable FingerTable
	start       uint64
	end         uint64
	id          uint64
}

var GlobalNode *Node = &Node{
	hashTable: make(map[uint64][]byte),
	fingerTable: FingerTable{
		start: 0,
		selfEntry: Entry{
			host: "",
			hash: 0,
		},
	},
}

func SHAToUint64(hash [64]byte) uint64 {
	return binary.BigEndian.Uint64(hash[:8])
}

func (node *Node) SetId(host *string) {
	hashSum := sha512.Sum512([]byte(*host))
	node.fingerTable.selfEntry.hash = SHAToUint64(hashSum)
}

func (node *Node) ProcessGet(request *pbClient.GetRequest) *pbClient.GetResponse {
	response := &pbClient.GetResponse{}
	key := SHAToUint64(sha512.Sum512(request.Key))
	value, ok := node.hashTable[key]
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
	_, ok := node.hashTable[key]
	if ok {
		response.Status = 0
		delete(node.hashTable, key)
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
	_, exist := node.hashTable[key]
	if exist {
		response.Status = 1
	} else {
		response.Status = 0
		// temp if for server_tests. Should create mock or something to avoid this runtime check
		if node.hashTable == nil {
			node.hashTable = make(map[uint64][]byte)
		}
		node.hashTable[key] = request.Value
		//temp print, just to know that everything is alright with client's PUT
		fmt.Printf("added to node with next args: key: %d, value: %s \n", key, request.Value)
	}
	return response
}

type ClientServer struct{}

// These methods are not thread-safe (TODO)
// I'm not sure about error handling here (nothing to handle)
func (s *ClientServer) ProcessGet(ctx context.Context, in *pbClient.GetRequest) (*pbClient.GetResponse, error) {
	fmt.Println("starting Process GET")
	return GlobalNode.ProcessGet(in), nil
}

func (s *ClientServer) ProcessPut(ctx context.Context, in *pbClient.PutRequest) (*pbClient.PutResponse, error) {
	fmt.Println("starting Process PUT")
	return GlobalNode.ProcessPut(in), nil
}

func (s *ClientServer) ProcessDelete(ctx context.Context, in *pbClient.DeleteRequest) (*pbClient.DeleteResponse, error) {
	fmt.Println("starting Process DELETE")
	return GlobalNode.ProcessDelete(in), nil
}

type NodeServer struct {
}

func (s *NodeServer) ProcessJoin(ctx context.Context, in *pbNode.JoinRequest) (*pbNode.JoinResponse, error) {
	var entry *Entry = &Entry{
		hash: in.Id,
		host: in.Host,
	}

	GlobalNode.fingerTable.add(entry)
	response := &pbNode.JoinResponse{}

	return response, nil
}

func (s *NodeServer) ProcessFingerTable(ctx context.Context, in *pbNode.FingerTableRequest) (*pbNode.FingerTable, error) {
	response := &pbNode.FingerTable{}

	return response, nil
}

func (s *NodeServer) ProcessRoute(ctx context.Context, in *pbNode.RouteRequest) (*pbNode.RouteResponse, error) {
	response := &pbNode.RouteResponse{}

	return response, nil
}
