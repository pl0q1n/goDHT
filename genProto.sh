protoc --go_out=plugins=grpc:. ./node_proto/node.proto
protoc --go_out=plugins=grpc:. ./client_proto/client.proto
