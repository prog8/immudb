// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authorizationschema

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

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	OpenSessionV2(ctx context.Context, in *OpenSessionRequestV2, opts ...grpc.CallOption) (*OpenSessionResponseV2, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) OpenSessionV2(ctx context.Context, in *OpenSessionRequestV2, opts ...grpc.CallOption) (*OpenSessionResponseV2, error) {
	out := new(OpenSessionResponseV2)
	err := c.cc.Invoke(ctx, "/immudb.authorizationschema.AuthorizationService/OpenSessionV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations should embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	OpenSessionV2(context.Context, *OpenSessionRequestV2) (*OpenSessionResponseV2, error)
}

// UnimplementedAuthorizationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) OpenSessionV2(context.Context, *OpenSessionRequestV2) (*OpenSessionResponseV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSessionV2 not implemented")
}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_OpenSessionV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSessionRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).OpenSessionV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.authorizationschema.AuthorizationService/OpenSessionV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).OpenSessionV2(ctx, req.(*OpenSessionRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "immudb.authorizationschema.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenSessionV2",
			Handler:    _AuthorizationService_OpenSessionV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorizationschema.proto",
}
