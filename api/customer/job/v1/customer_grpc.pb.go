// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: v1/customer.proto

package v1

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

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeReply, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeReply, error) {
	out := new(ConsumeReply)
	err := c.cc.Invoke(ctx, "/sentinel.job.v1.Customer/Consume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility
type CustomerServer interface {
	Consume(context.Context, *ConsumeRequest) (*ConsumeReply, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (UnimplementedCustomerServer) Consume(context.Context, *ConsumeRequest) (*ConsumeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.job.v1.Customer/Consume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.job.v1.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Consume",
			Handler:    _Customer_Consume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/customer.proto",
}
