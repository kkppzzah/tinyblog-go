// Package app 组织应用的各个组件。
package app

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ppzzl.com/tinyblog-go/user/common"
	"ppzzl.com/tinyblog-go/user/interfaces"
	"ppzzl.com/tinyblog-go/user/repository"
	"ppzzl.com/tinyblog-go/user/rpc"
	"ppzzl.com/tinyblog-go/user/service"
)

// ContextImpl 存放各个应用组件。
type ContextImpl struct {
	gRPCServer         *rpc.Server
	redisClient        *redis.Client
	userDB             *gorm.DB
	userRepository     interfaces.UserRepository
	sessionRepository  interfaces.SessionRepository
	userEventPublisher interfaces.UserEventPublisher
}

// NewContextImpl 创建Context实例。
func NewContextImpl() *ContextImpl {
	ctx := &ContextImpl{}
	// 创建各个基础组件。
	ctx.redisClient = redis.NewClient(&redis.Options{
		Addr: common.MustGetEnv(common.EnvRedisConnStr, ""),
	})
	userDB, err := gorm.Open(mysql.Open(common.MustLoadSecretAsString(common.EnvUserDBConnStr, common.EnvUserDBConnStrSecretFile)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database (user), %v", err))
	}
	ctx.userDB = userDB
	// 数据存储接口。
	ctx.userRepository = repository.NewUserRepositoryDB(ctx)
	ctx.sessionRepository = repository.NewSessionRepositoryRedis(ctx)
	// 用户事件通知。
	ctx.userEventPublisher = service.NewUserEventPublisherKafka()
	// 创建gRPC服务。
	ctx.gRPCServer = rpc.NewServer(ctx)
	return ctx
}

// GetGRPCServer 获取gRPC服务。
func (c *ContextImpl) GetGRPCServer() *rpc.Server {
	return c.gRPCServer
}

// GetUserRepository 获取文章数据存储接口。
func (c *ContextImpl) GetUserRepository() interfaces.UserRepository {
	return c.userRepository
}

// GetSessionRepository 获取会话数据存储接口。
func (c *ContextImpl) GetSessionRepository() interfaces.SessionRepository {
	return c.sessionRepository
}

// GetUserDB 获取用户数据库DB
func (c *ContextImpl) GetUserDB() *gorm.DB {
	return c.userDB
}

// GetRedisClient 获取Redis客户端。
func (c *ContextImpl) GetRedisClient() *redis.Client {
	return c.redisClient
}

// GetUserEventPublisher 返回用户事件发布服务
func (c *ContextImpl) GetUserEventPublisher() interfaces.UserEventPublisher {
	return c.userEventPublisher
}
