// Package model 数据模型。
package model

import "time"

// Article 存储在搜索引擎中的文档。
type Article struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	PublishTime time.Time `json:"publish_time"`
}
