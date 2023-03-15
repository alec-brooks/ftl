// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ftlv1

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

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	// Ping service for readiness.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Serve a module as part of the mesh.
	Serve(ctx context.Context, in *ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.AgentService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Serve(ctx context.Context, in *ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error) {
	out := new(ServeResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.AgentService/Serve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations should embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	// Ping service for readiness.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Serve a module as part of the mesh.
	Serve(context.Context, *ServeRequest) (*ServeResponse, error)
}

// UnimplementedAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgentServiceServer) Serve(context.Context, *ServeRequest) (*ServeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Serve not implemented")
}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.AgentService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Serve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Serve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.AgentService/Serve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Serve(ctx, req.(*ServeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xyz.block.ftl.v1.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AgentService_Ping_Handler,
		},
		{
			MethodName: "Serve",
			Handler:    _AgentService_Serve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xyz/block/ftl/v1/ftl.proto",
}

// VerbServiceClient is the client API for VerbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerbServiceClient interface {
	// Ping service for readiness.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Call a Verb on the Drive.
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	// List the Verbs available on the service.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type verbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerbServiceClient(cc grpc.ClientConnInterface) VerbServiceClient {
	return &verbServiceClient{cc}
}

func (c *verbServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.VerbService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verbServiceClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.VerbService/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verbServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.VerbService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerbServiceServer is the server API for VerbService service.
// All implementations should embed UnimplementedVerbServiceServer
// for forward compatibility
type VerbServiceServer interface {
	// Ping service for readiness.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Call a Verb on the Drive.
	Call(context.Context, *CallRequest) (*CallResponse, error)
	// List the Verbs available on the service.
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedVerbServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVerbServiceServer struct {
}

func (UnimplementedVerbServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedVerbServiceServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedVerbServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeVerbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerbServiceServer will
// result in compilation errors.
type UnsafeVerbServiceServer interface {
	mustEmbedUnimplementedVerbServiceServer()
}

func RegisterVerbServiceServer(s grpc.ServiceRegistrar, srv VerbServiceServer) {
	s.RegisterService(&VerbService_ServiceDesc, srv)
}

func _VerbService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerbServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.VerbService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerbServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerbService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerbServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.VerbService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerbServiceServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerbService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerbServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.VerbService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerbServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerbService_ServiceDesc is the grpc.ServiceDesc for VerbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xyz.block.ftl.v1.VerbService",
	HandlerType: (*VerbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _VerbService_Ping_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _VerbService_Call_Handler,
		},
		{
			MethodName: "List",
			Handler:    _VerbService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xyz/block/ftl/v1/ftl.proto",
}

// DevelServiceClient is the client API for DevelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevelServiceClient interface {
	// Ping service for readiness.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// FileChange is called when a file is changed.
	//
	// The Drive should hot reload the module if a change to the file warrants it.
	FileChange(ctx context.Context, in *FileChangeRequest, opts ...grpc.CallOption) (*FileChangeResponse, error)
}

type develServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevelServiceClient(cc grpc.ClientConnInterface) DevelServiceClient {
	return &develServiceClient{cc}
}

func (c *develServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.DevelService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *develServiceClient) FileChange(ctx context.Context, in *FileChangeRequest, opts ...grpc.CallOption) (*FileChangeResponse, error) {
	out := new(FileChangeResponse)
	err := c.cc.Invoke(ctx, "/xyz.block.ftl.v1.DevelService/FileChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevelServiceServer is the server API for DevelService service.
// All implementations should embed UnimplementedDevelServiceServer
// for forward compatibility
type DevelServiceServer interface {
	// Ping service for readiness.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// FileChange is called when a file is changed.
	//
	// The Drive should hot reload the module if a change to the file warrants it.
	FileChange(context.Context, *FileChangeRequest) (*FileChangeResponse, error)
}

// UnimplementedDevelServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDevelServiceServer struct {
}

func (UnimplementedDevelServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDevelServiceServer) FileChange(context.Context, *FileChangeRequest) (*FileChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileChange not implemented")
}

// UnsafeDevelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevelServiceServer will
// result in compilation errors.
type UnsafeDevelServiceServer interface {
	mustEmbedUnimplementedDevelServiceServer()
}

func RegisterDevelServiceServer(s grpc.ServiceRegistrar, srv DevelServiceServer) {
	s.RegisterService(&DevelService_ServiceDesc, srv)
}

func _DevelService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.DevelService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelService_FileChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelServiceServer).FileChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xyz.block.ftl.v1.DevelService/FileChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelServiceServer).FileChange(ctx, req.(*FileChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DevelService_ServiceDesc is the grpc.ServiceDesc for DevelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xyz.block.ftl.v1.DevelService",
	HandlerType: (*DevelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DevelService_Ping_Handler,
		},
		{
			MethodName: "FileChange",
			Handler:    _DevelService_FileChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xyz/block/ftl/v1/ftl.proto",
}
