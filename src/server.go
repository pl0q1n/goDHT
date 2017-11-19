package DHT

import (
	proto "../DHT_proto"
	"net"
	"fmt"
	"flag"
	"bufio"
	"encoding/binary"
	"crypto/sha512"
)


var (
	host *string = flag.String("host", "127.0.0.1:8081", "host")
)


type (
	GetRequest proto.GetRequest
	DeleteRequest proto.DeleteRequest
	PutRequest proto.PutRequest

	GetResponse proto.GetResponse
	DeleteResponse proto.DeleteResponse
	PutResponse proto.PutResponse

	GetStatus proto.GetResponse_Status
	DeleteStatus proto.DeleteResponse_Status
	PutStatus proto.PutResponse_Status
)


type Node struct {
	hashTable map[uint64] string
	start uint64
	end uint64
	id uint64
}


func SHAToUint64(hash [64]byte) uint64 {
	return binary.BigEndian.Uint64(hash[:8])
}


func (node *Node) GetProcessing(request GetRequest) *GetResponse {
	response := &GetResponse{}
	value, ok := node.hashTable[request.Key]
	if !ok {
		response.Status = 1  // I don't get how to take "value-name" of enum from proto
	} else {
		response.Status = 0
	}
	response.Value = value
	return response
	}


func (node *Node) DeleteProcessing(request DeleteRequest) *DeleteResponse {
	response := &DeleteResponse{}
	_, ok := node.hashTable[request.Key]
	if ok {
		response.Status = 0
		delete(node.hashTable, request.Key)
	} else {
		response.Status = 1
	}
	return response
}


func (node *Node) PutProcessing(request PutRequest) *PutResponse {
	response := &PutResponse{}
	valueBytes := []byte(request.Value)
	key := SHAToUint64(sha512.Sum512(valueBytes))
	_, ok := node.hashTable[key]
	if ok {
		response.Status = 1
	} else {
		response.Key = key
		response.Status = 0
		if node.hashTable == nil{
			node.hashTable = make(map[uint64]string)
		}
		node.hashTable[key] = request.Value
	}
	return response
}

func main() {
	flag.Parse()
	ln, _  := net.Listen("tcp", *(host))
	conn, _ := ln.Accept()

	for  {
		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Print("Message Received:", string(message))

	}


}