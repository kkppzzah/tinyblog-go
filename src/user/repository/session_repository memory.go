// Package repository 数据存储接口。
package repository

import "ppzzl.com/tinyblog-go/user/common"

// SessionRepositoryMemory 用来操作会话存储的接口。
type SessionRepositoryMemory struct {
	store map[string]int64
}

// NewSessionRepositoryMemory 创建SessionRepositoryMemory实例。
func NewSessionRepositoryMemory() *SessionRepositoryMemory {
	return &SessionRepositoryMemory{
		store: make(map[string]int64),
	}
}

// Create 创建会话记录。
func (r *SessionRepositoryMemory) Create(sessionID string, userID int64) error {
	r.store[sessionID] = userID
	return nil
}

// Get 获取会话记录。
func (r *SessionRepositoryMemory) Get(sessionID string) (int64, error) {
	userID := r.store[sessionID]
	if userID == 0 {
		return userID, common.NewError(common.ErrorCodeNoFound, nil)
	}
	return userID, nil
}

// Delete 删除会话记录。
func (r *SessionRepositoryMemory) Delete(sessionID string) error {
	delete(r.store, sessionID)
	return nil
}
