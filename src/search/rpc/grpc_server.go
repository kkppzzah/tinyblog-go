// Package rpc 处理gRPC服务请求。
package rpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"ppzzl.com/tinyblog-go/search/common"
	healthpb "ppzzl.com/tinyblog-go/search/genproto/grpc/health/v1"
	pb "ppzzl.com/tinyblog-go/search/genproto/search"
	"ppzzl.com/tinyblog-go/search/interfaces"
)

// Server gRPC服务端。
type Server struct {
	address    string
	grpcServer *grpc.Server
	ctx        interfaces.Context
}

// NewServer 创建gRPC服务端。
func NewServer(ctx interfaces.Context) *Server {
	s := &Server{
		address:    common.MustGetEnv(common.EnvVarNameListenAddress, ""),
		grpcServer: grpc.NewServer(),
		ctx:        ctx,
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
	pb.RegisterSearchServiceServer(s.grpcServer, NewSearchService(s.ctx))
	healthpb.RegisterHealthServer(s.grpcServer, NewHealthService(s.ctx))
}
