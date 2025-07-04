// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: command/v1/service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CommandService_CreateCommand_FullMethodName                  = "/command.v1.CommandService/CreateCommand"
	CommandService_GetCommand_FullMethodName                     = "/command.v1.CommandService/GetCommand"
	CommandService_CancelCurrentProcessingCommand_FullMethodName = "/command.v1.CommandService/CancelCurrentProcessingCommand"
)

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	// CreateCommand creates a new command.
	CreateCommand(ctx context.Context, in *CreateCommandRequest, opts ...grpc.CallOption) (*CreateCommandResponse, error)
	// GetCommand gets a command by its ID.
	GetCommand(ctx context.Context, in *GetCommandRequest, opts ...grpc.CallOption) (*GetCommandResponse, error)
	// CancelCurrentProcessingCommand cancels the current processing command.
	CancelCurrentProcessingCommand(ctx context.Context, in *CancelCurrentProcessingCommandRequest, opts ...grpc.CallOption) (*CancelCurrentProcessingCommandResponse, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) CreateCommand(ctx context.Context, in *CreateCommandRequest, opts ...grpc.CallOption) (*CreateCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommandResponse)
	err := c.cc.Invoke(ctx, CommandService_CreateCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) GetCommand(ctx context.Context, in *GetCommandRequest, opts ...grpc.CallOption) (*GetCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommandResponse)
	err := c.cc.Invoke(ctx, CommandService_GetCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) CancelCurrentProcessingCommand(ctx context.Context, in *CancelCurrentProcessingCommandRequest, opts ...grpc.CallOption) (*CancelCurrentProcessingCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelCurrentProcessingCommandResponse)
	err := c.cc.Invoke(ctx, CommandService_CancelCurrentProcessingCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility.
type CommandServiceServer interface {
	// CreateCommand creates a new command.
	CreateCommand(context.Context, *CreateCommandRequest) (*CreateCommandResponse, error)
	// GetCommand gets a command by its ID.
	GetCommand(context.Context, *GetCommandRequest) (*GetCommandResponse, error)
	// CancelCurrentProcessingCommand cancels the current processing command.
	CancelCurrentProcessingCommand(context.Context, *CancelCurrentProcessingCommandRequest) (*CancelCurrentProcessingCommandResponse, error)
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommandServiceServer struct{}

func (UnimplementedCommandServiceServer) CreateCommand(context.Context, *CreateCommandRequest) (*CreateCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommand not implemented")
}
func (UnimplementedCommandServiceServer) GetCommand(context.Context, *GetCommandRequest) (*GetCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommand not implemented")
}
func (UnimplementedCommandServiceServer) CancelCurrentProcessingCommand(context.Context, *CancelCurrentProcessingCommandRequest) (*CancelCurrentProcessingCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelCurrentProcessingCommand not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}
func (UnimplementedCommandServiceServer) testEmbeddedByValue()                        {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommandServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommandService_ServiceDesc, srv)
}

func _CommandService_CreateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).CreateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_CreateCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).CreateCommand(ctx, req.(*CreateCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_GetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).GetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_GetCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).GetCommand(ctx, req.(*GetCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_CancelCurrentProcessingCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelCurrentProcessingCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).CancelCurrentProcessingCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_CancelCurrentProcessingCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).CancelCurrentProcessingCommand(ctx, req.(*CancelCurrentProcessingCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandService_ServiceDesc is the grpc.ServiceDesc for CommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "command.v1.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommand",
			Handler:    _CommandService_CreateCommand_Handler,
		},
		{
			MethodName: "GetCommand",
			Handler:    _CommandService_GetCommand_Handler,
		},
		{
			MethodName: "CancelCurrentProcessingCommand",
			Handler:    _CommandService_CancelCurrentProcessingCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command/v1/service.proto",
}
