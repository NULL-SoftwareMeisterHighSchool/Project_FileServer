// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: comments/service.proto

package comments_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CommentService_CreateComment_FullMethodName          = "/commnets.CommentService/CreateComment"
	CommentService_DeleteCommnet_FullMethodName          = "/commnets.CommentService/DeleteCommnet"
	CommentService_GetCommentsByArticleID_FullMethodName = "/commnets.CommentService/GetCommentsByArticleID"
	CommentService_GetRepliesByCommentID_FullMethodName  = "/commnets.CommentService/GetRepliesByCommentID"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCommnet(ctx context.Context, in *DeleteCommnetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCommentsByArticleID(ctx context.Context, in *GetCommentsByArticleIDRequest, opts ...grpc.CallOption) (*GetCommentsByArticleIDResponse, error)
	GetRepliesByCommentID(ctx context.Context, in *GetRepliesByCommentIDRequest, opts ...grpc.CallOption) (*GetRepliesByCommentIDResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CommentService_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) DeleteCommnet(ctx context.Context, in *DeleteCommnetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CommentService_DeleteCommnet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentsByArticleID(ctx context.Context, in *GetCommentsByArticleIDRequest, opts ...grpc.CallOption) (*GetCommentsByArticleIDResponse, error) {
	out := new(GetCommentsByArticleIDResponse)
	err := c.cc.Invoke(ctx, CommentService_GetCommentsByArticleID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetRepliesByCommentID(ctx context.Context, in *GetRepliesByCommentIDRequest, opts ...grpc.CallOption) (*GetRepliesByCommentIDResponse, error) {
	out := new(GetRepliesByCommentIDResponse)
	err := c.cc.Invoke(ctx, CommentService_GetRepliesByCommentID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*emptypb.Empty, error)
	DeleteCommnet(context.Context, *DeleteCommnetRequest) (*emptypb.Empty, error)
	GetCommentsByArticleID(context.Context, *GetCommentsByArticleIDRequest) (*GetCommentsByArticleIDResponse, error)
	GetRepliesByCommentID(context.Context, *GetRepliesByCommentIDRequest) (*GetRepliesByCommentIDResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) DeleteCommnet(context.Context, *DeleteCommnetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommnet not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentsByArticleID(context.Context, *GetCommentsByArticleIDRequest) (*GetCommentsByArticleIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByArticleID not implemented")
}
func (UnimplementedCommentServiceServer) GetRepliesByCommentID(context.Context, *GetRepliesByCommentIDRequest) (*GetRepliesByCommentIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepliesByCommentID not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_DeleteCommnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).DeleteCommnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_DeleteCommnet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).DeleteCommnet(ctx, req.(*DeleteCommnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentsByArticleID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByArticleIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentsByArticleID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetCommentsByArticleID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentsByArticleID(ctx, req.(*GetCommentsByArticleIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetRepliesByCommentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepliesByCommentIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetRepliesByCommentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetRepliesByCommentID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetRepliesByCommentID(ctx, req.(*GetRepliesByCommentIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commnets.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "DeleteCommnet",
			Handler:    _CommentService_DeleteCommnet_Handler,
		},
		{
			MethodName: "GetCommentsByArticleID",
			Handler:    _CommentService_GetCommentsByArticleID_Handler,
		},
		{
			MethodName: "GetRepliesByCommentID",
			Handler:    _CommentService_GetRepliesByCommentID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comments/service.proto",
}
