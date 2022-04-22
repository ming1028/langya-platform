// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1--rc1
// source: langya_platform_app.proto

package app

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

// LangYaPlatformClient is the client API for LangYaPlatform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LangYaPlatformClient interface {
	ServiceContractBook(ctx context.Context, in *ContractBookReq, opts ...grpc.CallOption) (*ContractBook, error)
}

type langYaPlatformClient struct {
	cc grpc.ClientConnInterface
}

func NewLangYaPlatformClient(cc grpc.ClientConnInterface) LangYaPlatformClient {
	return &langYaPlatformClient{cc}
}

func (c *langYaPlatformClient) ServiceContractBook(ctx context.Context, in *ContractBookReq, opts ...grpc.CallOption) (*ContractBook, error) {
	out := new(ContractBook)
	err := c.cc.Invoke(ctx, "/app.LangYaPlatform/ServiceContractBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LangYaPlatformServer is the server API for LangYaPlatform service.
// All implementations must embed UnimplementedLangYaPlatformServer
// for forward compatibility
type LangYaPlatformServer interface {
	ServiceContractBook(context.Context, *ContractBookReq) (*ContractBook, error)
	mustEmbedUnimplementedLangYaPlatformServer()
}

// UnimplementedLangYaPlatformServer must be embedded to have forward compatible implementations.
type UnimplementedLangYaPlatformServer struct {
}

func (UnimplementedLangYaPlatformServer) ServiceContractBook(context.Context, *ContractBookReq) (*ContractBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceContractBook not implemented")
}
func (UnimplementedLangYaPlatformServer) mustEmbedUnimplementedLangYaPlatformServer() {}

// UnsafeLangYaPlatformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LangYaPlatformServer will
// result in compilation errors.
type UnsafeLangYaPlatformServer interface {
	mustEmbedUnimplementedLangYaPlatformServer()
}

func RegisterLangYaPlatformServer(s grpc.ServiceRegistrar, srv LangYaPlatformServer) {
	s.RegisterService(&LangYaPlatform_ServiceDesc, srv)
}

func _LangYaPlatform_ServiceContractBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LangYaPlatformServer).ServiceContractBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.LangYaPlatform/ServiceContractBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LangYaPlatformServer).ServiceContractBook(ctx, req.(*ContractBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LangYaPlatform_ServiceDesc is the grpc.ServiceDesc for LangYaPlatform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LangYaPlatform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.LangYaPlatform",
	HandlerType: (*LangYaPlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceContractBook",
			Handler:    _LangYaPlatform_ServiceContractBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "langya_platform_app.proto",
}
