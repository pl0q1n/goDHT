// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package node_proto is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	JoinRequest
	JoinResponse
	FingerTable
	FingerTableRequest
*/
package node_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JoinResponse_Status int32

const (
	JoinResponse_Success JoinResponse_Status = 0
	JoinResponse_Failure JoinResponse_Status = 101
)

var JoinResponse_Status_name = map[int32]string{
	0:   "Success",
	101: "Failure",
}
var JoinResponse_Status_value = map[string]int32{
	"Success": 0,
	"Failure": 101,
}

func (x JoinResponse_Status) String() string {
	return proto.EnumName(JoinResponse_Status_name, int32(x))
}
func (JoinResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type JoinRequest struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JoinRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *JoinRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type JoinResponse struct {
	FingerTable *FingerTable        `protobuf:"bytes,1,opt,name=fingerTable" json:"fingerTable,omitempty"`
	InitStorage map[uint64][]byte   `protobuf:"bytes,2,rep,name=initStorage" json:"initStorage,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status      JoinResponse_Status `protobuf:"varint,3,opt,name=status,enum=node_proto.JoinResponse_Status" json:"status,omitempty"`
}

func (m *JoinResponse) Reset()                    { *m = JoinResponse{} }
func (m *JoinResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()               {}
func (*JoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JoinResponse) GetFingerTable() *FingerTable {
	if m != nil {
		return m.FingerTable
	}
	return nil
}

func (m *JoinResponse) GetInitStorage() map[uint64][]byte {
	if m != nil {
		return m.InitStorage
	}
	return nil
}

func (m *JoinResponse) GetStatus() JoinResponse_Status {
	if m != nil {
		return m.Status
	}
	return JoinResponse_Success
}

type FingerTable struct {
	Entry []*FingerTable_Entry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *FingerTable) Reset()                    { *m = FingerTable{} }
func (m *FingerTable) String() string            { return proto.CompactTextString(m) }
func (*FingerTable) ProtoMessage()               {}
func (*FingerTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FingerTable) GetEntry() []*FingerTable_Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type FingerTable_Entry struct {
	Hash uint64 `protobuf:"varint,1,opt,name=hash" json:"hash,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *FingerTable_Entry) Reset()                    { *m = FingerTable_Entry{} }
func (m *FingerTable_Entry) String() string            { return proto.CompactTextString(m) }
func (*FingerTable_Entry) ProtoMessage()               {}
func (*FingerTable_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *FingerTable_Entry) GetHash() uint64 {
	if m != nil {
		return m.Hash
	}
	return 0
}

func (m *FingerTable_Entry) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type FingerTableRequest struct {
}

func (m *FingerTableRequest) Reset()                    { *m = FingerTableRequest{} }
func (m *FingerTableRequest) String() string            { return proto.CompactTextString(m) }
func (*FingerTableRequest) ProtoMessage()               {}
func (*FingerTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*JoinRequest)(nil), "node_proto.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "node_proto.JoinResponse")
	proto.RegisterType((*FingerTable)(nil), "node_proto.FingerTable")
	proto.RegisterType((*FingerTable_Entry)(nil), "node_proto.FingerTable.Entry")
	proto.RegisterType((*FingerTableRequest)(nil), "node_proto.FingerTableRequest")
	proto.RegisterEnum("node_proto.JoinResponse_Status", JoinResponse_Status_name, JoinResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	ProcessJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	ProcessFingerTable(ctx context.Context, in *FingerTableRequest, opts ...grpc.CallOption) (*FingerTable, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) ProcessJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := grpc.Invoke(ctx, "/node_proto.Node/ProcessJoin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ProcessFingerTable(ctx context.Context, in *FingerTableRequest, opts ...grpc.CallOption) (*FingerTable, error) {
	out := new(FingerTable)
	err := grpc.Invoke(ctx, "/node_proto.Node/ProcessFingerTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	ProcessJoin(context.Context, *JoinRequest) (*JoinResponse, error)
	ProcessFingerTable(context.Context, *FingerTableRequest) (*FingerTable, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_ProcessJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ProcessJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node_proto.Node/ProcessJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ProcessJoin(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ProcessFingerTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FingerTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ProcessFingerTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node_proto.Node/ProcessFingerTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ProcessFingerTable(ctx, req.(*FingerTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node_proto.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessJoin",
			Handler:    _Node_ProcessJoin_Handler,
		},
		{
			MethodName: "ProcessFingerTable",
			Handler:    _Node_ProcessFingerTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0x6d, 0xd2, 0x8f, 0xcb, 0x9d, 0x2d, 0x25, 0x0c, 0x85, 0x86, 0xc2, 0xbd, 0x37, 0xe4, 0x29,
	0xf7, 0x25, 0x62, 0xfa, 0xe0, 0xc7, 0x83, 0x0f, 0x82, 0x05, 0x15, 0x54, 0x52, 0xdf, 0x25, 0x6d,
	0xc6, 0x76, 0xb1, 0x64, 0x6b, 0x76, 0x23, 0xf4, 0xb7, 0xf8, 0x33, 0xfc, 0x83, 0x92, 0xdd, 0x40,
	0x17, 0x35, 0x6f, 0x67, 0x92, 0x73, 0xce, 0xcc, 0x9c, 0x59, 0x80, 0x42, 0xe4, 0x14, 0xef, 0x4a,
	0xa1, 0x04, 0x6a, 0xfc, 0xa4, 0x71, 0x78, 0x0c, 0xec, 0x46, 0xf0, 0x22, 0xa5, 0xd7, 0x8a, 0xa4,
	0xc2, 0x11, 0xb8, 0x3c, 0xf7, 0x9d, 0xc0, 0x89, 0x7a, 0xa9, 0xcb, 0x73, 0x44, 0xe8, 0x6d, 0x84,
	0x54, 0xbe, 0x1b, 0x38, 0xd1, 0xef, 0x54, 0xe3, 0xf0, 0xc3, 0x85, 0xa1, 0xd1, 0xc8, 0x9d, 0x28,
	0x24, 0xe1, 0x19, 0xb0, 0x67, 0x5e, 0xac, 0xa9, 0x7c, 0xcc, 0x96, 0x5b, 0xd2, 0x6a, 0x96, 0x4c,
	0xe2, 0x43, 0x97, 0x78, 0x7e, 0xf8, 0x9d, 0xda, 0x5c, 0xbc, 0x05, 0xc6, 0x0b, 0xae, 0x16, 0x4a,
	0x94, 0xd9, 0x9a, 0x7c, 0x37, 0xe8, 0x46, 0x2c, 0xf9, 0x6f, 0x4b, 0xed, 0x4e, 0xf1, 0xf5, 0x81,
	0x7b, 0x55, 0xa8, 0x72, 0x9f, 0xda, 0x6a, 0x3c, 0x81, 0x81, 0x54, 0x99, 0xaa, 0xa4, 0xdf, 0x0d,
	0x9c, 0x68, 0x94, 0xfc, 0x6b, 0xf5, 0x59, 0x68, 0x5a, 0xda, 0xd0, 0xa7, 0x17, 0xe0, 0x7d, 0x75,
	0x46, 0x0f, 0xba, 0x2f, 0xb4, 0x6f, 0xa2, 0xa8, 0x21, 0x8e, 0xa1, 0xff, 0x96, 0x6d, 0x2b, 0xd2,
	0x61, 0x0c, 0x53, 0x53, 0x9c, 0xbb, 0xa7, 0x4e, 0x18, 0xc2, 0xc0, 0x38, 0x22, 0x83, 0x5f, 0x8b,
	0x6a, 0xb5, 0x22, 0x29, 0xbd, 0x4e, 0x5d, 0xcc, 0x33, 0xbe, 0xad, 0x4a, 0xf2, 0x28, 0x94, 0xc0,
	0xac, 0x14, 0x70, 0x06, 0x7d, 0xaa, 0xfb, 0xf8, 0x8e, 0x5e, 0xf9, 0x4f, 0x4b, 0x5a, 0xb1, 0x59,
	0xd3, 0x70, 0xa7, 0x47, 0xd0, 0x37, 0xc3, 0xd5, 0x67, 0xc9, 0xe4, 0xa6, 0x99, 0x4e, 0xe3, 0x1f,
	0x4f, 0x35, 0x06, 0xb4, 0xa3, 0x37, 0x47, 0x4e, 0xde, 0x1d, 0xe8, 0xdd, 0x89, 0x9c, 0xf0, 0x12,
	0xd8, 0x43, 0x29, 0xea, 0x69, 0xeb, 0x74, 0x70, 0xf2, 0x3d, 0x2f, 0x2d, 0x98, 0xfa, 0x6d, 0x41,
	0x86, 0x1d, 0xbc, 0x07, 0x6c, 0x3c, 0xec, 0xf5, 0xfe, 0xb6, 0x5d, 0xbf, 0x71, 0x6c, 0x7b, 0x1d,
	0x61, 0x67, 0x39, 0xd0, 0x1f, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x44, 0x6b, 0x3a,
	0xb2, 0x02, 0x00, 0x00,
}