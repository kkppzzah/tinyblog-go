// Package app 组织应用的各个组件。
package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ppzzl.com/tinyblog-go/recommend/common"
	"ppzzl.com/tinyblog-go/recommend/interfaces"
	"ppzzl.com/tinyblog-go/recommend/repository"
	"ppzzl.com/tinyblog-go/recommend/rpc"
	"ppzzl.com/tinyblog-go/recommend/service"
)

// ContextImpl 存放各个应用组件。
type ContextImpl struct {
	gRPCServer          *rpc.Server
	recommendDB         *gorm.DB
	articleRepository   interfaces.ArticleRepository
	articleEventHandler interfaces.ArticleEventHandler
}

// NewContextImpl 创建Context实例。
func NewContextImpl() *ContextImpl {
	ctx := &ContextImpl{}
	// 创建各个服务。
	recommendDB, err := gorm.Open(mysql.Open(common.MustLoadSecretAsString(common.EnvRecommendDBConnStr, common.EnvRecommendDBConnStrSecretFile)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database (recommend), %v", err))
	}
	ctx.recommendDB = recommendDB
	// 数据存储接口。
	ctx.articleRepository = repository.NewArticleRepositoryDB(ctx)
	// 事件处理。
	ctx.articleEventHandler = service.NewArticleEventHandlerKafka(ctx)
	// 创建gRPC服务。
	ctx.gRPCServer = rpc.NewServer(ctx)
	return ctx
}

// GetGRPCServer 获取gRPC服务。
func (c *ContextImpl) GetGRPCServer() *rpc.Server {
	return c.gRPCServer
}

// GetArticleRepository 获取文章数据存储接口
func (c *ContextImpl) GetArticleRepository() interfaces.ArticleRepository {
	return c.articleRepository
}

// GetRecommendDB 获取文章数据库DB
func (c *ContextImpl) GetRecommendDB() *gorm.DB {
	return c.recommendDB
}

// GetArticleEventHandler 获取文章事件处理服务
func (c *ContextImpl) GetArticleEventHandler() interfaces.ArticleEventHandler {
	return c.articleEventHandler
}
