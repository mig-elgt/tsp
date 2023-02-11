// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: table/table.proto

package table

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

// TableServiceClient is the client API for TableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TableServiceClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
}

type tableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTableServiceClient(cc grpc.ClientConnInterface) TableServiceClient {
	return &tableServiceClient{cc}
}

func (c *tableServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/TableService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TableServiceServer is the server API for TableService service.
// All implementations must embed UnimplementedTableServiceServer
// for forward compatibility
type TableServiceServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	mustEmbedUnimplementedTableServiceServer()
}

// UnimplementedTableServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTableServiceServer struct {
}

func (UnimplementedTableServiceServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedTableServiceServer) mustEmbedUnimplementedTableServiceServer() {}

// UnsafeTableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TableServiceServer will
// result in compilation errors.
type UnsafeTableServiceServer interface {
	mustEmbedUnimplementedTableServiceServer()
}

func RegisterTableServiceServer(s grpc.ServiceRegistrar, srv TableServiceServer) {
	s.RegisterService(&TableService_ServiceDesc, srv)
}

func _TableService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TableService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TableService_ServiceDesc is the grpc.ServiceDesc for TableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TableService",
	HandlerType: (*TableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _TableService_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "table/table.proto",
}
