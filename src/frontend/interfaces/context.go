// Package interfaces 定义各个接口。
package interfaces

import (
	"ppzzl.com/tinyblog-go/frontend/service"
)

// Context 用来获取系统各个组件的接口。
type Context interface {
	//
	GetRecommendService() *service.RecommendService
	// GetArticleService 获取推荐服务。
	GetArticleService() *service.ArticleService
	// GetAuthService 获取鉴权服务。
	GetAuthService() *service.AuthService
	// GetSearchService 获取搜索服务。
	GetSearchService() *service.SearchService
	// GetWebService 获取web服务。
	GetWebService() WebService
}
