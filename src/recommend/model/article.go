// Package model 数据模型。
package model

// Article 文章元数据（不含文章内容）。
type Article struct {
	ID int64 `gorm:"column:id;primaryKey;<-:create;not null;type:bigint;autoIncrement"`
	// UserID int64 `gorm:"column:user_id;<-:create;not null;type:bigint"`
}

// TableName 指定ArticleMeta对应的数据库表名。
func (Article) TableName() string {
	return "article"
}
