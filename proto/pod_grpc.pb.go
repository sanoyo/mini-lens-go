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

// PodServiceClient is the client API for PodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodServiceClient interface {
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PodResponse, error)
}

type podServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPodServiceClient(cc grpc.ClientConnInterface) PodServiceClient {
	return &podServiceClient{cc}
}

func (c *podServiceClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PodResponse, error) {
	out := new(PodResponse)
	err := c.cc.Invoke(ctx, "/proto.PodService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodServiceServer is the server API for PodService service.
// All implementations must embed UnimplementedPodServiceServer
// for forward compatibility
type PodServiceServer interface {
	GetStatus(context.Context, *Empty) (*PodResponse, error)
	mustEmbedUnimplementedPodServiceServer()
}

// UnimplementedPodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPodServiceServer struct {
}

func (UnimplementedPodServiceServer) GetStatus(context.Context, *Empty) (*PodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedPodServiceServer) mustEmbedUnimplementedPodServiceServer() {}

// UnsafePodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PodServiceServer will
// result in compilation errors.
type UnsafePodServiceServer interface {
	mustEmbedUnimplementedPodServiceServer()
}

func RegisterPodServiceServer(s grpc.ServiceRegistrar, srv PodServiceServer) {
	s.RegisterService(&PodService_ServiceDesc, srv)
}

func _PodService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PodService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServiceServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PodService_ServiceDesc is the grpc.ServiceDesc for PodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PodService",
	HandlerType: (*PodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _PodService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pod.proto",
}