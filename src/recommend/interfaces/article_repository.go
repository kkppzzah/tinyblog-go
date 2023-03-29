// Package interfaces 定义各个接口。
package interfaces

import "ppzzl.com/tinyblog-go/recommend/model"

// ArticleRepository 用来操作文章存储的接口。
type ArticleRepository interface {
	Create(article *model.Article) (*model.Article, error)
	Update(*model.Article) error
	Delete(id int64) error
	GetRandomly(num int) ([]*model.Article, error)
}
