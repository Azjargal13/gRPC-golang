// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/protobuf.proto

/*
Package calculator is a generated protocol buffer package.

It is generated from these files:
	proto/protobuf.proto

It has these top-level messages:
	SumRequest
	SumResponse
	PrimeNumberDecompositionRequest
	PrimeNumberDecompositionResponse
*/
package calculator

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

type SumRequest struct {
	FNum int32 `protobuf:"varint,1,opt,name=f_num,json=fNum" json:"f_num,omitempty"`
	SNum int32 `protobuf:"varint,2,opt,name=s_num,json=sNum" json:"s_num,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SumRequest) GetFNum() int32 {
	if m != nil {
		return m.FNum
	}
	return 0
}

func (m *SumRequest) GetSNum() int32 {
	if m != nil {
		return m.SNum
	}
	return 0
}

type SumResponse struct {
	SumRes int32 `protobuf:"varint,1,opt,name=sum_res,json=sumRes" json:"sum_res,omitempty"`
}

func (m *SumResponse) Reset()                    { *m = SumResponse{} }
func (m *SumResponse) String() string            { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()               {}
func (*SumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SumResponse) GetSumRes() int32 {
	if m != nil {
		return m.SumRes
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *PrimeNumberDecompositionRequest) Reset()                    { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string            { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()               {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PrimeNumberDecompositionRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	Primefactor int64 `protobuf:"varint,1,opt,name=primefactor" json:"primefactor,omitempty"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *PrimeNumberDecompositionResponse) GetPrimefactor() int64 {
	if m != nil {
		return m.Primefactor
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CalculatorService service

type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := grpc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CalculatorService_serviceDesc.Streams[0], c.cc, "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CalculatorService service

type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/protobuf.proto",
}

func init() { proto.RegisterFile("proto/protobuf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xad, 0x73, 0x15, 0x6e, 0x4f, 0x46, 0xd9, 0xc6, 0x5e, 0x2c, 0x79, 0x10, 0x41, 0xa9,
	0xa2, 0x20, 0xfa, 0xec, 0x9e, 0x8b, 0xb4, 0x1f, 0x60, 0xb4, 0xe1, 0x0a, 0x85, 0x5d, 0x53, 0x73,
	0x39, 0xc1, 0x6f, 0xe8, 0xc7, 0x92, 0xa5, 0xab, 0x1b, 0xc8, 0xd0, 0x97, 0x90, 0xfc, 0xef, 0x7e,
	0xf9, 0xe7, 0xfe, 0x81, 0x8b, 0xce, 0x59, 0x6f, 0xef, 0xc2, 0x5a, 0x49, 0x9d, 0x86, 0x8d, 0x02,
	0x53, 0xae, 0x8d, 0xac, 0x4b, 0x6f, 0x9d, 0x7e, 0x02, 0x28, 0x84, 0x72, 0x7c, 0x17, 0x64, 0xaf,
	0xce, 0x61, 0x5c, 0xaf, 0x5a, 0xa1, 0x79, 0x94, 0x44, 0xd7, 0xe3, 0xfc, 0xa4, 0xce, 0x84, 0x36,
	0x22, 0x07, 0xf1, 0xb8, 0x17, 0x39, 0x13, 0xd2, 0x57, 0x30, 0x09, 0x1c, 0x77, 0xb6, 0x65, 0x54,
	0x33, 0x38, 0x65, 0xa1, 0x95, 0x43, 0xde, 0xa2, 0x31, 0x87, 0xaa, 0x7e, 0x81, 0xcb, 0x37, 0xd7,
	0x10, 0x66, 0x42, 0x15, 0xba, 0x25, 0x1a, 0x4b, 0x9d, 0xe5, 0xc6, 0x37, 0xb6, 0x1d, 0x4c, 0xa7,
	0x10, 0xb7, 0xa1, 0x3a, 0xa0, 0xfd, 0x49, 0x2f, 0x21, 0x39, 0x8c, 0x6e, 0x7d, 0x13, 0x98, 0x74,
	0x9b, 0x9e, 0xba, 0x34, 0xde, 0xf6, 0x17, 0x8c, 0xf2, 0x7d, 0xe9, 0xe1, 0x2b, 0x82, 0xb3, 0xd7,
	0x9f, 0x79, 0x0b, 0x74, 0x1f, 0x8d, 0x41, 0xf5, 0x0c, 0xa3, 0x42, 0x48, 0x4d, 0xd3, 0x5d, 0x14,
	0xe9, 0x2e, 0x87, 0xc5, 0xec, 0x97, 0xde, 0xfb, 0xe9, 0x23, 0xf5, 0x09, 0xf3, 0x43, 0xaf, 0x52,
	0x37, 0xfb, 0xd8, 0x1f, 0x63, 0x2f, 0x6e, 0xff, 0xd7, 0x3c, 0x18, 0xdf, 0x47, 0x55, 0x1c, 0xbe,
	0xef, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x06, 0x7e, 0xf4, 0x20, 0xd6, 0x01, 0x00, 0x00,
}
