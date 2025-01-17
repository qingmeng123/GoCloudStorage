// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: storage/storage.proto

package storage

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

const (
	Storage_FindStorageByHash_FullMethodName   = "/storage.Storage/FindStorageByHash"
	Storage_CreateStorage_FullMethodName       = "/storage.Storage/CreateStorage"
	Storage_GenerateDownloadURL_FullMethodName = "/storage.Storage/GenerateDownloadURL"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	FindStorageByHash(ctx context.Context, in *FindStorageByHashReq, opts ...grpc.CallOption) (*FindStorageByHashResp, error)
	CreateStorage(ctx context.Context, in *CreateStorageReq, opts ...grpc.CallOption) (*CreateStorageResp, error)
	GenerateDownloadURL(ctx context.Context, in *GenerateDownloadURLReq, opts ...grpc.CallOption) (*GenerateDownloadURLResp, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) FindStorageByHash(ctx context.Context, in *FindStorageByHashReq, opts ...grpc.CallOption) (*FindStorageByHashResp, error) {
	out := new(FindStorageByHashResp)
	err := c.cc.Invoke(ctx, Storage_FindStorageByHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) CreateStorage(ctx context.Context, in *CreateStorageReq, opts ...grpc.CallOption) (*CreateStorageResp, error) {
	out := new(CreateStorageResp)
	err := c.cc.Invoke(ctx, Storage_CreateStorage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GenerateDownloadURL(ctx context.Context, in *GenerateDownloadURLReq, opts ...grpc.CallOption) (*GenerateDownloadURLResp, error) {
	out := new(GenerateDownloadURLResp)
	err := c.cc.Invoke(ctx, Storage_GenerateDownloadURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	FindStorageByHash(context.Context, *FindStorageByHashReq) (*FindStorageByHashResp, error)
	CreateStorage(context.Context, *CreateStorageReq) (*CreateStorageResp, error)
	GenerateDownloadURL(context.Context, *GenerateDownloadURLReq) (*GenerateDownloadURLResp, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) FindStorageByHash(context.Context, *FindStorageByHashReq) (*FindStorageByHashResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStorageByHash not implemented")
}
func (UnimplementedStorageServer) CreateStorage(context.Context, *CreateStorageReq) (*CreateStorageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorage not implemented")
}
func (UnimplementedStorageServer) GenerateDownloadURL(context.Context, *GenerateDownloadURLReq) (*GenerateDownloadURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDownloadURL not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_FindStorageByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageByHashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).FindStorageByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_FindStorageByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).FindStorageByHash(ctx, req.(*FindStorageByHashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_CreateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).CreateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_CreateStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).CreateStorage(ctx, req.(*CreateStorageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GenerateDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDownloadURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GenerateDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GenerateDownloadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GenerateDownloadURL(ctx, req.(*GenerateDownloadURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindStorageByHash",
			Handler:    _Storage_FindStorageByHash_Handler,
		},
		{
			MethodName: "CreateStorage",
			Handler:    _Storage_CreateStorage_Handler,
		},
		{
			MethodName: "GenerateDownloadURL",
			Handler:    _Storage_GenerateDownloadURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage/storage.proto",
}
