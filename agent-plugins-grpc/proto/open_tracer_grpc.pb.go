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

// KernelOpenTracerClient is the client API for KernelOpenTracer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KernelOpenTracerClient interface {
	GetTraceInfo(ctx context.Context, in *TraceInfoRequest, opts ...grpc.CallOption) (KernelOpenTracer_GetTraceInfoClient, error)
	GetBillOfMaterials(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BillOfMaterials, error)
	ClearBillOfMaterials(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type kernelOpenTracerClient struct {
	cc grpc.ClientConnInterface
}

func NewKernelOpenTracerClient(cc grpc.ClientConnInterface) KernelOpenTracerClient {
	return &kernelOpenTracerClient{cc}
}

func (c *kernelOpenTracerClient) GetTraceInfo(ctx context.Context, in *TraceInfoRequest, opts ...grpc.CallOption) (KernelOpenTracer_GetTraceInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &KernelOpenTracer_ServiceDesc.Streams[0], "/open_tracer.KernelOpenTracer/GetTraceInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &kernelOpenTracerGetTraceInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KernelOpenTracer_GetTraceInfoClient interface {
	Recv() (*TraceInfo, error)
	grpc.ClientStream
}

type kernelOpenTracerGetTraceInfoClient struct {
	grpc.ClientStream
}

func (x *kernelOpenTracerGetTraceInfoClient) Recv() (*TraceInfo, error) {
	m := new(TraceInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kernelOpenTracerClient) GetBillOfMaterials(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BillOfMaterials, error) {
	out := new(BillOfMaterials)
	err := c.cc.Invoke(ctx, "/open_tracer.KernelOpenTracer/GetBillOfMaterials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelOpenTracerClient) ClearBillOfMaterials(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/open_tracer.KernelOpenTracer/ClearBillOfMaterials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KernelOpenTracerServer is the server API for KernelOpenTracer service.
// All implementations must embed UnimplementedKernelOpenTracerServer
// for forward compatibility
type KernelOpenTracerServer interface {
	GetTraceInfo(*TraceInfoRequest, KernelOpenTracer_GetTraceInfoServer) error
	GetBillOfMaterials(context.Context, *Empty) (*BillOfMaterials, error)
	ClearBillOfMaterials(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedKernelOpenTracerServer()
}

// UnimplementedKernelOpenTracerServer must be embedded to have forward compatible implementations.
type UnimplementedKernelOpenTracerServer struct {
}

func (UnimplementedKernelOpenTracerServer) GetTraceInfo(*TraceInfoRequest, KernelOpenTracer_GetTraceInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTraceInfo not implemented")
}
func (UnimplementedKernelOpenTracerServer) GetBillOfMaterials(context.Context, *Empty) (*BillOfMaterials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillOfMaterials not implemented")
}
func (UnimplementedKernelOpenTracerServer) ClearBillOfMaterials(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearBillOfMaterials not implemented")
}
func (UnimplementedKernelOpenTracerServer) mustEmbedUnimplementedKernelOpenTracerServer() {}

// UnsafeKernelOpenTracerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KernelOpenTracerServer will
// result in compilation errors.
type UnsafeKernelOpenTracerServer interface {
	mustEmbedUnimplementedKernelOpenTracerServer()
}

func RegisterKernelOpenTracerServer(s grpc.ServiceRegistrar, srv KernelOpenTracerServer) {
	s.RegisterService(&KernelOpenTracer_ServiceDesc, srv)
}

func _KernelOpenTracer_GetTraceInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraceInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KernelOpenTracerServer).GetTraceInfo(m, &kernelOpenTracerGetTraceInfoServer{stream})
}

type KernelOpenTracer_GetTraceInfoServer interface {
	Send(*TraceInfo) error
	grpc.ServerStream
}

type kernelOpenTracerGetTraceInfoServer struct {
	grpc.ServerStream
}

func (x *kernelOpenTracerGetTraceInfoServer) Send(m *TraceInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _KernelOpenTracer_GetBillOfMaterials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelOpenTracerServer).GetBillOfMaterials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_tracer.KernelOpenTracer/GetBillOfMaterials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelOpenTracerServer).GetBillOfMaterials(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KernelOpenTracer_ClearBillOfMaterials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelOpenTracerServer).ClearBillOfMaterials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_tracer.KernelOpenTracer/ClearBillOfMaterials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelOpenTracerServer).ClearBillOfMaterials(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// KernelOpenTracer_ServiceDesc is the grpc.ServiceDesc for KernelOpenTracer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KernelOpenTracer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "open_tracer.KernelOpenTracer",
	HandlerType: (*KernelOpenTracerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillOfMaterials",
			Handler:    _KernelOpenTracer_GetBillOfMaterials_Handler,
		},
		{
			MethodName: "ClearBillOfMaterials",
			Handler:    _KernelOpenTracer_ClearBillOfMaterials_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTraceInfo",
			Handler:       _KernelOpenTracer_GetTraceInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "open_tracer.proto",
}
