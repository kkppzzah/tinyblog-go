// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: recommend.proto

package recommend

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
	RecommendService_RecommendForHome_FullMethodName    = "/recommend.RecommendService/RecommendForHome"
	RecommendService_RecommendForArticle_FullMethodName = "/recommend.RecommendService/RecommendForArticle"
)

// RecommendServiceClient is the client API for RecommendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendServiceClient interface {
	// 为网站首页进行推荐。
	RecommendForHome(ctx context.Context, in *RecommendForHomeRequest, opts ...grpc.CallOption) (*RecommendForHomeResponse, error)
	// 为单个文章进行推荐。
	RecommendForArticle(ctx context.Context, in *RecommendForArticleRequest, opts ...grpc.CallOption) (*RecommendForArticleResponse, error)
}

type recommendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendServiceClient(cc grpc.ClientConnInterface) RecommendServiceClient {
	return &recommendServiceClient{cc}
}

func (c *recommendServiceClient) RecommendForHome(ctx context.Context, in *RecommendForHomeRequest, opts ...grpc.CallOption) (*RecommendForHomeResponse, error) {
	out := new(RecommendForHomeResponse)
	err := c.cc.Invoke(ctx, RecommendService_RecommendForHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendServiceClient) RecommendForArticle(ctx context.Context, in *RecommendForArticleRequest, opts ...grpc.CallOption) (*RecommendForArticleResponse, error) {
	out := new(RecommendForArticleResponse)
	err := c.cc.Invoke(ctx, RecommendService_RecommendForArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendServiceServer is the server API for RecommendService service.
// All implementations should embed UnimplementedRecommendServiceServer
// for forward compatibility
type RecommendServiceServer interface {
	// 为网站首页进行推荐。
	RecommendForHome(context.Context, *RecommendForHomeRequest) (*RecommendForHomeResponse, error)
	// 为单个文章进行推荐。
	RecommendForArticle(context.Context, *RecommendForArticleRequest) (*RecommendForArticleResponse, error)
}

// UnimplementedRecommendServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRecommendServiceServer struct {
}

func (UnimplementedRecommendServiceServer) RecommendForHome(context.Context, *RecommendForHomeRequest) (*RecommendForHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendForHome not implemented")
}
func (UnimplementedRecommendServiceServer) RecommendForArticle(context.Context, *RecommendForArticleRequest) (*RecommendForArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendForArticle not implemented")
}

// UnsafeRecommendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendServiceServer will
// result in compilation errors.
type UnsafeRecommendServiceServer interface {
	mustEmbedUnimplementedRecommendServiceServer()
}

func RegisterRecommendServiceServer(s grpc.ServiceRegistrar, srv RecommendServiceServer) {
	s.RegisterService(&RecommendService_ServiceDesc, srv)
}

func _RecommendService_RecommendForHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendForHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).RecommendForHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_RecommendForHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).RecommendForHome(ctx, req.(*RecommendForHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendService_RecommendForArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendForArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).RecommendForArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_RecommendForArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).RecommendForArticle(ctx, req.(*RecommendForArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendService_ServiceDesc is the grpc.ServiceDesc for RecommendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recommend.RecommendService",
	HandlerType: (*RecommendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecommendForHome",
			Handler:    _RecommendService_RecommendForHome_Handler,
		},
		{
			MethodName: "RecommendForArticle",
			Handler:    _RecommendService_RecommendForArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommend.proto",
}