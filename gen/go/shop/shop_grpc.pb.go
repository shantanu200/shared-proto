// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: shop/shop.proto

package shop

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
	ShopService_GetShop_FullMethodName                = "/shop.ShopService/GetShop"
	ShopService_GetStoreAdminAnalytics_FullMethodName = "/shop.ShopService/GetStoreAdminAnalytics"
	ShopService_GetStoreCountAnalytics_FullMethodName = "/shop.ShopService/GetStoreCountAnalytics"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	GetShop(ctx context.Context, in *ShopReq, opts ...grpc.CallOption) (*Shop, error)
	GetStoreAdminAnalytics(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*StoreAdminAnalyticsResult, error)
	GetStoreCountAnalytics(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*StoreCountAnalytics, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *ShopReq, opts ...grpc.CallOption) (*Shop, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shop)
	err := c.cc.Invoke(ctx, ShopService_GetShop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStoreAdminAnalytics(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*StoreAdminAnalyticsResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreAdminAnalyticsResult)
	err := c.cc.Invoke(ctx, ShopService_GetStoreAdminAnalytics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStoreCountAnalytics(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*StoreCountAnalytics, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreCountAnalytics)
	err := c.cc.Invoke(ctx, ShopService_GetStoreCountAnalytics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility.
type ShopServiceServer interface {
	GetShop(context.Context, *ShopReq) (*Shop, error)
	GetStoreAdminAnalytics(context.Context, *EmptyReq) (*StoreAdminAnalyticsResult, error)
	GetStoreCountAnalytics(context.Context, *EmptyReq) (*StoreCountAnalytics, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShopServiceServer struct{}

func (UnimplementedShopServiceServer) GetShop(context.Context, *ShopReq) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}
func (UnimplementedShopServiceServer) GetStoreAdminAnalytics(context.Context, *EmptyReq) (*StoreAdminAnalyticsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreAdminAnalytics not implemented")
}
func (UnimplementedShopServiceServer) GetStoreCountAnalytics(context.Context, *EmptyReq) (*StoreCountAnalytics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreCountAnalytics not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}
func (UnimplementedShopServiceServer) testEmbeddedByValue()                     {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	// If the following call pancis, it indicates UnimplementedShopServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*ShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStoreAdminAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStoreAdminAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetStoreAdminAnalytics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStoreAdminAnalytics(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStoreCountAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStoreCountAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetStoreCountAnalytics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStoreCountAnalytics(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "GetStoreAdminAnalytics",
			Handler:    _ShopService_GetStoreAdminAnalytics_Handler,
		},
		{
			MethodName: "GetStoreCountAnalytics",
			Handler:    _ShopService_GetStoreCountAnalytics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/shop.proto",
}
