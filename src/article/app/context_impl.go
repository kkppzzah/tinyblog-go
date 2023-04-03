// Package app 组织应用的各个组件。
package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ppzzl.com/tinyblog-go/article/common"
	"ppzzl.com/tinyblog-go/article/interfaces"
	"ppzzl.com/tinyblog-go/article/repository"
	"ppzzl.com/tinyblog-go/article/rpc"
	"ppzzl.com/tinyblog-go/article/service"
)

// ContextImpl 存放各个应用组件。
type ContextImpl struct {
	gRPCServer            *rpc.Server
	articleRepository     interfaces.ArticleRepository
	articleDB             *gorm.DB
	userRepository        interfaces.UserRepository
	userEventHandler      interfaces.UserEventHandler
	articleEventPublisher interfaces.ArticleEventPublisher
}

// NewContextImpl 创建Context实例。
func NewContextImpl() *ContextImpl {
	ctx := &ContextImpl{}
	// 创建各个基础组件。
	articleDB, err := gorm.Open(mysql.Open(common.MustLoadSecretAsString(common.EnvArticleDBConnStr, common.EnvArticleDBConnStrSecretFile)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database (article), %v", err))
	}
	ctx.articleDB = articleDB
	// 数据存储接口。
	ctx.articleRepository = repository.NewArticleRepositoryDB(ctx)
	ctx.userRepository = repository.NewUserRepositoryDB(ctx)
	// 时间发布。
	ctx.articleEventPublisher = service.NewArticleEventPublisherKafka()
	// 事件处理。
	ctx.userEventHandler = service.NewUserEventHandlerKafka(ctx)
	// 创建gRPC服务。
	ctx.gRPCServer = rpc.NewServer(ctx)
	return ctx
}

// GetGRPCServer 获取gRPC服务。
func (c *ContextImpl) GetGRPCServer() *rpc.Server {
	return c.gRPCServer
}

// GetArticleRepository 获取文章数据存储接口。
func (c *ContextImpl) GetArticleRepository() interfaces.ArticleRepository {
	return c.articleRepository
}

// GetArticleDB 获取文章数据库DB。
func (c *ContextImpl) GetArticleDB() *gorm.DB {
	return c.articleDB
}

// GetUserRepository 获取用户数据存储接口。
func (c *ContextImpl) GetUserRepository() interfaces.UserRepository {
	return c.userRepository
}

// GetArticleEventPublisher 获取文章事件发布接口
func (c *ContextImpl) GetArticleEventPublisher() interfaces.ArticleEventPublisher {
	return c.articleEventPublisher
}
