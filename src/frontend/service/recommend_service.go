// Package service 各个服务。
package service

import (
	"context"

	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/recommend"
)

// RecommendService 推荐服务。
type RecommendService struct {
	RPCServiceBase
}

// NewRecommendService 创建推荐服务实例。
func NewRecommendService(mustConnect bool) *RecommendService {
	svc := &RecommendService{
		RPCServiceBase: RPCServiceBase{
			name:           "recommend",
			serviceAddress: common.MustGetEnv(common.EnvRecommendServiceAddress, ""),
		},
	}
	svc.initialize(mustConnect)
	return svc
}

// RecommendForHome 首页推荐。
func (svc *RecommendService) RecommendForHome(ctx context.Context, userID int64) ([]int64, error) {
	rsp, err := pb.NewRecommendServiceClient(svc.GetConnection()).RecommendForHome(ctx, &pb.RecommendForHomeRequest{UserId: userID})
	if err == nil {
		return rsp.ArticleIds, err
	}
	return []int64{}, err
}

// RecommendForArticle 对特定文章的推荐。
func (svc *RecommendService) RecommendForArticle(ctx context.Context, articleID int64, userID int64) ([]int64, error) {
	rsp, err := pb.NewRecommendServiceClient(svc.GetConnection()).RecommendForArticle(ctx, &pb.RecommendForArticleRequest{ArticleId: articleID, UserId: userID})
	if err == nil {
		return rsp.ArticleIds, err
	}
	return []int64{}, err
}
