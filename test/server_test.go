package DHT_test

import (
	"crypto/sha512"
	"testing"
	server "../src"
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
	testRequest := &server.GetRequest{
		Key: 1337,
	}

	response := node.GetProcessing(*testRequest)
	if response.Status != 1 {
		t.Errorf("wrong GetResponse_Status: %d", response.Status)
	}
}

func TestPutProcessingSuccess(t *testing.T) {
	node := &server.Node{}
	testRequest := &server.PutRequest{
		Value: "PutProcessing test",
	}

	response := node.PutProcessing(*testRequest)
	if response.Status != 0 {
		t.Errorf("invalid GetResponse_Status. Expected: %d, but got: %d", 0, response.Status)
	}
}

func TestGetProcessingSuccess(t *testing.T){
	node := &server.Node{}
	testPutRequest := &server.PutRequest{
		Value: "PutRequest for GetRequest",
	}
	testPutResponse := node.PutProcessing(*testPutRequest)
	key := testPutResponse.Key
	testGetRequest := &server.GetRequest{
		Key: key,
	}
	testResponse := node.GetProcessing(*testGetRequest)
	if testResponse.Value != testPutRequest.Value {
		t.Errorf("Wrong Value. Expected: %s, but got: %s", testPutRequest.Value, testResponse.Value)
	}
}

func TestPutProcessingAlreadyExist(t *testing.T) {
	node := &server.Node{}
	testRequest := &server.PutRequest{
		Value: "PutProcessing test",
	}

	node.PutProcessing(*testRequest)
	response := node.PutProcessing(*testRequest)
	if response.Status != 1 {
		t.Errorf("invalid PutProcessing_Status. Expected: %d, but got: %d", 1, response.Status)
	}
}

func TestDeleteProcessingNotFound(t *testing.T) {
	node := &server.Node{}
	testRequest := &server.DeleteRequest{
		1337,
	}
	response := node.DeleteProcessing(*testRequest)
	if response.Status != 1 {
		t.Errorf("invalid DeleteResponse_Status. Expected: %d, but got: %d", 1, response.Status)
	}
}

func TestDeleteProcessingSuccess(t *testing.T){
	node := &server.Node{}
	testPutRequest := &server.PutRequest{
		Value: "DeleteProcessing test",
	}

	testPutResponse := node.PutProcessing(*testPutRequest)
	testDeleteRequest := &server.DeleteRequest{
		Key: testPutResponse.Key,
	}
	response := node.DeleteProcessing(*testDeleteRequest)
	if response.Status != 0 {
		t.Errorf("invalid DeleteResponse_Status. Expected: %d, but got: %d", 0, response.Status)
	}

	testGetRequest := &server.GetRequest{
		Key: 1337,
	}

	testGetResponse := node.GetProcessing(*testGetRequest)
	if testGetResponse.Status != 1 {
		t.Errorf("wrong GetResponse_Status: %d", testGetResponse.Status)
	}
}