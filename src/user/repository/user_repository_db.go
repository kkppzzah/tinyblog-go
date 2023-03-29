// Package repository 数据存储接口。
package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"ppzzl.com/tinyblog-go/user/common"
	"ppzzl.com/tinyblog-go/user/interfaces"
	"ppzzl.com/tinyblog-go/user/model"
)

// UserRepositoryDB 用来操作用户存储的接口（存储在数据库中）。
type UserRepositoryDB struct {
	userDB *gorm.DB
}

// NewUserRepositoryDB 创建用户存储接口。
func NewUserRepositoryDB(ctx interfaces.Context) *UserRepositoryDB {
	r := &UserRepositoryDB{
		userDB: ctx.GetUserDB(),
	}
	return r
}

// Create 创建用户。
func (r *UserRepositoryDB) Create(user *model.User) (*model.User, error) {
	if err := r.userDB.Save(user).Error; err != nil {
		log.Printf("failed to save user to db, user name: %s, %v", user.Name, err)
		return user, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return user, nil
}

// Update 更新用户。
func (r *UserRepositoryDB) Update(user *model.User) error {
	if err := r.userDB.Save(user).Error; err != nil {
		log.Printf("failed to update user, user name: %s, %v", user.Name, err)
		return common.NewError(common.ErrorCodeDBOpError, err)
	}
	return nil
}

// Delete 删除用户。
func (r *UserRepositoryDB) Delete(id int64) error {
	if err := r.userDB.Delete(&model.User{}, id).Error; err != nil {
		log.Printf("failed to delete user, user id: %d, %v", id, err)
		return common.NewError(common.ErrorCodeDBOpError, err)
	}
	return nil
}

// Get 获取用户。
func (r *UserRepositoryDB) Get(id int64) (*model.User, error) {
	user := &model.User{
		ID: id,
	}
	if err := r.userDB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewError(common.ErrorCodeNoFound, err)
		}
		log.Printf("failed to get user, user id: %d, %v", id, err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return user, nil
}

// GetByName 通过用户名获取用户。
func (r *UserRepositoryDB) GetByName(name string) (*model.User, error) {
	user := &model.User{}
	if err := r.userDB.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewError(common.ErrorCodeNoFound, err)
		}
		log.Printf("failed to get user, user name: %s, %v", name, err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return user, nil
}
