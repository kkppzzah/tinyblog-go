// Package service 各个服务。
package service

import (
	"context"

	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/search"
)

// SearchService 搜索服务。
type SearchService struct {
	RPCServiceBase
}

// NewSearchService 创建搜索服务。。
func NewSearchService(mustConnect bool) *SearchService {
	svc := &SearchService{
		RPCServiceBase: RPCServiceBase{
			name:           "search",
			serviceAddress: common.MustGetEnv(common.EnvSearchServiceAddress, ""),
		},
	}
	svc.initialize(mustConnect)
	return svc
}

// SimpleSearch 简单搜索。
func (svc *SearchService) SimpleSearch(ctx context.Context, req *pb.SimpleSearchRequest) (*pb.SimpleSearchResponse, error) {
	return svc.getServiceClient().SimpleSearch(ctx, req)
}

func (svc *SearchService) getServiceClient() pb.SearchServiceClient {
	return pb.NewSearchServiceClient(svc.GetConnection())
}
