// Package interfaces 定义各个接口。
package interfaces

import "ppzzl.com/tinyblog-go/article/model"

// ArticleRepository 用来操作文章存储的接口。
type ArticleRepository interface {
	Create(article *model.Article) (*model.Article, error)
	GetArticleMeta(id int64) (*model.ArticleMeta, error)
	Get(id int64) (*model.Article, error)
	Update(*model.Article) error
	Delete(id int64) error
	GetByUser(userID int64) ([]*model.Article, error)
	GetByIds(ids []int64) ([]*model.Article, error)
}
