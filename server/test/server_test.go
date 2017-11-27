package DHT_test

import (
	"crypto/sha512"
	"encoding/binary"
	"reflect"
	"testing"

	pb "github.com/pl0q1n/goDHT/client_proto"
	server "github.com/pl0q1n/goDHT/server/api"
)

func TestSHAToUint64(t *testing.T) {
	testString := "SHAToUint"
	var expectedResult uint64 = 4936774056552412463
	testBytes := []byte(testString)
	testHash := sha512.Sum512(testBytes)
	testOut := server.SHAToUint64(testHash)
	if testOut != expectedResult {
		t.Errorf("wrong SHA to UINT64 conversion, should be: %d, but got %d", expectedResult, testOut)
	}
}

func TestGetProcessingNotFound(t *testing.T) {
	node := &server.Node{}
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, 1337)
	testRequest := &pb.GetRequest{
		Key: key,
	}

	response := node.ProcessGet(testRequest)
	if response.Status != 1 {
		t.Errorf("wrong GetResponse_Status: %d", response.Status)
	}
}

func TestPutProcessingSuccess(t *testing.T) {
	node := &server.Node{}
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, 1337)
	testRequest := &pb.PutRequest{
		Value: []byte("PutProcessing test"),
		Key:   key,
	}

	response := node.ProcessPut(testRequest)
	if response.Status != 0 {
		t.Errorf("invalid GetResponse_Status. Expected: %d, but got: %d", 0, response.Status)
	}
}

func TestGetProcessingSuccess(t *testing.T) {
	node := &server.Node{}
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, 1337)
	testPutRequest := &pb.PutRequest{
		Value: []byte("PutRequest for GetRequest"),
		Key:   key,
	}
	node.ProcessPut(testPutRequest)
	testGetRequest := &pb.GetRequest{
		Key: key,
	}
	testResponse := node.ProcessGet(testGetRequest)
	if !reflect.DeepEqual(testResponse.Value, testPutRequest.Value) {
		t.Errorf("Wrong Value. Expected: %s, but got: %s", testPutRequest.Value, testResponse.Value)
	}
}

func TestPutProcessingAlreadyExist(t *testing.T) {
	node := &server.Node{}
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, 1337)
	testRequest := &pb.PutRequest{
		Value: []byte("PutProcessing test"),
		Key:   key,
	}

	node.ProcessPut(testRequest)
	response := node.ProcessPut(testRequest)
	if response.Status != 1 {
		t.Errorf("invalid PutProcessing_Status. Expected: %d, but got: %d", 1, response.Status)
	}
}

func TestDeleteProcessingNotFound(t *testing.T) {
	node := &server.Node{}
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, 1337)
	testRequest := &pb.DeleteRequest{
		key,
	}
	response := node.ProcessDelete(testRequest)
	if response.Status != 1 {
		t.Errorf("invalid DeleteResponse_Status. Expected: %d, but got: %d", 1, response.Status)
	}
}

func TestDeleteProcessingSuccess(t *testing.T) {
	node := &server.Node{}
	key := make([]byte, 4)
	binary.LittleEndian.PutUint32(key, 1337)
	testPutRequest := &pb.PutRequest{
		Value: []byte("DeleteProcessing test"),
		Key:   key,
	}

	node.ProcessPut(testPutRequest)
	testDeleteRequest := &pb.DeleteRequest{
		Key: key,
	}
	response := node.ProcessDelete(testDeleteRequest)
	if response.Status != 0 {
		t.Errorf("invalid DeleteResponse_Status. Expected: %d, but got: %d", 0, response.Status)
	}

	testGetRequest := &pb.GetRequest{
		Key: key,
	}

	testGetResponse := node.ProcessGet(testGetRequest)
	if testGetResponse.Status != 1 {
		t.Errorf("wrong GetResponse_Status: %d", testGetResponse.Status)
	}
}
