// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bank/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("bank/v1beta/querier.proto", fileDescriptor_ee07edbbdc40ed12) }

var fileDescriptor_ee07edbbdc40ed12 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4a, 0xcc, 0xcb,
	0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcd, 0x48, 0xcd, 0xcc, 0x4d, 0x49, 0xcc, 0xc9, 0xd1,
	0x03, 0xa9, 0xd1, 0x83, 0xa8, 0x31, 0x94, 0xd2, 0x4a, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x4f,
	0x4a, 0x2c, 0x4e, 0x05, 0xeb, 0xa8, 0x84, 0x6a, 0x37, 0xd4, 0x2f, 0x48, 0x4c, 0xcf, 0xcc, 0x4b,
	0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0x18, 0x61, 0xc4, 0xce, 0xc5, 0x1a, 0x08, 0x52, 0xe1, 0xe4, 0x7e,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0x89, 0x25, 0x99, 0xc9, 0x79, 0xa9, 0x25, 0xe5, 0xf9,
	0x45, 0xd9, 0xfa, 0x30, 0xdb, 0xf5, 0x2b, 0xf4, 0xc1, 0x6e, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x4e,
	0x62, 0x03, 0x1b, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x19, 0x57, 0x7f, 0x8b, 0xb8, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.bank.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "bank/v1beta/querier.proto",
}
