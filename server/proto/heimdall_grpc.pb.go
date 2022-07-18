// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: server/proto/heimdall.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HeimdallClient is the client API for Heimdall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeimdallClient interface {
	GetSpan(ctx context.Context, in *GetSpanRequest, opts ...grpc.CallOption) (*GetSpanResponse, error)
	GetEventRecords(ctx context.Context, in *GetEventRecordsRequest, opts ...grpc.CallOption) (Heimdall_GetEventRecordsClient, error)
}

type heimdallClient struct {
	cc grpc.ClientConnInterface
}

func NewHeimdallClient(cc grpc.ClientConnInterface) HeimdallClient {
	return &heimdallClient{cc}
}

func (c *heimdallClient) GetSpan(ctx context.Context, in *GetSpanRequest, opts ...grpc.CallOption) (*GetSpanResponse, error) {
	out := new(GetSpanResponse)
	err := c.cc.Invoke(ctx, "/proto.Heimdall/GetSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) GetEventRecords(ctx context.Context, in *GetEventRecordsRequest, opts ...grpc.CallOption) (Heimdall_GetEventRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Heimdall_ServiceDesc.Streams[0], "/proto.Heimdall/GetEventRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &heimdallGetEventRecordsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Heimdall_GetEventRecordsClient interface {
	Recv() (*GetEventRecordsResponse, error)
	grpc.ClientStream
}

type heimdallGetEventRecordsClient struct {
	grpc.ClientStream
}

func (x *heimdallGetEventRecordsClient) Recv() (*GetEventRecordsResponse, error) {
	m := new(GetEventRecordsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeimdallServer is the server API for Heimdall service.
// All implementations must embed UnimplementedHeimdallServer
// for forward compatibility
type HeimdallServer interface {
	GetSpan(context.Context, *GetSpanRequest) (*GetSpanResponse, error)
	GetEventRecords(*GetEventRecordsRequest, Heimdall_GetEventRecordsServer) error
	mustEmbedUnimplementedHeimdallServer()
}

// UnimplementedHeimdallServer must be embedded to have forward compatible implementations.
type UnimplementedHeimdallServer struct {
}

func (UnimplementedHeimdallServer) GetSpan(context.Context, *GetSpanRequest) (*GetSpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpan not implemented")
}
func (UnimplementedHeimdallServer) GetEventRecords(*GetEventRecordsRequest, Heimdall_GetEventRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventRecords not implemented")
}
func (UnimplementedHeimdallServer) mustEmbedUnimplementedHeimdallServer() {}

// UnsafeHeimdallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeimdallServer will
// result in compilation errors.
type UnsafeHeimdallServer interface {
	mustEmbedUnimplementedHeimdallServer()
}

func RegisterHeimdallServer(s grpc.ServiceRegistrar, srv HeimdallServer) {
	s.RegisterService(&Heimdall_ServiceDesc, srv)
}

func _Heimdall_GetSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).GetSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Heimdall/GetSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).GetSpan(ctx, req.(*GetSpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_GetEventRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventRecordsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HeimdallServer).GetEventRecords(m, &heimdallGetEventRecordsServer{stream})
}

type Heimdall_GetEventRecordsServer interface {
	Send(*GetEventRecordsResponse) error
	grpc.ServerStream
}

type heimdallGetEventRecordsServer struct {
	grpc.ServerStream
}

func (x *heimdallGetEventRecordsServer) Send(m *GetEventRecordsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Heimdall_ServiceDesc is the grpc.ServiceDesc for Heimdall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Heimdall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Heimdall",
	HandlerType: (*HeimdallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSpan",
			Handler:    _Heimdall_GetSpan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEventRecords",
			Handler:       _Heimdall_GetEventRecords_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/proto/heimdall.proto",
}
