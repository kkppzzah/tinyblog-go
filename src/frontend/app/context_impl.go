// Package app 组织应用的各个组件。
package app

import (
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web"
)

// ContextImpl 存放各个应用组件。
type ContextImpl struct {
	recommendService *service.RecommendService
	articleService   *service.ArticleService
	authService      *service.AuthService
	searchService    *service.SearchService
	webService       *web.Service
}

// NewContextImpl 创建ContextImpl实例。
func NewContextImpl() *ContextImpl {
	ctx := &ContextImpl{}
	// 创建各个gRPC服务（此处的服务指的是对服务的访问，不是服务的服务端实现）。
	ctx.recommendService = service.NewRecommendService(false)
	ctx.articleService = service.NewArticleService(false)
	ctx.authService = service.NewAuthService(false)
	ctx.searchService = service.NewSearchService(false)
	// 创建Web服务。
	ctx.webService = web.NewWebService(ctx)
	//
	return ctx
}

// GetRecommendService 获取推荐服务。
func (c *ContextImpl) GetRecommendService() *service.RecommendService {
	return c.recommendService
}

// GetArticleService 获取推荐服务。
func (c *ContextImpl) GetArticleService() *service.ArticleService {
	return c.articleService
}

// GetWebService 获取web服务。
func (c *ContextImpl) GetWebService() interfaces.WebService {
	return c.webService
}

// GetAuthService 获取鉴权服务。
func (c *ContextImpl) GetAuthService() *service.AuthService {
	return c.authService
}

// GetSearchService 获取搜索服务。
func (c *ContextImpl) GetSearchService() *service.SearchService {
	return c.searchService
}
