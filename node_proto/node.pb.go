// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_proto/node.proto

/*
Package node_proto is a generated protocol buffer package.

It is generated from these files:
	node_proto/node.proto

It has these top-level messages:
	JoinRequest
	JoinResponse
	FingerTable
	FingerTableRequest
	RouteRequest
	RouteResponse
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
	JoinResponse_Success   JoinResponse_Status = 0
	JoinResponse_Failure   JoinResponse_Status = 101
	JoinResponse_WrongNode JoinResponse_Status = 102
)

var JoinResponse_Status_name = map[int32]string{
	0:   "Success",
	101: "Failure",
	102: "WrongNode",
}
var JoinResponse_Status_value = map[string]int32{
	"Success":   0,
	"Failure":   101,
	"WrongNode": 102,
}

func (x JoinResponse_Status) String() string {
	return proto.EnumName(JoinResponse_Status_name, int32(x))
}
func (JoinResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Add Status to messages
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
	Entry    []*FingerTable_Entry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
	Previous *FingerTable_Entry   `protobuf:"bytes,2,opt,name=previous" json:"previous,omitempty"`
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

func (m *FingerTable) GetPrevious() *FingerTable_Entry {
	if m != nil {
		return m.Previous
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

type RouteRequest struct {
	Hash uint64 `protobuf:"varint,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *RouteRequest) Reset()                    { *m = RouteRequest{} }
func (m *RouteRequest) String() string            { return proto.CompactTextString(m) }
func (*RouteRequest) ProtoMessage()               {}
func (*RouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RouteRequest) GetHash() uint64 {
	if m != nil {
		return m.Hash
	}
	return 0
}

type RouteResponse struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
}

func (m *RouteResponse) Reset()                    { *m = RouteResponse{} }
func (m *RouteResponse) String() string            { return proto.CompactTextString(m) }
func (*RouteResponse) ProtoMessage()               {}
func (*RouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RouteResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "node_proto.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "node_proto.JoinResponse")
	proto.RegisterType((*FingerTable)(nil), "node_proto.FingerTable")
	proto.RegisterType((*FingerTable_Entry)(nil), "node_proto.FingerTable.Entry")
	proto.RegisterType((*FingerTableRequest)(nil), "node_proto.FingerTableRequest")
	proto.RegisterType((*RouteRequest)(nil), "node_proto.RouteRequest")
	proto.RegisterType((*RouteResponse)(nil), "node_proto.RouteResponse")
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
	ProcessRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
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

func (c *nodeClient) ProcessRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := grpc.Invoke(ctx, "/node_proto.Node/ProcessRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	ProcessJoin(context.Context, *JoinRequest) (*JoinResponse, error)
	ProcessFingerTable(context.Context, *FingerTableRequest) (*FingerTable, error)
	ProcessRoute(context.Context, *RouteRequest) (*RouteResponse, error)
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

func _Node_ProcessRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ProcessRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node_proto.Node/ProcessRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ProcessRoute(ctx, req.(*RouteRequest))
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
		{
			MethodName: "ProcessRoute",
			Handler:    _Node_ProcessRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_proto/node.proto",
}

func init() { proto.RegisterFile("node_proto/node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x8d, 0x37, 0x1f, 0xd0, 0x71, 0x5a, 0xad, 0x46, 0x45, 0x5d, 0x22, 0x01, 0x91, 0xb9, 0x84,
	0xcb, 0x56, 0xdd, 0x1e, 0xa0, 0x1c, 0x38, 0x20, 0xb5, 0x12, 0x20, 0x01, 0x72, 0x90, 0x38, 0xa2,
	0x6d, 0xd7, 0x4d, 0x2d, 0xa2, 0x75, 0xb0, 0xbd, 0x95, 0xfa, 0xa7, 0xf8, 0x11, 0xfc, 0x14, 0x7e,
	0x09, 0xb2, 0xd7, 0x61, 0x0d, 0x64, 0xc5, 0x6d, 0x6c, 0xbf, 0x79, 0xcf, 0xf3, 0xde, 0xc0, 0x83,
	0x5a, 0x55, 0xe2, 0xcb, 0x46, 0x2b, 0xab, 0x8e, 0x5d, 0x99, 0xfb, 0x12, 0xa1, 0xbb, 0x66, 0x27,
	0x40, 0xdf, 0x2a, 0x59, 0x73, 0xf1, 0xad, 0x11, 0xc6, 0xe2, 0x01, 0x24, 0xb2, 0xca, 0xc8, 0x9c,
	0x2c, 0x46, 0x3c, 0x91, 0x15, 0x22, 0x8c, 0x6e, 0x94, 0xb1, 0x59, 0x32, 0x27, 0x8b, 0x3d, 0xee,
	0x6b, 0xf6, 0x23, 0x81, 0x69, 0xdb, 0x63, 0x36, 0xaa, 0x36, 0x02, 0xcf, 0x80, 0x5e, 0xcb, 0x7a,
	0x25, 0xf4, 0xa7, 0xf2, 0x72, 0x2d, 0x7c, 0x37, 0x2d, 0x8e, 0xf2, 0x4e, 0x25, 0xbf, 0xe8, 0x9e,
	0x79, 0x8c, 0xc5, 0x77, 0x40, 0x65, 0x2d, 0xed, 0xd2, 0x2a, 0x5d, 0xae, 0x44, 0x96, 0xcc, 0x87,
	0x0b, 0x5a, 0x3c, 0x8b, 0x5b, 0x63, 0xa5, 0xfc, 0x4d, 0x87, 0x3d, 0xaf, 0xad, 0xbe, 0xe3, 0x71,
	0x37, 0x3e, 0x87, 0x89, 0xb1, 0xa5, 0x6d, 0x4c, 0x36, 0x9c, 0x93, 0xc5, 0x41, 0xf1, 0xa4, 0x97,
	0x67, 0xe9, 0x61, 0x3c, 0xc0, 0x67, 0xaf, 0x20, 0xfd, 0x9b, 0x19, 0x53, 0x18, 0x7e, 0x15, 0x77,
	0xc1, 0x0a, 0x57, 0xe2, 0x21, 0x8c, 0x6f, 0xcb, 0x75, 0x23, 0xbc, 0x19, 0x53, 0xde, 0x1e, 0x5e,
	0x26, 0x2f, 0x08, 0x3b, 0x81, 0x49, 0xcb, 0x88, 0x14, 0xee, 0x2d, 0x9b, 0xab, 0x2b, 0x61, 0x4c,
	0x3a, 0x70, 0x87, 0x8b, 0x52, 0xae, 0x1b, 0x2d, 0x52, 0x81, 0xfb, 0xb0, 0xf7, 0x59, 0xab, 0x7a,
	0xf5, 0x5e, 0x55, 0x22, 0xbd, 0x66, 0xdf, 0x09, 0xd0, 0xc8, 0x15, 0x3c, 0x85, 0xb1, 0x70, 0xba,
	0x19, 0xf1, 0x16, 0x3c, 0xea, 0x71, 0x2f, 0x6f, 0xc7, 0x6e, 0xb1, 0x78, 0x06, 0xf7, 0x37, 0x5a,
	0xdc, 0x4a, 0xd5, 0x18, 0xff, 0xa9, 0xff, 0xf6, 0xfd, 0x86, 0xcf, 0x8e, 0x61, 0xdc, 0xce, 0xe9,
	0x12, 0x2e, 0xcd, 0x4d, 0x18, 0xd4, 0xd7, 0x3b, 0x53, 0x3f, 0x04, 0x8c, 0x53, 0x6c, 0xf7, 0x85,
	0x31, 0x98, 0x72, 0xd5, 0xd8, 0xed, 0x79, 0x17, 0x1b, 0x7b, 0x0a, 0xfb, 0x01, 0x13, 0xf6, 0x65,
	0x4b, 0x4f, 0x3a, 0xfa, 0xe2, 0x27, 0x81, 0x91, 0xb3, 0x06, 0x5f, 0x03, 0xfd, 0xa8, 0x95, 0x73,
	0xd0, 0x25, 0x86, 0x47, 0xff, 0x66, 0xe8, 0x95, 0x66, 0x59, 0x5f, 0xb8, 0x6c, 0x80, 0x1f, 0x00,
	0x03, 0x47, 0x6c, 0xf1, 0xe3, 0xbe, 0x8d, 0x0c, 0x8c, 0x7d, 0x1b, 0xcb, 0x06, 0x78, 0x0e, 0xd3,
	0x40, 0xe8, 0x27, 0xc1, 0x3f, 0xc4, 0x63, 0x03, 0x66, 0x0f, 0x77, 0xbc, 0x6c, 0xff, 0x75, 0x39,
	0xf1, 0xd7, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x64, 0x27, 0x32, 0x53, 0x98, 0x03, 0x00,
	0x00,
}
