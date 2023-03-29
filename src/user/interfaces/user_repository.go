// Package interfaces 定义各个接口。
package interfaces

import "ppzzl.com/tinyblog-go/user/model"

// UserRepository 用来操作文章存储的接口。
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Get(id int64) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Update(*model.User) error
	Delete(id int64) error
}
