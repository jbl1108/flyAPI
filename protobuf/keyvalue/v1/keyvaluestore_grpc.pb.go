// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: protobuf/keyvalue/v1/keyvaluestore.proto

package flyapi

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

// StoreKeyValueClient is the client API for StoreKeyValue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreKeyValueClient interface {
	//StoreKeyValue
	StoreKeyValue(ctx context.Context, in *KeyValues, opts ...grpc.CallOption) (*StatusReply, error)
}

type storeKeyValueClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreKeyValueClient(cc grpc.ClientConnInterface) StoreKeyValueClient {
	return &storeKeyValueClient{cc}
}

func (c *storeKeyValueClient) StoreKeyValue(ctx context.Context, in *KeyValues, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/keyvaluestore.StoreKeyValue/StoreKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreKeyValueServer is the server API for StoreKeyValue service.
// All implementations must embed UnimplementedStoreKeyValueServer
// for forward compatibility
type StoreKeyValueServer interface {
	//StoreKeyValue
	StoreKeyValue(context.Context, *KeyValues) (*StatusReply, error)
	mustEmbedUnimplementedStoreKeyValueServer()
}

// UnimplementedStoreKeyValueServer must be embedded to have forward compatible implementations.
type UnimplementedStoreKeyValueServer struct {
}

func (UnimplementedStoreKeyValueServer) StoreKeyValue(context.Context, *KeyValues) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreKeyValue not implemented")
}
func (UnimplementedStoreKeyValueServer) mustEmbedUnimplementedStoreKeyValueServer() {}

// UnsafeStoreKeyValueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreKeyValueServer will
// result in compilation errors.
type UnsafeStoreKeyValueServer interface {
	mustEmbedUnimplementedStoreKeyValueServer()
}

func RegisterStoreKeyValueServer(s grpc.ServiceRegistrar, srv StoreKeyValueServer) {
	s.RegisterService(&StoreKeyValue_ServiceDesc, srv)
}

func _StoreKeyValue_StoreKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValues)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreKeyValueServer).StoreKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvaluestore.StoreKeyValue/StoreKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreKeyValueServer).StoreKeyValue(ctx, req.(*KeyValues))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreKeyValue_ServiceDesc is the grpc.ServiceDesc for StoreKeyValue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreKeyValue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyvaluestore.StoreKeyValue",
	HandlerType: (*StoreKeyValueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreKeyValue",
			Handler:    _StoreKeyValue_StoreKeyValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/keyvalue/v1/keyvaluestore.proto",
}
