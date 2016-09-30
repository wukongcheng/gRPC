// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	NoParam
	IPList
*/
package helloworld

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NoParam struct {
}

func (m *NoParam) Reset()                    { *m = NoParam{} }
func (m *NoParam) String() string            { return proto.CompactTextString(m) }
func (*NoParam) ProtoMessage()               {}
func (*NoParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type IPList struct {
	Ip []string `protobuf:"bytes,1,rep,name=ip" json:"ip,omitempty"`
}

func (m *IPList) Reset()                    { *m = IPList{} }
func (m *IPList) String() string            { return proto.CompactTextString(m) }
func (*IPList) ProtoMessage()               {}
func (*IPList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*NoParam)(nil), "helloworld.NoParam")
	proto.RegisterType((*IPList)(nil), "helloworld.IPList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Whitelist service

type WhitelistClient interface {
	GetWhitelist(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*IPList, error)
}

type whitelistClient struct {
	cc *grpc.ClientConn
}

func NewWhitelistClient(cc *grpc.ClientConn) WhitelistClient {
	return &whitelistClient{cc}
}

func (c *whitelistClient) GetWhitelist(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*IPList, error) {
	out := new(IPList)
	err := grpc.Invoke(ctx, "/helloworld.Whitelist/GetWhitelist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Whitelist service

type WhitelistServer interface {
	GetWhitelist(context.Context, *NoParam) (*IPList, error)
}

func RegisterWhitelistServer(s *grpc.Server, srv WhitelistServer) {
	s.RegisterService(&_Whitelist_serviceDesc, srv)
}

func _Whitelist_GetWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhitelistServer).GetWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Whitelist/GetWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhitelistServer).GetWhitelist(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Whitelist_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Whitelist",
	HandlerType: (*WhitelistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWhitelist",
			Handler:    _Whitelist_GetWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4b, 0x43, 0x31,
	0x10, 0x87, 0xfb, 0x54, 0xfa, 0x7c, 0x47, 0x51, 0x89, 0x20, 0xa1, 0x2e, 0x92, 0x41, 0x9c, 0x82,
	0xd4, 0xc9, 0xc5, 0xa1, 0x83, 0xb5, 0x20, 0x12, 0xea, 0xd0, 0x39, 0xea, 0xd1, 0x06, 0x12, 0x13,
	0x93, 0x88, 0xf6, 0xbf, 0x37, 0x89, 0x2d, 0x2f, 0x43, 0xb7, 0xfb, 0x5d, 0xbe, 0xdc, 0x7d, 0x1c,
	0x9c, 0xad, 0x51, 0x6b, 0xfb, 0x63, 0xbd, 0xfe, 0xe0, 0xce, 0xdb, 0x68, 0x09, 0xf4, 0x1d, 0xc6,
	0x60, 0xf4, 0x94, 0xd3, 0x02, 0xbf, 0xbe, 0x31, 0x44, 0x42, 0xe0, 0xe8, 0x53, 0x1a, 0xa4, 0xcd,
	0x55, 0x73, 0xd3, 0x2d, 0x4a, 0xcd, 0xae, 0x01, 0xb6, 0x8c, 0xd3, 0x1b, 0x42, 0xa1, 0x35, 0x18,
	0x82, 0x5c, 0xed, 0xa0, 0x5d, 0x64, 0x1d, 0xb4, 0x2f, 0x56, 0x48, 0x2f, 0x0d, 0xa3, 0x30, 0x9c,
	0x8b, 0x67, 0x95, 0x06, 0x9e, 0xc0, 0x81, 0x72, 0x89, 0x3c, 0x4c, 0x64, 0xaa, 0x26, 0x73, 0x68,
	0x67, 0x1e, 0x31, 0xa2, 0x27, 0x0f, 0x70, 0xfc, 0x2a, 0x37, 0x65, 0x34, 0xa1, 0xbc, 0xd2, 0xac,
	0x8d, 0xc6, 0x17, 0x7b, 0x5e, 0x92, 0x07, 0x1b, 0x4c, 0x1e, 0xa1, 0x5b, 0xae, 0x55, 0x44, 0x9d,
	0xf7, 0xdc, 0xc3, 0x68, 0x86, 0xb1, 0xcf, 0xe7, 0xf5, 0xb7, 0xad, 0xd6, 0x98, 0xd4, 0xcd, 0x7f,
	0x41, 0x36, 0x98, 0xde, 0xc2, 0xa5, 0xb2, 0x7c, 0xe5, 0xdd, 0x3b, 0xc7, 0x5f, 0x69, 0x9c, 0xc6,
	0x50, 0x71, 0xd3, 0xd3, 0xb2, 0x74, 0x99, 0x6b, 0x91, 0xef, 0x27, 0x9a, 0xb7, 0x61, 0x39, 0xe4,
	0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0xcc, 0x02, 0xc8, 0x5c, 0x01, 0x00, 0x00,
}
