// Package interfaces 定义各个接口。
package interfaces

import "gorm.io/gorm"

// Context 用来管理应用的各个组件。
type Context interface {
	GetArticleRepository() ArticleRepository         // 获取文章数据存储接口
	GetArticleDB() *gorm.DB                          // 获取文章数据库DB
	GetUserRepository() UserRepository               // 获取用户数据存储接口
	GetArticleEventPublisher() ArticleEventPublisher // 获取文章事件发布接口
}
