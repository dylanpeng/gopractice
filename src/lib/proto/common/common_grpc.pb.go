// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: common.proto

package common

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

const (
	CommonService_CommonTest_FullMethodName = "/common.CommonService/CommonTest"
)

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonServiceClient interface {
	CommonTest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type commonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonServiceClient(cc grpc.ClientConnInterface) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) CommonTest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, CommonService_CommonTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
// All implementations must embed UnimplementedCommonServiceServer
// for forward compatibility
type CommonServiceServer interface {
	CommonTest(context.Context, *Empty) (*Response, error)
	mustEmbedUnimplementedCommonServiceServer()
}

// UnimplementedCommonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (UnimplementedCommonServiceServer) CommonTest(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommonTest not implemented")
}
func (UnimplementedCommonServiceServer) mustEmbedUnimplementedCommonServiceServer() {}

// UnsafeCommonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServiceServer will
// result in compilation errors.
type UnsafeCommonServiceServer interface {
	mustEmbedUnimplementedCommonServiceServer()
}

func RegisterCommonServiceServer(s grpc.ServiceRegistrar, srv CommonServiceServer) {
	s.RegisterService(&CommonService_ServiceDesc, srv)
}

func _CommonService_CommonTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).CommonTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommonService_CommonTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).CommonTest(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CommonService_ServiceDesc is the grpc.ServiceDesc for CommonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommonTest",
			Handler:    _CommonService_CommonTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}