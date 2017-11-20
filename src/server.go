package DHT

import (
	proto "../DHT_proto"
	"bufio"
	"crypto/sha512"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
)

var (
	host *string = flag.String("host", "127.0.0.1:8081", "host")
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

func (node *Node) ProcessGet(request proto.GetRequest) *proto.GetResponse {
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

func (node *Node) ProcessDelete(request proto.DeleteRequest) *proto.DeleteResponse {
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

func (node *Node) ProcessPut(request proto.PutRequest) *proto.PutResponse {
	response := &proto.PutResponse{}
	valueBytes := []byte(request.Value)
	key := SHAToUint64(sha512.Sum512(valueBytes))
	_, ok := node.hashTable[key]
	if ok {
		response.Status = 1
	} else {
		response.Key = key
		response.Status = 0
		if node.hashTable == nil {
			node.hashTable = make(map[uint64]string)
		}
		node.hashTable[key] = request.Value
	}
	return response
}

func main() {
	flag.Parse()
	ln, _ := net.Listen("tcp", *(host))
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Print("Message Received:", string(message))

	}

}
