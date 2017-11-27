package api

import (
	"encoding/binary"

	"context"
	"crypto/sha512"
	"fmt"

	pb "github.com/pl0q1n/goDHT/client_proto"
)

// Node..Brief implimentation of node structure for DHT
type Node struct {
	hashTable map[uint64][]byte
	start     uint64
	end       uint64
	id        uint64
}

var GlobalNode *Node = &Node{
	hashTable: make(map[uint64][]byte),
	start:     0,
	end:       0,
	id:        0,
}

func SHAToUint64(hash [64]byte) uint64 {
	return binary.BigEndian.Uint64(hash[:8])
}

func (node *Node) SetId(host *string) {
	hashSum := sha512.Sum512([]byte(*host))
	node.id = SHAToUint64(hashSum)
}

func (node *Node) ProcessGet(request *pb.GetRequest) *pb.GetResponse {
	response := &pb.GetResponse{}
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

func (node *Node) ProcessDelete(request *pb.DeleteRequest) *pb.DeleteResponse {
	response := &pb.DeleteResponse{}
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

func (node *Node) ProcessPut(request *pb.PutRequest) *pb.PutResponse {
	response := &pb.PutResponse{}

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

type Server struct{}

// I'm not sure about error handling here (nothing to handle)
func (s *Server) ProcessGet(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Println("starting Process GET")
	return GlobalNode.ProcessGet(in), nil
}

func (s *Server) ProcessPut(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
	fmt.Println("starting Process PUT")
	return GlobalNode.ProcessPut(in), nil
}

func (s *Server) ProcessDelete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	fmt.Println("starting Process DELETE")
	return GlobalNode.ProcessDelete(in), nil
}
