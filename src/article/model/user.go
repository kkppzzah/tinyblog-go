// Package model 数据模型。
package model

// User 用户信息。
type User struct {
	ID       int64  `gorm:"column:id;primaryKey;<-:create;not null;type:bigint;autoIncrement"`
	Name     string `gorm:"column:name;<-:create;not null;type:varchar;size:20;unique"`
	Nickname string `gorm:"column:nickname;not null;type:varchar;size:32"`
}

// TableName 指定User对应的数据库表名。
func (User) TableName() string {
	return "user"
}
