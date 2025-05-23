// protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative subscription.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: subscription.proto

package subscription

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
	SubscriptionService_Subscribe_FullMethodName                   = "/subscription.SubscriptionService/Subscribe"
	SubscriptionService_Unsubscribe_FullMethodName                 = "/subscription.SubscriptionService/Unsubscribe"
	SubscriptionService_GetSubscriptions_FullMethodName            = "/subscription.SubscriptionService/GetSubscriptions"
	SubscriptionService_GetSubscribers_FullMethodName              = "/subscription.SubscriptionService/GetSubscribers"
	SubscriptionService_CheckSubscription_FullMethodName           = "/subscription.SubscriptionService/CheckSubscription"
	SubscriptionService_GetWatchlistsBySubscription_FullMethodName = "/subscription.SubscriptionService/GetWatchlistsBySubscription"
	SubscriptionService_GetReviewsBySubscription_FullMethodName    = "/subscription.SubscriptionService/GetReviewsBySubscription"
)

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления подписками
type SubscriptionServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error)
	GetSubscribers(ctx context.Context, in *GetSubscribersRequest, opts ...grpc.CallOption) (*GetSubscribersResponse, error)
	CheckSubscription(ctx context.Context, in *CheckSubscriptionRequest, opts ...grpc.CallOption) (*CheckSubscriptionResponse, error)
	GetWatchlistsBySubscription(ctx context.Context, in *GetWatchlistsRequest, opts ...grpc.CallOption) (*GetWatchlistsResponse, error)
	GetReviewsBySubscription(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_Subscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_Unsubscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubscriptionsResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetSubscriptions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscribers(ctx context.Context, in *GetSubscribersRequest, opts ...grpc.CallOption) (*GetSubscribersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubscribersResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetSubscribers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CheckSubscription(ctx context.Context, in *CheckSubscriptionRequest, opts ...grpc.CallOption) (*CheckSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_CheckSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetWatchlistsBySubscription(ctx context.Context, in *GetWatchlistsRequest, opts ...grpc.CallOption) (*GetWatchlistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWatchlistsResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetWatchlistsBySubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetReviewsBySubscription(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewsResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetReviewsBySubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility.
//
// Сервис для управления подписками
type SubscriptionServiceServer interface {
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error)
	GetSubscribers(context.Context, *GetSubscribersRequest) (*GetSubscribersResponse, error)
	CheckSubscription(context.Context, *CheckSubscriptionRequest) (*CheckSubscriptionResponse, error)
	GetWatchlistsBySubscription(context.Context, *GetWatchlistsRequest) (*GetWatchlistsResponse, error)
	GetReviewsBySubscription(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubscriptionServiceServer struct{}

func (UnimplementedSubscriptionServiceServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubscriptionServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetSubscribers(context.Context, *GetSubscribersRequest) (*GetSubscribersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribers not implemented")
}
func (UnimplementedSubscriptionServiceServer) CheckSubscription(context.Context, *CheckSubscriptionRequest) (*CheckSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetWatchlistsBySubscription(context.Context, *GetWatchlistsRequest) (*GetWatchlistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWatchlistsBySubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetReviewsBySubscription(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewsBySubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}
func (UnimplementedSubscriptionServiceServer) testEmbeddedByValue()                             {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubscriptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_Unsubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscriptions(ctx, req.(*GetSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscribersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetSubscribers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscribers(ctx, req.(*GetSubscribersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CheckSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CheckSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CheckSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CheckSubscription(ctx, req.(*CheckSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetWatchlistsBySubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWatchlistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetWatchlistsBySubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetWatchlistsBySubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetWatchlistsBySubscription(ctx, req.(*GetWatchlistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetReviewsBySubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetReviewsBySubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetReviewsBySubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetReviewsBySubscription(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _SubscriptionService_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _SubscriptionService_Unsubscribe_Handler,
		},
		{
			MethodName: "GetSubscriptions",
			Handler:    _SubscriptionService_GetSubscriptions_Handler,
		},
		{
			MethodName: "GetSubscribers",
			Handler:    _SubscriptionService_GetSubscribers_Handler,
		},
		{
			MethodName: "CheckSubscription",
			Handler:    _SubscriptionService_CheckSubscription_Handler,
		},
		{
			MethodName: "GetWatchlistsBySubscription",
			Handler:    _SubscriptionService_GetWatchlistsBySubscription_Handler,
		},
		{
			MethodName: "GetReviewsBySubscription",
			Handler:    _SubscriptionService_GetReviewsBySubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
