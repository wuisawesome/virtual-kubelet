// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: cmd/virtual-kubelet/internal/provider/grpc/pod_provider.proto

package grpc

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
	PodProvider_ConfigureNode_FullMethodName = "/PodProvider/ConfigureNode"
)

// PodProviderClient is the client API for PodProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodProviderClient interface {
	ConfigureNode(ctx context.Context, in *ConfigureNodeRequest, opts ...grpc.CallOption) (*ConfigureNodeReply, error)
}

type podProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewPodProviderClient(cc grpc.ClientConnInterface) PodProviderClient {
	return &podProviderClient{cc}
}

func (c *podProviderClient) ConfigureNode(ctx context.Context, in *ConfigureNodeRequest, opts ...grpc.CallOption) (*ConfigureNodeReply, error) {
	out := new(ConfigureNodeReply)
	err := c.cc.Invoke(ctx, PodProvider_ConfigureNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodProviderServer is the server API for PodProvider service.
// All implementations must embed UnimplementedPodProviderServer
// for forward compatibility
type PodProviderServer interface {
	ConfigureNode(context.Context, *ConfigureNodeRequest) (*ConfigureNodeReply, error)
	mustEmbedUnimplementedPodProviderServer()
}

// UnimplementedPodProviderServer must be embedded to have forward compatible implementations.
type UnimplementedPodProviderServer struct {
}

func (UnimplementedPodProviderServer) ConfigureNode(context.Context, *ConfigureNodeRequest) (*ConfigureNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureNode not implemented")
}
func (UnimplementedPodProviderServer) mustEmbedUnimplementedPodProviderServer() {}

// UnsafePodProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PodProviderServer will
// result in compilation errors.
type UnsafePodProviderServer interface {
	mustEmbedUnimplementedPodProviderServer()
}

func RegisterPodProviderServer(s grpc.ServiceRegistrar, srv PodProviderServer) {
	s.RegisterService(&PodProvider_ServiceDesc, srv)
}

func _PodProvider_ConfigureNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodProviderServer).ConfigureNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PodProvider_ConfigureNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodProviderServer).ConfigureNode(ctx, req.(*ConfigureNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PodProvider_ServiceDesc is the grpc.ServiceDesc for PodProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PodProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PodProvider",
	HandlerType: (*PodProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureNode",
			Handler:    _PodProvider_ConfigureNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/virtual-kubelet/internal/provider/grpc/pod_provider.proto",
}
