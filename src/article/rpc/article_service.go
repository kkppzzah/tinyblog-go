// Package rpc 处理gRPC服务请求。
package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ppzzl.com/tinyblog-go/article/common"
	pb "ppzzl.com/tinyblog-go/article/genproto/article"
	"ppzzl.com/tinyblog-go/article/interfaces"
	"ppzzl.com/tinyblog-go/article/model"
)

// ArticleService 实现文章服务。
type ArticleService struct {
	articleRepository     interfaces.ArticleRepository
	articleEventPublisher interfaces.ArticleEventPublisher
}

// NewArticleService 创建ArticleService实例。
func NewArticleService(context interfaces.Context) *ArticleService {
	rs := &ArticleService{
		articleRepository:     context.GetArticleRepository(),
		articleEventPublisher: context.GetArticleEventPublisher(),
	}
	return rs
}

// Publish 发表文章。
func (svc *ArticleService) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	article := &model.Article{
		UserID:      req.UserId,
		Title:       req.Title,
		Tags:        req.Tags,
		Summary:     req.Summary,
		Content:     req.Content,
		PublishTime: time.Now(),
	}

	article, err := svc.articleRepository.Create(article)
	if err != nil {
		log.Printf("failed to creart article, user_id: %d, title: %s, error: %v", article.UserID, article.Title, err)
		return nil, err
	}

	rsp := &pb.PublishResponse{
		Id: article.ID,
	}
	// 发布文章发布事件。
	err = svc.articleEventPublisher.Publish(&interfaces.ArticleEvent{
		EventType: interfaces.ArticleEventTypeCreate,
		ArticleInfo: &interfaces.ArticleInfo{
			ID: article.ID,
		},
	})
	return rsp, nil
}

// Update 更新文章。
func (svc *ArticleService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Empty, error) {
	article := &model.Article{
		ID:      req.Id,
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
	}

	err := svc.articleRepository.Update(article)
	if err != nil {
		switch v := err.(type) {
		case *common.Error:
			if v.Code == common.ErrorCodeNoFound {
				return &pb.Empty{}, status.Error(codes.NotFound, fmt.Sprintf("article %d is not found", req.Id))
			}
			return &pb.Empty{}, status.Error(codes.Internal, v.Error())
		default:
			return &pb.Empty{}, status.Error(codes.Internal, v.Error())
		}
	}
	// 发布文章更新事件。
	err = svc.articleEventPublisher.Publish(&interfaces.ArticleEvent{
		EventType: interfaces.ArticleEventTypeUpdate,
		ArticleInfo: &interfaces.ArticleInfo{
			ID: article.ID,
		},
	})
	return &pb.Empty{}, nil
}

// Delete 删除文章。
func (svc *ArticleService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	err := svc.articleRepository.Delete(req.Id)
	if err != nil {
		return &pb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	// 发布文章删除事件。
	err = svc.articleEventPublisher.Publish(&interfaces.ArticleEvent{
		EventType: interfaces.ArticleEventTypeDelete,
		ArticleInfo: &interfaces.ArticleInfo{
			ID: req.Id,
		},
	})
	return &pb.Empty{}, nil
}

// Get 获取单个文章内容。
func (svc *ArticleService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	article, err := svc.articleRepository.Get(req.Id)
	if err != nil {
		log.Printf("failed to get article, article_id: %d", req.Id)
		return nil, err
	}

	if article == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("article %d is not found", req.Id))
	}

	rsp := &pb.GetResponse{
		Id:          article.ID,
		UserId:      article.UserID,
		Title:       article.Title,
		Tags:        article.Tags,
		Summary:     article.Summary,
		Content:     article.Content,
		PublishTime: article.PublishTime.Unix(),
		Nickname:    article.Nickname,
	}
	return rsp, nil
}

// GetByUser 获取单个用户的多个文章。
func (svc *ArticleService) GetByUser(ctx context.Context, req *pb.GetByUserRequest) (*pb.GetByUserResponse, error) {
	articles, err := svc.articleRepository.GetByUser(req.UserId)
	if err != nil {
		return nil, err
	}
	rspArticles := make([]*pb.UserArticle, len(articles))
	for idx, article := range articles {
		rspArticles[idx] = &pb.UserArticle{
			Id:          article.ID,
			UserId:      article.UserID,
			Title:       article.Title,
			Tags:        article.Tags,
			Summary:     article.Summary,
			UserName:    article.Nickname,
			PublishTime: article.PublishTime.Unix(),
			Nickname:    article.Nickname,
		}
	}
	rsp := &pb.GetByUserResponse{
		Articles: rspArticles,
	}
	return rsp, nil
}

// GetByIds 获取指定文章id列表的多个文章。
func (svc *ArticleService) GetByIds(ctx context.Context, req *pb.GetByIdsRequest) (*pb.GetByIdsResponse, error) {
	articles, err := svc.articleRepository.GetByIds(req.Ids)
	if err != nil {
		return nil, err
	}
	rspArticles := make([]*pb.UserArticle, len(articles))
	for idx, article := range articles {
		rspArticles[idx] = &pb.UserArticle{
			Id:          article.ID,
			UserId:      article.UserID,
			Title:       article.Title,
			Tags:        article.Tags,
			Summary:     article.Summary,
			UserName:    article.Nickname,
			PublishTime: article.PublishTime.Unix(),
			Nickname:    article.Nickname,
		}
	}
	rsp := &pb.GetByIdsResponse{
		Articles: rspArticles,
	}
	return rsp, nil
}
