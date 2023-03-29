// Package repository 数据存储接口。
package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"ppzzl.com/tinyblog-go/user/common"
	"ppzzl.com/tinyblog-go/user/interfaces"
)

// SessionRepositoryRedis 用来操作会话存储的接口。
type SessionRepositoryRedis struct {
	redisClient *redis.Client
	ctx         context.Context
}

// NewSessionRepositoryRedis 创建SessionRepositoryRedis实例。
func NewSessionRepositoryRedis(ctx interfaces.Context) *SessionRepositoryRedis {
	return &SessionRepositoryRedis{
		redisClient: ctx.GetRedisClient(),
		ctx:         context.Background(),
	}
}

// Create 创建会话记录。
func (r *SessionRepositoryRedis) Create(sessionID string, userID int64) error {
	sessionKey := r.createSessionKey(sessionID)
	ctx, cancel := context.WithTimeout(r.ctx, 1000*time.Millisecond)
	defer cancel()
	err := r.redisClient.Set(ctx, sessionKey, strconv.FormatInt(userID, 10), 20*24*time.Hour).Err()
	if err != nil {
		log.Printf("failed to set session for user %d, %v", userID, err)
		return err
	}
	return nil
}

// Get 获取会话记录。
func (r *SessionRepositoryRedis) Get(sessionID string) (int64, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 1000*time.Millisecond)
	defer cancel()
	sessionKey := r.createSessionKey(sessionID)
	val, err := r.redisClient.Get(ctx, sessionKey).Result()
	if err == redis.Nil {
		return 0, common.NewError(common.ErrorCodeNoFound, err)
	} else if err != nil {
		log.Printf("failed to get session data, %v", err)
		return 0, common.NewError(common.ErrorCodeDBOpError, err)
	}
	userID, _ := strconv.ParseInt(val, 10, 64)
	return userID, nil
}

// Delete 删除会话记录。
func (r *SessionRepositoryRedis) Delete(sessionID string) error {
	ctx, cancel := context.WithTimeout(r.ctx, 1000*time.Millisecond)
	defer cancel()
	sessionKey := r.createSessionKey(sessionID)
	err := r.redisClient.Del(ctx, sessionKey).Err()
	if err != nil && err != redis.Nil {
		log.Printf("failed to delete session data, %v", err)
		return common.NewError(common.ErrorCodeDBOpError, err)
	}
	return nil
}

func (r *SessionRepositoryRedis) createSessionKey(sessionID string) string {
	return fmt.Sprintf("%s%s", common.RedisSessionKeyPrefix, sessionID)
}
