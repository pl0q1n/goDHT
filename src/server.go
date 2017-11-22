package DHT

import (
	proto "../DHT_proto"
	"crypto/sha512"
	"encoding/binary"
	"flag"
	"net"
	"context"
	"log"
	"google.golang.org/grpc"
	"fmt"
)

var (
	host *string = flag.String("host", "127.0.0.1:8081", "host")
	node *Node = &Node{
		hashTable: make(map[uint64]string),
		start: 0,
		end: 0,
		id: 0,
	}
)

type Node struct {
	hashTable map[uint64]string
	start     uint64
	end       uint64
	id        uint64
}

func SHAToUint64(hash [64]byte) uint64 {
	return binary.BigEndian.Uint64(hash[:8])
}

func (node *Node) ProcessGet(request *proto.GetRequest) *proto.GetResponse {
	response := &proto.GetResponse{}
	value, ok := node.hashTable[request.Key]
	if !ok {
		response.Status = 1 // I don't get how to take "value-name" of enum from proto
	} else {
		response.Status = 0
	}
	response.Value = value
	return response
}

func (node *Node) ProcessDelete(request *proto.DeleteRequest) *proto.DeleteResponse {
	response := &proto.DeleteResponse{}
	_, ok := node.hashTable[request.Key]
	if ok {
		response.Status = 0
		delete(node.hashTable, request.Key)
	} else {
		response.Status = 1
	}
	return response
}

func (node *Node) ProcessPut(request *proto.PutRequest) *proto.PutResponse {
	response := &proto.PutResponse{}
	valueBytes := []byte(request.Value)
	key := SHAToUint64(sha512.Sum512(valueBytes))
	_, ok := node.hashTable[key]
	if ok {
		response.Status = 1
	} else {
		response.Key = key
		response.Status = 0
		// temp if for server_tests. Should create mock or something to avoid this runtime check
		if node.hashTable == nil {
			node.hashTable = make(map[uint64]string)
		}
		node.hashTable[key] = request.Value
		//temp print, just to know that everything is alright with client's PUT
		fmt.Printf("added to node with next args: key: %d, value: %s \n", key, request.Value)
	}
	return response
}

type server struct{}

// I'm not sure about error handling here (nothing to handle)
func (s *server) ProcessGet(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	fmt.Println("starting Process GET")
	return node.ProcessGet(in), nil
}

func (s *server) ProcessPut(ctx context.Context, in *proto.PutRequest) (*proto.PutResponse, error) {
	fmt.Println("starting Process PUT")
	return node.ProcessPut(in), nil
}

func (s *server) ProcessDelete(ctx context.Context, in *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	fmt.Println("starting Process DELETE")
	return node.ProcessDelete(in), nil
}


func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", *(host))
	if err != nil {
		log.Fatalf("failed to lwisten: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNodeServer(grpcServer, &server{})
	grpcServer.Serve(ln)
}
