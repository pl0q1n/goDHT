[![Build Status](https://travis-ci.org/pl0q1n/goDHT.svg?branch=master)](https://travis-ci.org/pl0q1n/goDHT)

# goDHT
A chord based distributed hash table (DHT) using Go & gRPC

## Prerequisites
* [grpc-go](https://github.com/grpc/grpc-go)


## Build
### Server:
```
$ protoc --go_out=plugins=grpc:. ./client_proto/client.proto
$ go build ./client
```
### Client:
```
$ protoc --go_out=plugins=grpc:. ./node_proto/node.proto
$ go build ./server
```

## Usage
### Server:
```
General options:
  -help                       Show help
  -host [ ip:port ]           Create host  
  -taget [ ip:port ]          Create host and connect it to existing one  
```

#### Example:
```
 $ ./server -host 127.0.0.1:8080
```

### Client:
I wrote this client just to give a simple example of how the server could be used. You can easily write your own client, just follow client protocol (node_proto/node.proto)

```
General options:
  -help                       Show help
  -host [ ip:port ]           Target host
  -action [ args ]            Action selection
  -key [ integer ]            Key for your value (key-value model)
  -value [ integer ]          Value for your key 

Args for action:
   p                          Put request
   d                          Delete request
   g                          Get request
```
#### Example:
```
$ ./client -host localhost:8080 -action p -key 1700 -value 1337
```

