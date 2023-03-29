// Package model 数据模型。
package model

import (
	"time"
)

// User 用户信息。
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;<-:false;not null;type:bigint;autoIncrement"`
	Name       string    `gorm:"column:name;<-:create;not null;type:varchar;size:20;unique"`
	Password   string    `gorm:"column:password;not null;type:varchar;size:45"`
	Avatar     string    `gorm:"column:avatar;not null;type:varchar;size:255"`
	Nickname   string    `gorm:"column:nickname;not null;type:varchar;size:32"`
	Bio        string    `gorm:"column:bio;not null;type:varchar;size:500"`
	CreateTime time.Time `gorm:"column:create_time;autoUpdateTime;<-:create;not null"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime;not null"`
	Version    int       `gorm:"column:version;not null;default:0"`
}

// TableName 指定User对应的数据库表名。
func (User) TableName() string {
	return "user"
}

// Role 角色。
type Role struct {
	ID   int64  `gorm:"column:id;primaryKey;<-:create;not null;type:bigint;autoIncrement"`
	name string `gorm:"column:name;not null;type:varchar;size:45;unique"`
}

// TableName 指定Role对应的数据库表名。
func (Role) TableName() string {
	return "role"
}

// Permission 权限。
type Permission struct {
	ID   int64  `gorm:"column:id;primaryKey;<-:false;not null;type:bigint;autoIncrement"`
	name string `gorm:"column:name;not null;type:varchar;size:45;unique"`
}

// TableName 指定Permission对应的数据库表名。
func (Permission) TableName() string {
	return "role"
}
