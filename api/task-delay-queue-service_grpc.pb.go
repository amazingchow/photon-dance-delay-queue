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

// TaskDelayQueueServiceClient is the client API for TaskDelayQueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskDelayQueueServiceClient interface {
	// for task producers
	PushTask(ctx context.Context, in *PushTaskRequest, opts ...grpc.CallOption) (*PushTaskResponse, error)
	FinishTask(ctx context.Context, in *FinishTaskRequest, opts ...grpc.CallOption) (*FinishTaskResponse, error)
	CheckTask(ctx context.Context, in *CheckTaskRequest, opts ...grpc.CallOption) (*CheckTaskResponse, error)
	// for task comsumers
	SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (*SubscribeTopicResponse, error)
	UnsubscribeTopic(ctx context.Context, in *UnsubscribeTopicRequest, opts ...grpc.CallOption) (*UnsubscribeTopicResponse, error)
}

type taskDelayQueueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskDelayQueueServiceClient(cc grpc.ClientConnInterface) TaskDelayQueueServiceClient {
	return &taskDelayQueueServiceClient{cc}
}

func (c *taskDelayQueueServiceClient) PushTask(ctx context.Context, in *PushTaskRequest, opts ...grpc.CallOption) (*PushTaskResponse, error) {
	out := new(PushTaskResponse)
	err := c.cc.Invoke(ctx, "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/PushTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskDelayQueueServiceClient) FinishTask(ctx context.Context, in *FinishTaskRequest, opts ...grpc.CallOption) (*FinishTaskResponse, error) {
	out := new(FinishTaskResponse)
	err := c.cc.Invoke(ctx, "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/FinishTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskDelayQueueServiceClient) CheckTask(ctx context.Context, in *CheckTaskRequest, opts ...grpc.CallOption) (*CheckTaskResponse, error) {
	out := new(CheckTaskResponse)
	err := c.cc.Invoke(ctx, "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/CheckTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskDelayQueueServiceClient) SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (*SubscribeTopicResponse, error) {
	out := new(SubscribeTopicResponse)
	err := c.cc.Invoke(ctx, "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/SubscribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskDelayQueueServiceClient) UnsubscribeTopic(ctx context.Context, in *UnsubscribeTopicRequest, opts ...grpc.CallOption) (*UnsubscribeTopicResponse, error) {
	out := new(UnsubscribeTopicResponse)
	err := c.cc.Invoke(ctx, "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/UnsubscribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskDelayQueueServiceServer is the server API for TaskDelayQueueService service.
// All implementations must embed UnimplementedTaskDelayQueueServiceServer
// for forward compatibility
type TaskDelayQueueServiceServer interface {
	// for task producers
	PushTask(context.Context, *PushTaskRequest) (*PushTaskResponse, error)
	FinishTask(context.Context, *FinishTaskRequest) (*FinishTaskResponse, error)
	CheckTask(context.Context, *CheckTaskRequest) (*CheckTaskResponse, error)
	// for task comsumers
	SubscribeTopic(context.Context, *SubscribeTopicRequest) (*SubscribeTopicResponse, error)
	UnsubscribeTopic(context.Context, *UnsubscribeTopicRequest) (*UnsubscribeTopicResponse, error)
	mustEmbedUnimplementedTaskDelayQueueServiceServer()
}

// UnimplementedTaskDelayQueueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskDelayQueueServiceServer struct {
}

func (UnimplementedTaskDelayQueueServiceServer) PushTask(context.Context, *PushTaskRequest) (*PushTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushTask not implemented")
}
func (UnimplementedTaskDelayQueueServiceServer) FinishTask(context.Context, *FinishTaskRequest) (*FinishTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTask not implemented")
}
func (UnimplementedTaskDelayQueueServiceServer) CheckTask(context.Context, *CheckTaskRequest) (*CheckTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTask not implemented")
}
func (UnimplementedTaskDelayQueueServiceServer) SubscribeTopic(context.Context, *SubscribeTopicRequest) (*SubscribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeTopic not implemented")
}
func (UnimplementedTaskDelayQueueServiceServer) UnsubscribeTopic(context.Context, *UnsubscribeTopicRequest) (*UnsubscribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribeTopic not implemented")
}
func (UnimplementedTaskDelayQueueServiceServer) mustEmbedUnimplementedTaskDelayQueueServiceServer() {}

// UnsafeTaskDelayQueueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskDelayQueueServiceServer will
// result in compilation errors.
type UnsafeTaskDelayQueueServiceServer interface {
	mustEmbedUnimplementedTaskDelayQueueServiceServer()
}

func RegisterTaskDelayQueueServiceServer(s grpc.ServiceRegistrar, srv TaskDelayQueueServiceServer) {
	s.RegisterService(&_TaskDelayQueueService_serviceDesc, srv)
}

func _TaskDelayQueueService_PushTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskDelayQueueServiceServer).PushTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/PushTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskDelayQueueServiceServer).PushTask(ctx, req.(*PushTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskDelayQueueService_FinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskDelayQueueServiceServer).FinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/FinishTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskDelayQueueServiceServer).FinishTask(ctx, req.(*FinishTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskDelayQueueService_CheckTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskDelayQueueServiceServer).CheckTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/CheckTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskDelayQueueServiceServer).CheckTask(ctx, req.(*CheckTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskDelayQueueService_SubscribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskDelayQueueServiceServer).SubscribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/SubscribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskDelayQueueServiceServer).SubscribeTopic(ctx, req.(*SubscribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskDelayQueueService_UnsubscribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskDelayQueueServiceServer).UnsubscribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amazingchow.photon_dance_delay_queue.TaskDelayQueueService/UnsubscribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskDelayQueueServiceServer).UnsubscribeTopic(ctx, req.(*UnsubscribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskDelayQueueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "amazingchow.photon_dance_delay_queue.TaskDelayQueueService",
	HandlerType: (*TaskDelayQueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushTask",
			Handler:    _TaskDelayQueueService_PushTask_Handler,
		},
		{
			MethodName: "FinishTask",
			Handler:    _TaskDelayQueueService_FinishTask_Handler,
		},
		{
			MethodName: "CheckTask",
			Handler:    _TaskDelayQueueService_CheckTask_Handler,
		},
		{
			MethodName: "SubscribeTopic",
			Handler:    _TaskDelayQueueService_SubscribeTopic_Handler,
		},
		{
			MethodName: "UnsubscribeTopic",
			Handler:    _TaskDelayQueueService_UnsubscribeTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/amazingchow/photon-dance-delay-queue/pb/task-delay-queue-service.proto",
}