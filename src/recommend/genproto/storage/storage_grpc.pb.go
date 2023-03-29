// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: storage.proto

package storage

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
	StorageService_Upload_FullMethodName        = "/storage.StorageService/Upload"
	StorageService_UploadBigFile_FullMethodName = "/storage.StorageService/UploadBigFile"
)

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	// 上传文件（不适合大文件）。
	Upload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
	// 上传文件（适合大文件）
	UploadBigFile(ctx context.Context, opts ...grpc.CallOption) (StorageService_UploadBigFileClient, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) Upload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, StorageService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) UploadBigFile(ctx context.Context, opts ...grpc.CallOption) (StorageService_UploadBigFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[0], StorageService_UploadBigFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServiceUploadBigFileClient{stream}
	return x, nil
}

type StorageService_UploadBigFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type storageServiceUploadBigFileClient struct {
	grpc.ClientStream
}

func (x *storageServiceUploadBigFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageServiceUploadBigFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations should embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	// 上传文件（不适合大文件）。
	Upload(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	// 上传文件（适合大文件）
	UploadBigFile(StorageService_UploadBigFileServer) error
}

// UnimplementedStorageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) Upload(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedStorageServiceServer) UploadBigFile(StorageService_UploadBigFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBigFile not implemented")
}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Upload(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_UploadBigFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceServer).UploadBigFile(&storageServiceUploadBigFileServer{stream})
}

type StorageService_UploadBigFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type storageServiceUploadBigFileServer struct {
	grpc.ServerStream
}

func (x *storageServiceUploadBigFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageServiceUploadBigFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _StorageService_Upload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadBigFile",
			Handler:       _StorageService_UploadBigFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}
