// Package app 组织应用的各个组件。
package app

import (
	"log"
	"strings"

	"ppzzl.com/tinyblog-go/search/common"
	"ppzzl.com/tinyblog-go/search/interfaces"
	"ppzzl.com/tinyblog-go/search/rpc"
	"ppzzl.com/tinyblog-go/search/service"

	"github.com/elastic/go-elasticsearch/v8"
)

// ContextImpl 存放各个应用组件。
type ContextImpl struct {
	gRPCServer          *rpc.Server
	articleEventHandler interfaces.ArticleEventHandler
	esClient            *elasticsearch.Client
	articleService      interfaces.ArticleService
}

// NewContextImpl 创建Context实例。
func NewContextImpl() *ContextImpl {
	ctx := &ContextImpl{}
	// 创建各个服务。
	addresses := strings.Split(common.MustGetEnv(common.EnvEsClusterAddresses, ""), ",")
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: addresses,
	})
	if err != nil {
		log.Fatalf("failed to create elasticsearch client, %v", err)
	}
	ctx.esClient = esClient
	ctx.articleService = service.NewArticleServiceImpl(false)
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

// GetArticleEventHandler 获取文章事件处理服务
func (c *ContextImpl) GetArticleEventHandler() interfaces.ArticleEventHandler {
	return c.articleEventHandler
}

// GetEsClient 获取Elasticsearch client。
func (c *ContextImpl) GetEsClient() *elasticsearch.Client {
	return c.esClient
}

// GetArticleService 获取文章服务
func (c *ContextImpl) GetArticleService() interfaces.ArticleService {
	return c.articleService
}
