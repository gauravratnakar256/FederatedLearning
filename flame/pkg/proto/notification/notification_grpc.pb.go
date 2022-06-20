// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: notification.proto

package notification

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

// EventRouteClient is the client API for EventRoute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventRouteClient interface {
	// This implements a server side streaming RPC
	GetEvent(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (EventRoute_GetEventClient, error)
}

type eventRouteClient struct {
	cc grpc.ClientConnInterface
}

func NewEventRouteClient(cc grpc.ClientConnInterface) EventRouteClient {
	return &eventRouteClient{cc}
}

func (c *eventRouteClient) GetEvent(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (EventRoute_GetEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventRoute_ServiceDesc.Streams[0], "/grpcNotification.EventRoute/GetEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventRouteGetEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventRoute_GetEventClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventRouteGetEventClient struct {
	grpc.ClientStream
}

func (x *eventRouteGetEventClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventRouteServer is the server API for EventRoute service.
// All implementations must embed UnimplementedEventRouteServer
// for forward compatibility
type EventRouteServer interface {
	// This implements a server side streaming RPC
	GetEvent(*TaskInfo, EventRoute_GetEventServer) error
	mustEmbedUnimplementedEventRouteServer()
}

// UnimplementedEventRouteServer must be embedded to have forward compatible implementations.
type UnimplementedEventRouteServer struct {
}

func (UnimplementedEventRouteServer) GetEvent(*TaskInfo, EventRoute_GetEventServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventRouteServer) mustEmbedUnimplementedEventRouteServer() {}

// UnsafeEventRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventRouteServer will
// result in compilation errors.
type UnsafeEventRouteServer interface {
	mustEmbedUnimplementedEventRouteServer()
}

func RegisterEventRouteServer(s grpc.ServiceRegistrar, srv EventRouteServer) {
	s.RegisterService(&EventRoute_ServiceDesc, srv)
}

func _EventRoute_GetEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventRouteServer).GetEvent(m, &eventRouteGetEventServer{stream})
}

type EventRoute_GetEventServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventRouteGetEventServer struct {
	grpc.ServerStream
}

func (x *eventRouteGetEventServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// EventRoute_ServiceDesc is the grpc.ServiceDesc for EventRoute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventRoute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcNotification.EventRoute",
	HandlerType: (*EventRouteServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvent",
			Handler:       _EventRoute_GetEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}

// TriggerRouteClient is the client API for TriggerRoute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TriggerRouteClient interface {
	Notify(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*Response, error)
}

type triggerRouteClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerRouteClient(cc grpc.ClientConnInterface) TriggerRouteClient {
	return &triggerRouteClient{cc}
}

func (c *triggerRouteClient) Notify(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpcNotification.TriggerRoute/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerRouteServer is the server API for TriggerRoute service.
// All implementations must embed UnimplementedTriggerRouteServer
// for forward compatibility
type TriggerRouteServer interface {
	Notify(context.Context, *EventRequest) (*Response, error)
	mustEmbedUnimplementedTriggerRouteServer()
}

// UnimplementedTriggerRouteServer must be embedded to have forward compatible implementations.
type UnimplementedTriggerRouteServer struct {
}

func (UnimplementedTriggerRouteServer) Notify(context.Context, *EventRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedTriggerRouteServer) mustEmbedUnimplementedTriggerRouteServer() {}

// UnsafeTriggerRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TriggerRouteServer will
// result in compilation errors.
type UnsafeTriggerRouteServer interface {
	mustEmbedUnimplementedTriggerRouteServer()
}

func RegisterTriggerRouteServer(s grpc.ServiceRegistrar, srv TriggerRouteServer) {
	s.RegisterService(&TriggerRoute_ServiceDesc, srv)
}

func _TriggerRoute_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerRouteServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcNotification.TriggerRoute/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerRouteServer).Notify(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TriggerRoute_ServiceDesc is the grpc.ServiceDesc for TriggerRoute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TriggerRoute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcNotification.TriggerRoute",
	HandlerType: (*TriggerRouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _TriggerRoute_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}
