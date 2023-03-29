// Package repository 数据存储接口。
package repository

import (
	"ppzzl.com/tinyblog-go/user/common"
	"ppzzl.com/tinyblog-go/user/model"
)

// UserRepositoryMemory 用来操作用户存储的接口（存储在内存中）。
type UserRepositoryMemory struct {
	id    int64
	store map[int64]*model.User
}

// NewUserRepositoryMemory 创建用户存储接口。
func NewUserRepositoryMemory() *UserRepositoryMemory {
	r := &UserRepositoryMemory{
		store: make(map[int64]*model.User),
	}
	return r
}

// Create 创建用户。
func (r *UserRepositoryMemory) Create(user *model.User) (*model.User, error) {
	r.id++
	user.ID = r.id
	r.store[r.id] = user
	return user, nil
}

// Update 更新用户。
func (r *UserRepositoryMemory) Update(user *model.User) error {
	u := r.store[user.ID]
	if u == nil {
		return common.NewError(common.ErrorCodeNoFound, nil)
	}
	u.Name = user.Name
	u.Nickname = user.Nickname
	u.Avatar = user.Avatar
	u.Bio = user.Bio
	return nil
}

// Delete 删除用户。
func (r *UserRepositoryMemory) Delete(id int64) error {
	delete(r.store, id)
	return nil
}

// Get 获取用户。
func (r *UserRepositoryMemory) Get(id int64) (*model.User, error) {
	return r.store[id], nil
}

// GetByName 通过用户名获取用户。
func (r *UserRepositoryMemory) GetByName(name string) (*model.User, error) {
	for _, v := range r.store {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, common.NewError(common.ErrorCodeNoFound, nil)
}
