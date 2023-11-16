// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// IOTServiceClient is the client API for IOTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IOTServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Control(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlResponse, error)
}

type iOTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIOTServiceClient(cc grpc.ClientConnInterface) IOTServiceClient {
	return &iOTServiceClient{cc}
}

func (c *iOTServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/iot.api.v1.IOTService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iOTServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/iot.api.v1.IOTService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iOTServiceClient) Control(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/iot.api.v1.IOTService/Control", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IOTServiceServer is the server API for IOTService service.
// All implementations must embed UnimplementedIOTServiceServer
// for forward compatibility
type IOTServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Control(context.Context, *ControlRequest) (*ControlResponse, error)
	mustEmbedUnimplementedIOTServiceServer()
}

// UnimplementedIOTServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIOTServiceServer struct {
}

func (UnimplementedIOTServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIOTServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIOTServiceServer) Control(context.Context, *ControlRequest) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}
func (UnimplementedIOTServiceServer) mustEmbedUnimplementedIOTServiceServer() {}

// UnsafeIOTServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IOTServiceServer will
// result in compilation errors.
type UnsafeIOTServiceServer interface {
	mustEmbedUnimplementedIOTServiceServer()
}

func RegisterIOTServiceServer(s grpc.ServiceRegistrar, srv IOTServiceServer) {
	s.RegisterService(&_IOTService_serviceDesc, srv)
}

func _IOTService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOTServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iot.api.v1.IOTService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOTServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IOTService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOTServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iot.api.v1.IOTService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOTServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IOTService_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOTServiceServer).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iot.api.v1.IOTService/Control",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOTServiceServer).Control(ctx, req.(*ControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IOTService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iot.api.v1.IOTService",
	HandlerType: (*IOTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _IOTService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _IOTService_Login_Handler,
		},
		{
			MethodName: "Control",
			Handler:    _IOTService_Control_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}