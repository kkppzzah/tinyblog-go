// Package rpc 处理gRPC服务请求。
package rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "ppzzl.com/tinyblog-go/recommend/genproto/recommend"
	"ppzzl.com/tinyblog-go/recommend/interfaces"
)

// RecommendService 实现推荐服务。
type RecommendService struct {
	articleRepository interfaces.ArticleRepository
}

// NewRecommendService 创建RecommendService实例。
func NewRecommendService(ctx interfaces.Context) *RecommendService {
	rs := &RecommendService{
		articleRepository: ctx.GetArticleRepository(),
	}
	return rs
}

// RecommendForHome 为网站首页进行推荐。
func (rs *RecommendService) RecommendForHome(context.Context, *pb.RecommendForHomeRequest) (*pb.RecommendForHomeResponse, error) {
	articles, err := rs.articleRepository.GetRandomly(20)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	articlesIds := make([]int64, len(articles))
	for i, article := range articles {
		articlesIds[i] = article.ID
	}
	rsp := pb.RecommendForHomeResponse{
		ArticleIds: articlesIds,
	}
	return &rsp, nil
}

// RecommendForArticle 为单个文章进行推荐。
func (rs *RecommendService) RecommendForArticle(context.Context, *pb.RecommendForArticleRequest) (*pb.RecommendForArticleResponse, error) {
	articles, err := rs.articleRepository.GetRandomly(20)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	articlesIds := make([]int64, len(articles))
	for i, article := range articles {
		articlesIds[i] = article.ID
	}
	rsp := pb.RecommendForArticleResponse{
		ArticleIds: articlesIds,
	}
	return &rsp, nil
}
