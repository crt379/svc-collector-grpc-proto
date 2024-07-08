// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: register/register.proto

package register

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Register_Register_FullMethodName    = "/service.collector.register.Register/Register"
	Register_GetRegister_FullMethodName = "/service.collector.register.Register/GetRegister"
	Register_Unregister_FullMethodName  = "/service.collector.register.Register/Unregister"
)

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	GetRegister(ctx context.Context, in *GetRegisterRequest, opts ...grpc.CallOption) (*GetRegisterReply, error)
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterReply, error)
}

type registerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterClient(cc grpc.ClientConnInterface) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, Register_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) GetRegister(ctx context.Context, in *GetRegisterRequest, opts ...grpc.CallOption) (*GetRegisterReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRegisterReply)
	err := c.cc.Invoke(ctx, Register_GetRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnregisterReply)
	err := c.cc.Invoke(ctx, Register_Unregister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
// All implementations must embed UnimplementedRegisterServer
// for forward compatibility
type RegisterServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	GetRegister(context.Context, *GetRegisterRequest) (*GetRegisterReply, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterReply, error)
	mustEmbedUnimplementedRegisterServer()
}

// UnimplementedRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (UnimplementedRegisterServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegisterServer) GetRegister(context.Context, *GetRegisterRequest) (*GetRegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegister not implemented")
}
func (UnimplementedRegisterServer) Unregister(context.Context, *UnregisterRequest) (*UnregisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (UnimplementedRegisterServer) mustEmbedUnimplementedRegisterServer() {}

// UnsafeRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServer will
// result in compilation errors.
type UnsafeRegisterServer interface {
	mustEmbedUnimplementedRegisterServer()
}

func RegisterRegisterServer(s grpc.ServiceRegistrar, srv RegisterServer) {
	s.RegisterService(&Register_ServiceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_GetRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).GetRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_GetRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).GetRegister(ctx, req.(*GetRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_Unregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Register_ServiceDesc is the grpc.ServiceDesc for Register service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Register_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.collector.register.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
		{
			MethodName: "GetRegister",
			Handler:    _Register_GetRegister_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _Register_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register/register.proto",
}