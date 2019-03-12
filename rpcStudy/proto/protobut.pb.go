// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/protobut.proto

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	proto/protobut.proto

It has these top-level messages:
	CustomerRequest
	CustomerResponse
	CustomerFilter
*/
package customer

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

// Request message for creating a new customer
type CustomerRequest struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email     string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone     string                     `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Addresses []*CustomerRequest_Address `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *CustomerRequest) Reset()                    { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()               {}
func (*CustomerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CustomerRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CustomerRequest) GetAddresses() []*CustomerRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type CustomerRequest_Address struct {
	Street            string `protobuf:"bytes,1,opt,name=street" json:"street,omitempty"`
	City              string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	State             string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Zip               string `protobuf:"bytes,4,opt,name=zip" json:"zip,omitempty"`
	IsShippingAddress bool   `protobuf:"varint,5,opt,name=isShippingAddress" json:"isShippingAddress,omitempty"`
}

func (m *CustomerRequest_Address) Reset()                    { *m = CustomerRequest_Address{} }
func (m *CustomerRequest_Address) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest_Address) ProtoMessage()               {}
func (*CustomerRequest_Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CustomerRequest_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *CustomerRequest_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CustomerRequest_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CustomerRequest_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *CustomerRequest_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

type CustomerResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *CustomerResponse) Reset()                    { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string            { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()               {}
func (*CustomerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CustomerResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CustomerFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *CustomerFilter) Reset()                    { *m = CustomerFilter{} }
func (m *CustomerFilter) String() string            { return proto.CompactTextString(m) }
func (*CustomerFilter) ProtoMessage()               {}
func (*CustomerFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomerFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomerRequest)(nil), "customer.CustomerRequest")
	proto.RegisterType((*CustomerRequest_Address)(nil), "customer.CustomerRequest.Address")
	proto.RegisterType((*CustomerResponse)(nil), "customer.CustomerResponse")
	proto.RegisterType((*CustomerFilter)(nil), "customer.CustomerFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Customer service

type CustomerClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error)
	// Create a new Customer - A simple RPC
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Customer_serviceDesc.Streams[0], c.cc, "/customer.Customer/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_GetCustomersClient interface {
	Recv() (*CustomerRequest, error)
	grpc.ClientStream
}

type customerGetCustomersClient struct {
	grpc.ClientStream
}

func (x *customerGetCustomersClient) Recv() (*CustomerRequest, error) {
	m := new(CustomerRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := grpc.Invoke(ctx, "/customer.Customer/CreateCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Customer service

type CustomerServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(*CustomerFilter, Customer_GetCustomersServer) error
	// Create a new Customer - A simple RPC
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).GetCustomers(m, &customerGetCustomersServer{stream})
}

type Customer_GetCustomersServer interface {
	Send(*CustomerRequest) error
	grpc.ServerStream
}

type customerGetCustomersServer struct {
	grpc.ServerStream
}

func (x *customerGetCustomersServer) Send(m *CustomerRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _Customer_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/protobut.proto",
}

func init() { proto.RegisterFile("proto/protobut.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xef, 0x4a, 0xc3, 0x30,
	0x10, 0xc0, 0x97, 0x6e, 0xdd, 0x9f, 0x53, 0xe6, 0x0c, 0x43, 0x62, 0x3f, 0xd5, 0x7e, 0x2a, 0x22,
	0x55, 0xe6, 0x57, 0x41, 0x64, 0xe0, 0xf0, 0x6b, 0x7d, 0x82, 0xae, 0x3d, 0x5c, 0x70, 0x6b, 0x6b,
	0x92, 0x22, 0xf3, 0x15, 0x7c, 0x07, 0x9f, 0xc1, 0x47, 0x94, 0xa4, 0x89, 0x03, 0xe7, 0xbe, 0x94,
	0xfb, 0x5d, 0xee, 0x2e, 0xbf, 0x1c, 0x85, 0x69, 0x2d, 0x2a, 0x55, 0x5d, 0x9b, 0xef, 0xb2, 0x51,
	0x89, 0x09, 0xe8, 0x30, 0x6f, 0xa4, 0xaa, 0x36, 0x28, 0xa2, 0x6f, 0x0f, 0x4e, 0xe6, 0x16, 0x52,
	0x7c, 0x6b, 0x50, 0x2a, 0x3a, 0x06, 0x8f, 0x17, 0x8c, 0x84, 0x24, 0xf6, 0x53, 0x8f, 0x17, 0x94,
	0x42, 0xaf, 0xcc, 0x36, 0xc8, 0xbc, 0x90, 0xc4, 0xa3, 0xd4, 0xc4, 0x74, 0x0a, 0x3e, 0x6e, 0x32,
	0xbe, 0x66, 0x5d, 0x93, 0x6c, 0x41, 0x67, 0xeb, 0x55, 0x55, 0x22, 0xeb, 0xb5, 0x59, 0x03, 0xf4,
	0x1e, 0x46, 0x59, 0x51, 0x08, 0x94, 0x12, 0x25, 0xf3, 0xc3, 0x6e, 0x7c, 0x34, 0xbb, 0x48, 0x9c,
	0x41, 0xf2, 0xe7, 0xf6, 0xe4, 0xa1, 0x2d, 0x4d, 0x77, 0x3d, 0xc1, 0x27, 0x81, 0x81, 0x4d, 0xd3,
	0x33, 0xe8, 0x4b, 0x25, 0x10, 0x95, 0x11, 0x1c, 0xa5, 0x96, 0xb4, 0x64, 0xce, 0xd5, 0xd6, 0x49,
	0xea, 0x58, 0xeb, 0x48, 0x95, 0x29, 0x74, 0x92, 0x06, 0xe8, 0x04, 0xba, 0x1f, 0xbc, 0xb6, 0x8a,
	0x3a, 0xa4, 0x57, 0x70, 0xca, 0xe5, 0xf3, 0x8a, 0xd7, 0x35, 0x2f, 0x5f, 0xec, 0x45, 0xcc, 0x0f,
	0x49, 0x3c, 0x4c, 0xf7, 0x0f, 0xa2, 0x3b, 0x98, 0xec, 0x9c, 0x65, 0x5d, 0x95, 0x12, 0xf7, 0x56,
	0xc6, 0x60, 0x20, 0x9b, 0x3c, 0xd7, 0x73, 0x3c, 0x33, 0xc7, 0x61, 0x74, 0x09, 0x63, 0xd7, 0xfd,
	0xc8, 0xd7, 0x0a, 0x85, 0xae, 0x7d, 0xc5, 0xed, 0x7b, 0x25, 0x0a, 0xfb, 0x24, 0x87, 0xb3, 0x2f,
	0x02, 0x43, 0x57, 0x4c, 0x17, 0x70, 0xbc, 0x40, 0xe5, 0x50, 0x52, 0xb6, 0xbf, 0xc2, 0x76, 0x60,
	0x70, 0x7e, 0x70, 0xb9, 0x51, 0xe7, 0x86, 0xd0, 0x27, 0x18, 0xcf, 0x05, 0x66, 0x0a, 0x7f, 0x47,
	0x1f, 0x6e, 0x08, 0x82, 0xff, 0x8e, 0xda, 0x47, 0x47, 0x9d, 0x65, 0xdf, 0xfc, 0x4e, 0xb7, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xfa, 0x64, 0x6d, 0x66, 0x02, 0x00, 0x00,
}