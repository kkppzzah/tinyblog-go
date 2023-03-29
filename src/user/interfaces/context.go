// Package interfaces 定义各个接口。
package interfaces

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Context 用来管理应用的各个组件。
type Context interface {
	GetUserRepository() UserRepository         // 获取用户数据存储接口
	GetSessionRepository() SessionRepository   // 获取会话数据存储接口
	GetUserDB() *gorm.DB                       // 获取用户数据库DB
	GetRedisClient() *redis.Client             // 获取Redis客户端
	GetUserEventPublisher() UserEventPublisher // 返回用户事件发布服务
}
