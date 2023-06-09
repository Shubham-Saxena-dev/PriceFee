// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: fees.proto

package stubs

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

// FeeServiceClient is the client API for FeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeeServiceClient interface {
	GetFee(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFeeResponse, error)
	AddFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error)
	RemoveFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error)
}

type feeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeeServiceClient(cc grpc.ClientConnInterface) FeeServiceClient {
	return &feeServiceClient{cc}
}

func (c *feeServiceClient) GetFee(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFeeResponse, error) {
	out := new(GetFeeResponse)
	err := c.cc.Invoke(ctx, "/repositories.FeeService/GetFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) AddFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/repositories.FeeService/AddFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) RemoveFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/repositories.FeeService/RemoveFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeeServiceServer is the server API for FeeService service.
// All implementations must embed UnimplementedFeeServiceServer
// for forward compatibility
type FeeServiceServer interface {
	GetFee(context.Context, *GetFeeRequest) (*GetFeeResponse, error)
	AddFee(context.Context, *FeeRequest) (*FeeResponse, error)
	RemoveFee(context.Context, *FeeRequest) (*FeeResponse, error)
	mustEmbedUnimplementedFeeServiceServer()
}

// UnimplementedFeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeeServiceServer struct {
}

func (UnimplementedFeeServiceServer) GetFee(context.Context, *GetFeeRequest) (*GetFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}
func (UnimplementedFeeServiceServer) AddFee(context.Context, *FeeRequest) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFee not implemented")
}
func (UnimplementedFeeServiceServer) RemoveFee(context.Context, *FeeRequest) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFee not implemented")
}
func (UnimplementedFeeServiceServer) mustEmbedUnimplementedFeeServiceServer() {}

// UnsafeFeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeeServiceServer will
// result in compilation errors.
type UnsafeFeeServiceServer interface {
	mustEmbedUnimplementedFeeServiceServer()
}

func RegisterFeeServiceServer(s grpc.ServiceRegistrar, srv FeeServiceServer) {
	s.RegisterService(&FeeService_ServiceDesc, srv)
}

func _FeeService_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repositories.FeeService/GetFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).GetFee(ctx, req.(*GetFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_AddFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).AddFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repositories.FeeService/AddFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).AddFee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_RemoveFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).RemoveFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repositories.FeeService/RemoveFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).RemoveFee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeeService_ServiceDesc is the grpc.ServiceDesc for FeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "repositories.FeeService",
	HandlerType: (*FeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFee",
			Handler:    _FeeService_GetFee_Handler,
		},
		{
			MethodName: "AddFee",
			Handler:    _FeeService_AddFee_Handler,
		},
		{
			MethodName: "RemoveFee",
			Handler:    _FeeService_RemoveFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fees.proto",
}
