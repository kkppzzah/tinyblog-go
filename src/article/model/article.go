// Package model 数据模型。
package model

import (
	"time"
)

// ArticleMeta 文章元数据（不含文章内容）。
type ArticleMeta struct {
	ID             int64          `gorm:"column:id;primaryKey;<-:false;not null;type:bigint;autoIncrement"`
	UserID         int64          `gorm:"column:user_id;<-:create;not null;type:bigint"`
	Title          string         `gorm:"column:title;not null;type:varchar;size:120"`
	Tags           string         `gorm:"column:tags;not null;type:varchar;size:40"`
	Summary        string         `gorm:"column:summary;not null;type:varchar;size:200"`
	PublishTime    time.Time      `gorm:"column:publish_time;autoUpdateTime;<-:create;not null"`
	UpdateTime     time.Time      `gorm:"column:update_time;autoUpdateTime;not null"`
	Version        int            `gorm:"column:version;not null;default:0"`
	ArticleContent ArticleContent `gorm:"foreignKey:ID;references:ID"`
}

// TableName 指定ArticleMeta对应的数据库表名。
func (ArticleMeta) TableName() string {
	return "article_meta"
}

// ArticleContent 文章内容。
type ArticleContent struct {
	ID      int64  `gorm:"column:id;primaryKey;<-:create;not null;type:bigint"`
	Content string `gorm:"column:content;not null;type:varchar;size:10000"`
}

// TableName 指定ArticleMeta对应的数据库表名。
func (ArticleContent) TableName() string {
	return "article_content"
}

// Article 文章，这是ArticleMeta和ArticleContent的合集。
type Article struct {
	ID          int64     `gorm:"column:id;primaryKey;<-:false;not null;type:bigint;autoIncrement"`
	UserID      int64     `gorm:"column:user_id;<-:create;not null;type:bigint"`
	Title       string    `gorm:"column:title;not null;type:varchar;size:120"`
	Tags        string    `gorm:"column:tags;not null;type:varchar;size:40"`
	Summary     string    `gorm:"column:summary;not null;type:varchar;size:200"`
	PublishTime time.Time `gorm:"column:publish_time;autoUpdateTime;<-:create;not null"`
	Content     string
	Nickname    string `gorm:"column:nickname;not null;type:varchar;size:120"`
}

// TableName 指定Article对应的数据库表名。
func (Article) TableName() string {
	return "article_meta"
}
