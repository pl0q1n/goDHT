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

func (node *Node) SetId(host *string) {
	hashSum := sha512.Sum512([]byte(*host))
	node.fingerTable.SelfEntry.Hash = SHAToUint64(hashSum)
}

func (node *Node) GetRange(start uint64, end uint64) map[uint64][]byte {
	rangeMap := make(map[uint64][]byte)
	for key, value := range node.hashTable {
		if key >= start && key < end {
			rangeMap[key] = value
		}
	}
	return rangeMap
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
		Hash: in.Id,
		Host: in.Host,
	}

	if in.Id < GlobalNode.id && in.Id > GlobalNode.fingerTable.PreviousEntry.Hash {
		status := pbNode.JoinResponse_WrongNode
		response := &pbNode.JoinResponse{}
		response.Status = status
		return response, nil
	}

	GlobalNode.fingerTable.Add(entry)
	protoFingerTable := GlobalNode.fingerTable.GetProtoFingerTable()

	response := &pbNode.JoinResponse{}
	response.FingerTable = protoFingerTable
	response.InitStorage = GlobalNode.GetRange(GlobalNode.fingerTable.PreviousEntry.Hash, in.Id)
	response.Status = pbNode.JoinResponse_Success

	return response, nil
}

func (s *NodeServer) ProcessFingerTable(ctx context.Context, in *pbNode.FingerTableRequest) (*pbNode.FingerTable, error) {
	response := GlobalNode.fingerTable.GetProtoFingerTable()
	return response, nil
}

func (s *NodeServer) ProcessRoute(ctx context.Context, in *pbNode.RouteRequest) (*pbNode.RouteResponse, error) {
	response := &pbNode.RouteResponse{}

	return response, nil
}
