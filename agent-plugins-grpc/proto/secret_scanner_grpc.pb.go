// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SecretScannerClient is the client API for SecretScanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretScannerClient interface {
	FindSecretInfo(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResult, error)
}

type secretScannerClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretScannerClient(cc grpc.ClientConnInterface) SecretScannerClient {
	return &secretScannerClient{cc}
}

func (c *secretScannerClient) FindSecretInfo(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResult, error) {
	out := new(FindResult)
	err := c.cc.Invoke(ctx, "/secret_scanner.SecretScanner/FindSecretInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretScannerServer is the server API for SecretScanner service.
// All implementations must embed UnimplementedSecretScannerServer
// for forward compatibility
type SecretScannerServer interface {
	FindSecretInfo(context.Context, *FindRequest) (*FindResult, error)
	mustEmbedUnimplementedSecretScannerServer()
}

// UnimplementedSecretScannerServer must be embedded to have forward compatible implementations.
type UnimplementedSecretScannerServer struct {
}

func (UnimplementedSecretScannerServer) FindSecretInfo(context.Context, *FindRequest) (*FindResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSecretInfo not implemented")
}
func (UnimplementedSecretScannerServer) mustEmbedUnimplementedSecretScannerServer() {}

// UnsafeSecretScannerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretScannerServer will
// result in compilation errors.
type UnsafeSecretScannerServer interface {
	mustEmbedUnimplementedSecretScannerServer()
}

func RegisterSecretScannerServer(s grpc.ServiceRegistrar, srv SecretScannerServer) {
	s.RegisterService(&SecretScanner_ServiceDesc, srv)
}

func _SecretScanner_FindSecretInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretScannerServer).FindSecretInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secret_scanner.SecretScanner/FindSecretInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretScannerServer).FindSecretInfo(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretScanner_ServiceDesc is the grpc.ServiceDesc for SecretScanner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretScanner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "secret_scanner.SecretScanner",
	HandlerType: (*SecretScannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSecretInfo",
			Handler:    _SecretScanner_FindSecretInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret_scanner.proto",
}
