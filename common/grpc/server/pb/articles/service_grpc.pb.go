// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: articles/service.proto

package articles_pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArticleService_CreateArticle_FullMethodName        = "/articles.ArticleService/CreateArticle"
	ArticleService_ListArticle_FullMethodName          = "/articles.ArticleService/ListArticle"
	ArticleService_GetArticle_FullMethodName           = "/articles.ArticleService/GetArticle"
	ArticleService_UpdateArticleBody_FullMethodName    = "/articles.ArticleService/UpdateArticleBody"
	ArticleService_UpdateArticleTitle_FullMethodName   = "/articles.ArticleService/UpdateArticleTitle"
	ArticleService_DeleteArticle_FullMethodName        = "/articles.ArticleService/DeleteArticle"
	ArticleService_SetArticleVisibility_FullMethodName = "/articles.ArticleService/SetArticleVisibility"
	ArticleService_ToggleArticleLike_FullMethodName    = "/articles.ArticleService/ToggleArticleLike"
)

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (*ListArticleResponse, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error)
	UpdateArticleBody(ctx context.Context, in *UpdateArticleBodyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateArticleTitle(ctx context.Context, in *UpdateArticleTitleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetArticleVisibility(ctx context.Context, in *SetArticleVisibilityRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ToggleArticleLike(ctx context.Context, in *ToggleArticleLikeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ArticleService_CreateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (*ListArticleResponse, error) {
	out := new(ListArticleResponse)
	err := c.cc.Invoke(ctx, ArticleService_ListArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error) {
	out := new(GetArticleResponse)
	err := c.cc.Invoke(ctx, ArticleService_GetArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticleBody(ctx context.Context, in *UpdateArticleBodyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ArticleService_UpdateArticleBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticleTitle(ctx context.Context, in *UpdateArticleTitleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ArticleService_UpdateArticleTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ArticleService_DeleteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) SetArticleVisibility(ctx context.Context, in *SetArticleVisibilityRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ArticleService_SetArticleVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ToggleArticleLike(ctx context.Context, in *ToggleArticleLikeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ArticleService_ToggleArticleLike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*empty.Empty, error)
	ListArticle(context.Context, *ListArticleRequest) (*ListArticleResponse, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error)
	UpdateArticleBody(context.Context, *UpdateArticleBodyRequest) (*empty.Empty, error)
	UpdateArticleTitle(context.Context, *UpdateArticleTitleRequest) (*empty.Empty, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*empty.Empty, error)
	SetArticleVisibility(context.Context, *SetArticleVisibilityRequest) (*empty.Empty, error)
	ToggleArticleLike(context.Context, *ToggleArticleLikeRequest) (*empty.Empty, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *CreateArticleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) ListArticle(context.Context, *ListArticleRequest) (*ListArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticleBody(context.Context, *UpdateArticleBodyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticleBody not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticleTitle(context.Context, *UpdateArticleTitleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticleTitle not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServiceServer) SetArticleVisibility(context.Context, *SetArticleVisibilityRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArticleVisibility not implemented")
}
func (UnimplementedArticleServiceServer) ToggleArticleLike(context.Context, *ToggleArticleLikeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleArticleLike not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ListArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ListArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_ListArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ListArticle(ctx, req.(*ListArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticleBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleBodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticleBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_UpdateArticleBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticleBody(ctx, req.(*UpdateArticleBodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticleTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticleTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_UpdateArticleTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticleTitle(ctx, req.(*UpdateArticleTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_DeleteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_SetArticleVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArticleVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).SetArticleVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_SetArticleVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).SetArticleVisibility(ctx, req.(*SetArticleVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ToggleArticleLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleArticleLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ToggleArticleLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArticleService_ToggleArticleLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ToggleArticleLike(ctx, req.(*ToggleArticleLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "articles.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "ListArticle",
			Handler:    _ArticleService_ListArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ArticleService_GetArticle_Handler,
		},
		{
			MethodName: "UpdateArticleBody",
			Handler:    _ArticleService_UpdateArticleBody_Handler,
		},
		{
			MethodName: "UpdateArticleTitle",
			Handler:    _ArticleService_UpdateArticleTitle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticleService_DeleteArticle_Handler,
		},
		{
			MethodName: "SetArticleVisibility",
			Handler:    _ArticleService_SetArticleVisibility_Handler,
		},
		{
			MethodName: "ToggleArticleLike",
			Handler:    _ArticleService_ToggleArticleLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "articles/service.proto",
}
