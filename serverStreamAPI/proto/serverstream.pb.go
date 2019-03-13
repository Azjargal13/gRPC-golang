// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/serverstream.proto

/*
Package serverstream is a generated protocol buffer package.

It is generated from these files:
	proto/serverstream.proto

It has these top-level messages:
	GreetManyRequest
	GreetManyResponse
	Greeting
*/
package serverstream

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

type GreetManyRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetManyRequest) Reset()                    { *m = GreetManyRequest{} }
func (m *GreetManyRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetManyRequest) ProtoMessage()               {}
func (*GreetManyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GreetManyRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManyResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *GreetManyResponse) Reset()                    { *m = GreetManyResponse{} }
func (m *GreetManyResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetManyResponse) ProtoMessage()               {}
func (*GreetManyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetManyResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Greeting struct {
	Firstname string `protobuf:"bytes,1,opt,name=firstname" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname" json:"lastname,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Greeting) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Greeting) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetManyRequest)(nil), "serverstream.GreetManyRequest")
	proto.RegisterType((*GreetManyResponse)(nil), "serverstream.GreetManyResponse")
	proto.RegisterType((*Greeting)(nil), "serverstream.Greeting")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GreetingService service

type GreetingServiceClient interface {
	Greet(ctx context.Context, in *GreetManyRequest, opts ...grpc.CallOption) (GreetingService_GreetClient, error)
}

type greetingServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetingServiceClient(cc *grpc.ClientConn) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) Greet(ctx context.Context, in *GreetManyRequest, opts ...grpc.CallOption) (GreetingService_GreetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GreetingService_serviceDesc.Streams[0], c.cc, "/serverstream.GreetingService/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceGreetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetingService_GreetClient interface {
	Recv() (*GreetManyResponse, error)
	grpc.ClientStream
}

type greetingServiceGreetClient struct {
	grpc.ClientStream
}

func (x *greetingServiceGreetClient) Recv() (*GreetManyResponse, error) {
	m := new(GreetManyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GreetingService service

type GreetingServiceServer interface {
	Greet(*GreetManyRequest, GreetingService_GreetServer) error
}

func RegisterGreetingServiceServer(s *grpc.Server, srv GreetingServiceServer) {
	s.RegisterService(&_GreetingService_serviceDesc, srv)
}

func _GreetingService_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetingServiceServer).Greet(m, &greetingServiceGreetServer{stream})
}

type GreetingService_GreetServer interface {
	Send(*GreetManyResponse) error
	grpc.ServerStream
}

type greetingServiceGreetServer struct {
	grpc.ServerStream
}

func (x *greetingServiceGreetServer) Send(m *GreetManyResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverstream.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _GreetingService_Greet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/serverstream.proto",
}

func init() { proto.RegisterFile("proto/serverstream.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x2a, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x03,
	0x0b, 0x09, 0xf1, 0x20, 0x8b, 0x29, 0xb9, 0x71, 0x09, 0xb8, 0x17, 0xa5, 0xa6, 0x96, 0xf8, 0x26,
	0xe6, 0x55, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x71, 0x71, 0xa4, 0x83, 0xc4,
	0x32, 0xf3, 0xd2, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xc4, 0xf4, 0x50, 0x0c, 0x72, 0x87,
	0xca, 0x06, 0xc1, 0xd5, 0x29, 0x69, 0x73, 0x09, 0x22, 0x99, 0x53, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x36, 0x86, 0x33, 0x08, 0xca,
	0x53, 0x72, 0xe1, 0xe2, 0x80, 0x19, 0x21, 0x24, 0xc3, 0xc5, 0x99, 0x96, 0x59, 0x54, 0x5c, 0x92,
	0x97, 0x98, 0x9b, 0x0a, 0x55, 0x86, 0x10, 0x10, 0x92, 0xe2, 0xe2, 0xc8, 0x49, 0x84, 0x4a, 0x32,
	0x81, 0x25, 0xe1, 0x7c, 0xa3, 0x58, 0x2e, 0x7e, 0x98, 0x29, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9,
	0xa9, 0x42, 0x5e, 0x5c, 0xac, 0x60, 0x21, 0x21, 0x39, 0x2c, 0x0e, 0x46, 0xf2, 0xa2, 0x94, 0x3c,
	0x4e, 0x79, 0x88, 0xd3, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0xc1, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x40, 0x22, 0x49, 0x1a, 0x4a, 0x01, 0x00, 0x00,
}