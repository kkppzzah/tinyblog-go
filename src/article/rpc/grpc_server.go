// Package rpc 处理gRPC服务请求。
package rpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"ppzzl.com/tinyblog-go/article/common"
	pb "ppzzl.com/tinyblog-go/article/genproto/article"
	"ppzzl.com/tinyblog-go/article/interfaces"
)

// Server gRPC服务端。
type Server struct {
	address    string
	grpcServer *grpc.Server
	context    interfaces.Context
}

// NewServer 创建gRPC服务端。
func NewServer(ctx interfaces.Context) *Server {
	s := &Server{
		address:    common.MustGetEnv(common.EnvVarNameListenAddress, ""),
		grpcServer: grpc.NewServer(),
		context:    ctx,
	}
	s.registerGRPCServcies()
	return s
}

// Run 运行gRPC服务端。
func (s *Server) Run() {
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("failed to listen '%s', %v", s.address, err)
	}
	s.grpcServer.Serve(l)
}

func (s *Server) registerGRPCServcies() {
	pb.RegisterArticleServiceServer(s.grpcServer, NewArticleService(s.context))
}
