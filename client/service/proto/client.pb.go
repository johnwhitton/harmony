// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package client

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// FetchAccountStateRequest is the request to fetch an account's balance and nonce.
type FetchAccountStateRequest struct {
	// The account address
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAccountStateRequest) Reset()         { *m = FetchAccountStateRequest{} }
func (m *FetchAccountStateRequest) String() string { return proto.CompactTextString(m) }
func (*FetchAccountStateRequest) ProtoMessage()    {}
func (*FetchAccountStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *FetchAccountStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAccountStateRequest.Unmarshal(m, b)
}
func (m *FetchAccountStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAccountStateRequest.Marshal(b, m, deterministic)
}
func (m *FetchAccountStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAccountStateRequest.Merge(m, src)
}
func (m *FetchAccountStateRequest) XXX_Size() int {
	return xxx_messageInfo_FetchAccountStateRequest.Size(m)
}
func (m *FetchAccountStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAccountStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAccountStateRequest proto.InternalMessageInfo

func (m *FetchAccountStateRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// FetchAccountStateResponse is the response of FetchAccountStateRequest.
type FetchAccountStateResponse struct {
	// The balance of the account (big.Int)
	Balance []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	// The nonce of the account
	Nonce                uint64   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAccountStateResponse) Reset()         { *m = FetchAccountStateResponse{} }
func (m *FetchAccountStateResponse) String() string { return proto.CompactTextString(m) }
func (*FetchAccountStateResponse) ProtoMessage()    {}
func (*FetchAccountStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func (m *FetchAccountStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAccountStateResponse.Unmarshal(m, b)
}
func (m *FetchAccountStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAccountStateResponse.Marshal(b, m, deterministic)
}
func (m *FetchAccountStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAccountStateResponse.Merge(m, src)
}
func (m *FetchAccountStateResponse) XXX_Size() int {
	return xxx_messageInfo_FetchAccountStateResponse.Size(m)
}
func (m *FetchAccountStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAccountStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAccountStateResponse proto.InternalMessageInfo

func (m *FetchAccountStateResponse) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *FetchAccountStateResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchAccountStateRequest)(nil), "client.FetchAccountStateRequest")
	proto.RegisterType((*FetchAccountStateResponse)(nil), "client.FetchAccountStateResponse")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x4c, 0xb8, 0x24,
	0xdc, 0x52, 0x4b, 0x92, 0x33, 0x1c, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x82, 0x4b, 0x12, 0x4b,
	0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a,
	0x52, 0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x25, 0x6f, 0x2e, 0x49,
	0x2c, 0xba, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0xda, 0x92, 0x12, 0x73, 0x12, 0xf3, 0x92,
	0x53, 0x61, 0xda, 0xa0, 0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xbc, 0x7c, 0x90, 0x38, 0x93, 0x02, 0xa3,
	0x06, 0x4b, 0x10, 0x84, 0x63, 0x94, 0xcd, 0xc5, 0xeb, 0x0c, 0x76, 0x4c, 0x70, 0x6a, 0x51, 0x59,
	0x66, 0x72, 0xaa, 0x50, 0x14, 0x97, 0x20, 0x86, 0xe9, 0x42, 0x0a, 0x7a, 0x50, 0xf7, 0xe3, 0x72,
	0xae, 0x94, 0x22, 0x1e, 0x15, 0x10, 0xa7, 0x29, 0x31, 0x24, 0xb1, 0x81, 0xbd, 0x6f, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0x07, 0xa9, 0xb6, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	FetchAccountState(ctx context.Context, in *FetchAccountStateRequest, opts ...grpc.CallOption) (*FetchAccountStateResponse, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) FetchAccountState(ctx context.Context, in *FetchAccountStateRequest, opts ...grpc.CallOption) (*FetchAccountStateResponse, error) {
	out := new(FetchAccountStateResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/FetchAccountState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	FetchAccountState(context.Context, *FetchAccountStateRequest) (*FetchAccountStateResponse, error)
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_FetchAccountState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAccountStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).FetchAccountState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/FetchAccountState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).FetchAccountState(ctx, req.(*FetchAccountStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAccountState",
			Handler:    _ClientService_FetchAccountState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}