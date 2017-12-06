package DHT_test

import (
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"

	proto "github.com/golang/protobuf/proto"
	pbClient "github.com/pl0q1n/goDHT/client_proto"
	pbNode "github.com/pl0q1n/goDHT/node_proto"
	server "github.com/pl0q1n/goDHT/server/api"
)

func getTestFingerTable() *server.FingerTable {
	fingerTable := &server.FingerTable{}
	prevEntry := &server.Entry{
		Host: "127.13.37.0",
		Hash: 1337,
	}
	selfEntry := &server.Entry{
		Host: "127.14.88.0",
		Hash: 1480,
	}

	for i := 0; i < len(fingerTable.Entries); i++ {
		fingerTable.Entries[i].Hash = uint64(i * 30)
		fingerTable.Entries[i].Host = fmt.Sprintf("127.0.0.%d", i)
	}

	fingerTable.PreviousEntry = *prevEntry
	fingerTable.SelfEntry = *selfEntry

	return fingerTable
}

func assertEqual(expected interface{}, actual interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assertion failed. Expected: %v, but got: %v", expected, actual)
	}
}

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
	testRequest := &pbClient.GetRequest{
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
	testRequest := &pbClient.PutRequest{
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
	testPutRequest := &pbClient.PutRequest{
		Value: []byte("PutRequest for GetRequest"),
		Key:   key,
	}
	node.ProcessPut(testPutRequest)
	testGetRequest := &pbClient.GetRequest{
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
	testRequest := &pbClient.PutRequest{
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
	testRequest := &pbClient.DeleteRequest{
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
	testPutRequest := &pbClient.PutRequest{
		Value: []byte("DeleteProcessing test"),
		Key:   key,
	}

	node.ProcessPut(testPutRequest)
	testDeleteRequest := &pbClient.DeleteRequest{
		Key: key,
	}
	response := node.ProcessDelete(testDeleteRequest)
	if response.Status != 0 {
		t.Errorf("invalid DeleteResponse_Status. Expected: %d, but got: %d", 0, response.Status)
	}

	testGetRequest := &pbClient.GetRequest{
		Key: key,
	}

	testGetResponse := node.ProcessGet(testGetRequest)
	if testGetResponse.Status != 1 {
		t.Errorf("wrong GetResponse_Status: %d", testGetResponse.Status)
	}
}

func TestGetProtoFingerTable(t *testing.T) {
	fingerTable := getTestFingerTable()
	protoFingerTable := &pbNode.FingerTable{}
	protoEntry := &pbNode.FingerTable_Entry{}

	protoEntry.Hash = fingerTable.PreviousEntry.Hash
	protoEntry.Host = fingerTable.PreviousEntry.Host
	protoFingerTable.Previous = protoEntry

	var entrySlice []*pbNode.FingerTable_Entry

	for _, elem := range fingerTable.Entries {
		protoEntry = &pbNode.FingerTable_Entry{}
		protoEntry.Hash = elem.Hash
		protoEntry.Host = elem.Host
		entrySlice = append(entrySlice, protoEntry)
	}

	protoFingerTable.Entry = entrySlice

	testProtoFingerTable := fingerTable.GetProtoFingerTable()

	if !proto.Equal(protoFingerTable, testProtoFingerTable) {
		t.Error("FingerTable messages are not equal")
	}
}

func TestRoute(t *testing.T) {
	fingerTable := getTestFingerTable()

	//Primitive one
	assertEqual("127.0.0.4", fingerTable.Route(90), t)

	//Over 0
	assertEqual("127.0.0.0", fingerTable.Route(1900), t)

	//Last id
	assertEqual("127.0.0.0", fingerTable.Route(1890), t)

	//self
	assertEqual(fingerTable.SelfEntry.Host,
		fingerTable.Route(fingerTable.SelfEntry.Hash-1), t)

	tempHash := fingerTable.SelfEntry.Hash
	fingerTable.SelfEntry.Hash = fingerTable.PreviousEntry.Hash
	fingerTable.PreviousEntry.Hash = tempHash

	//selfhash = 1338, prev = 1480
	assertEqual(fingerTable.SelfEntry.Host, fingerTable.Route(1500), t)
}

func TestAdd(t *testing.T) {
	fingerTable := &server.FingerTable{
		PreviousEntry: server.Entry{
			Host: "127.13.37.0",
			Hash: 1337,
		},
		SelfEntry: server.Entry{
			Host: "127.14.88.0",
			Hash: 1480,
		},
	}
	firstEntry := &server.Entry{
		Host: "00.00.00.0",
		Hash: 1900,
	}

	for _, elem := range fingerTable.Entries {
		assertEqual(server.Entry{}, elem, t)
	}

	fingerTable.Add(firstEntry)

	for _, elem := range fingerTable.Entries {
		assertEqual(*firstEntry, elem, t)
	}

	secondEntry := &server.Entry{
		Host: "11.11.11.11",
		Hash: 2900,
	}

	fingerTable.Add(secondEntry)
	assertEqual(secondEntry.Host, fingerTable.Route(1905), t)
	assertEqual(firstEntry.Host, fingerTable.Route(secondEntry.Hash), t)

}
