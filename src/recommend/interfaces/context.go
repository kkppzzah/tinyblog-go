// Package interfaces 定义各个接口。
package interfaces

import "gorm.io/gorm"

// Context 用来管理应用的各个组件。
type Context interface {
	GetArticleRepository() ArticleRepository     // 获取文章数据存储接口
	GetRecommendDB() *gorm.DB                    // 获取文章数据库DB
	GetArticleEventHandler() ArticleEventHandler // 获取文章事件处理服务
}
