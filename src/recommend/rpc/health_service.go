// Package rpc 处理gRPC服务请求。
package rpc

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	pb "ppzzl.com/tinyblog-go/recommend/genproto/grpc/health/v1"
	"ppzzl.com/tinyblog-go/recommend/interfaces"
)

// HealthService 实现健康状态反馈服务。
type HealthService struct {
	esClient *elasticsearch.Client
	ctx      context.Context
}

// NewHealthService 创建HealthService实例。
func NewHealthService(ctx interfaces.Context) *HealthService {
	hs := &HealthService{}
	return hs
}

// Check 处理Health检测请求。
func (svc *HealthService) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	rsp := &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}
	return rsp, nil
}
