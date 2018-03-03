package main

import (
	"flag"
	"log"

	pb "github.com/pl0q1n/goDHT/client_proto"

	"encoding/binary"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func create_put(key *int, value *string) *pb.PutRequest {
	bkey := make([]byte, 4)
	binary.LittleEndian.PutUint32(bkey, uint32(*key))
	request := &pb.PutRequest{
		Key:   bkey,
		Value: []byte(*value),
	}
	return request
}

func create_delete(key *int) *pb.DeleteRequest {
	bkey := make([]byte, 4)
	binary.LittleEndian.PutUint32(bkey, uint32(*key))
	request := &pb.DeleteRequest{
		Key: bkey,
	}
	return request
}

func create_get(key *int) *pb.GetRequest {
	bkey := make([]byte, 4)
	binary.LittleEndian.PutUint32(bkey, uint32(*key))
	request := &pb.GetRequest{
		Key: bkey,
	}
	return request
}

func main() {
	action := flag.String("action", "p", "action type")
	key := flag.Int("key", 1337, "key for request")
	address := flag.String("host", "localhost:8081", "host to connect")
	value := flag.String("value", "1480", "value for put request")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := pb.NewKeyValueClient(conn)
	if *action == "p" {
		request := create_put(key, value)
		r, err := cl.ProcessPut(context.Background(), request)
		if err != nil {
			log.Fatalf("Put request failed: %v", err)
		}
		log.Printf("Status of PutResponse: %d", r.Status)

	} else if *action == "d" {
		request := create_delete(key)
		r, err := cl.ProcessDelete(context.Background(), request)
		if err != nil {
			log.Fatalf("Delete request failed: %v", err)
		}
		log.Printf("Status of DeleteResponse: %d", r.Status)

	} else if *action == "g"{
		request := create_get(key)
		r, err := cl.ProcessGet(context.Background(), request)
		if err != nil {
			log.Fatalf("Get request failed: %v", err)
		}

		log.Printf("Status of GetResponse: %d", r.Status)
		log.Printf("Value of GetResponse: %s", string(r.Value[:]))

	} else {
		log.Printf("Wrong action parameter: %s", *action)
	}
}
