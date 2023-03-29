// Package interfaces 定义各个接口。
package interfaces

import "github.com/elastic/go-elasticsearch/v8"

// Context 用来管理应用的各个组件。
type Context interface {
	GetArticleEventHandler() ArticleEventHandler // 获取文章事件处理服务
	GetEsClient() *elasticsearch.Client          // 获取Elasticsearch client
	GetArticleService() ArticleService           // 获取文章服务
}
