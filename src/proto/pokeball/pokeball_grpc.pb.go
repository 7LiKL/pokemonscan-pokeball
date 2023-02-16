// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/pokeball/pokeball.proto

package pokeball

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatArgs, opts ...grpc.CallOption) (*HeartbeatReply, error)
	GetTask(ctx context.Context, in *GetTaskArgs, opts ...grpc.CallOption) (*GetTaskReply, error)
	ContinueTask(ctx context.Context, in *GetTaskArgs, opts ...grpc.CallOption) (*ContinueTaskReply, error)
	GetRegisteredConfig(ctx context.Context, in *GetRegisteredConfigArgs, opts ...grpc.CallOption) (*GetRegisteredConfigReply, error)
	ReportInfoResult(ctx context.Context, in *ReportInfoArgs, opts ...grpc.CallOption) (*ReportInfoReply, error)
	ReportVulResult(ctx context.Context, in *ReportVulArgs, opts ...grpc.CallOption) (*ReportVulReply, error)
	ReportCompletionStatus(ctx context.Context, in *CompletionStatusArgs, opts ...grpc.CallOption) (*CompletionStatusReply, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Heartbeat(ctx context.Context, in *HeartbeatArgs, opts ...grpc.CallOption) (*HeartbeatReply, error) {
	out := new(HeartbeatReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *GetTaskArgs, opts ...grpc.CallOption) (*GetTaskReply, error) {
	out := new(GetTaskReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ContinueTask(ctx context.Context, in *GetTaskArgs, opts ...grpc.CallOption) (*ContinueTaskReply, error) {
	out := new(ContinueTaskReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/ContinueTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetRegisteredConfig(ctx context.Context, in *GetRegisteredConfigArgs, opts ...grpc.CallOption) (*GetRegisteredConfigReply, error) {
	out := new(GetRegisteredConfigReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/GetRegisteredConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ReportInfoResult(ctx context.Context, in *ReportInfoArgs, opts ...grpc.CallOption) (*ReportInfoReply, error) {
	out := new(ReportInfoReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/ReportInfoResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ReportVulResult(ctx context.Context, in *ReportVulArgs, opts ...grpc.CallOption) (*ReportVulReply, error) {
	out := new(ReportVulReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/ReportVulResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ReportCompletionStatus(ctx context.Context, in *CompletionStatusArgs, opts ...grpc.CallOption) (*CompletionStatusReply, error) {
	out := new(CompletionStatusReply)
	err := c.cc.Invoke(ctx, "/pokemon.proto.pokeball.TaskService/ReportCompletionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	Heartbeat(context.Context, *HeartbeatArgs) (*HeartbeatReply, error)
	GetTask(context.Context, *GetTaskArgs) (*GetTaskReply, error)
	ContinueTask(context.Context, *GetTaskArgs) (*ContinueTaskReply, error)
	GetRegisteredConfig(context.Context, *GetRegisteredConfigArgs) (*GetRegisteredConfigReply, error)
	ReportInfoResult(context.Context, *ReportInfoArgs) (*ReportInfoReply, error)
	ReportVulResult(context.Context, *ReportVulArgs) (*ReportVulReply, error)
	ReportCompletionStatus(context.Context, *CompletionStatusArgs) (*CompletionStatusReply, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) Heartbeat(context.Context, *HeartbeatArgs) (*HeartbeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedTaskServiceServer) GetTask(context.Context, *GetTaskArgs) (*GetTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskServiceServer) ContinueTask(context.Context, *GetTaskArgs) (*ContinueTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContinueTask not implemented")
}
func (UnimplementedTaskServiceServer) GetRegisteredConfig(context.Context, *GetRegisteredConfigArgs) (*GetRegisteredConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredConfig not implemented")
}
func (UnimplementedTaskServiceServer) ReportInfoResult(context.Context, *ReportInfoArgs) (*ReportInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportInfoResult not implemented")
}
func (UnimplementedTaskServiceServer) ReportVulResult(context.Context, *ReportVulArgs) (*ReportVulReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportVulResult not implemented")
}
func (UnimplementedTaskServiceServer) ReportCompletionStatus(context.Context, *CompletionStatusArgs) (*CompletionStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportCompletionStatus not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Heartbeat(ctx, req.(*HeartbeatArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*GetTaskArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ContinueTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ContinueTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/ContinueTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ContinueTask(ctx, req.(*GetTaskArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetRegisteredConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisteredConfigArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetRegisteredConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/GetRegisteredConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetRegisteredConfig(ctx, req.(*GetRegisteredConfigArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ReportInfoResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ReportInfoResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/ReportInfoResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ReportInfoResult(ctx, req.(*ReportInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ReportVulResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportVulArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ReportVulResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/ReportVulResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ReportVulResult(ctx, req.(*ReportVulArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ReportCompletionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionStatusArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ReportCompletionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.proto.pokeball.TaskService/ReportCompletionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ReportCompletionStatus(ctx, req.(*CompletionStatusArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.proto.pokeball.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _TaskService_Heartbeat_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "ContinueTask",
			Handler:    _TaskService_ContinueTask_Handler,
		},
		{
			MethodName: "GetRegisteredConfig",
			Handler:    _TaskService_GetRegisteredConfig_Handler,
		},
		{
			MethodName: "ReportInfoResult",
			Handler:    _TaskService_ReportInfoResult_Handler,
		},
		{
			MethodName: "ReportVulResult",
			Handler:    _TaskService_ReportVulResult_Handler,
		},
		{
			MethodName: "ReportCompletionStatus",
			Handler:    _TaskService_ReportCompletionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pokeball/pokeball.proto",
}
