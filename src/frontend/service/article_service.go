// Package service 各个服务。
package service

import (
	"context"

	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/article"
)

// ArticleService 推荐服务。
type ArticleService struct {
	RPCServiceBase
}

// NewArticleService 创建推荐服务实例。
func NewArticleService(mustConnect bool) *ArticleService {
	svc := &ArticleService{
		RPCServiceBase: RPCServiceBase{
			name:           "article",
			serviceAddress: common.MustGetEnv(common.EnvArticleServiceAddress, ""),
		},
	}
	svc.initialize(mustConnect)
	return svc
}

// Publish 发表文章。
func (svc *ArticleService) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	return pb.NewArticleServiceClient(svc.GetConnection()).Publish(ctx, req)
}

// Update 更新文章。
func (svc *ArticleService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Empty, error) {
	return pb.NewArticleServiceClient(svc.GetConnection()).Update(ctx, req)
}

// Delete 删除文章。
func (svc *ArticleService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	return pb.NewArticleServiceClient(svc.GetConnection()).Delete(ctx, req)
}

// Get 获取单个文章内容。
func (svc *ArticleService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return pb.NewArticleServiceClient(svc.GetConnection()).Get(ctx, req)
}

// GetByUser 获取单个用户的多个文章。
func (svc *ArticleService) GetByUser(ctx context.Context, req *pb.GetByUserRequest) (*pb.GetByUserResponse, error) {
	return pb.NewArticleServiceClient(svc.GetConnection()).GetByUser(ctx, req)
}

// GetByIds 获取指定文章id列表的多个文章。
func (svc *ArticleService) GetByIds(ctx context.Context, ids []int64) (*pb.GetByIdsResponse, error) {
	return pb.NewArticleServiceClient(svc.GetConnection()).GetByIds(ctx, &pb.GetByIdsRequest{Ids: ids})
}
