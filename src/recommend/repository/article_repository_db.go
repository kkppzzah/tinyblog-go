// Package repository 数据存储接口。
package repository

import (
	"log"

	"gorm.io/gorm"
	"ppzzl.com/tinyblog-go/recommend/common"
	"ppzzl.com/tinyblog-go/recommend/interfaces"
	"ppzzl.com/tinyblog-go/recommend/model"
)

// ArticleRepositoryDB 用来操作文章存储的接口（存储在内存中）。
type ArticleRepositoryDB struct {
	articleDB *gorm.DB
}

// NewArticleRepositoryDB 创建文章存储接口。
func NewArticleRepositoryDB(ctx interfaces.Context) *ArticleRepositoryDB {
	r := &ArticleRepositoryDB{
		articleDB: ctx.GetRecommendDB(),
	}
	return r
}

// Create 创建文章。
func (r *ArticleRepositoryDB) Create(article *model.Article) (*model.Article, error) {
	if err := r.articleDB.Create(article).Error; err != nil {
		log.Printf("failed to save article to db, %v", err)
		return article, common.NewError(common.ErrorCodeNoFound, err)
	}
	return article, nil
}

// Update 更新文章。
func (r *ArticleRepositoryDB) Update(article *model.Article) error {
	// TODO 更新文章。
	return nil
}

// Delete 删除文章。
func (r *ArticleRepositoryDB) Delete(id int64) error {
	// TODO 删除文章。
	return nil
}

var getArticlesRandomlySQL = "select * from article order by rand() limit ?"

// GetRandomly 随机获取指定数量的文章。
func (r *ArticleRepositoryDB) GetRandomly(num int) ([]*model.Article, error) {
	var articles []*model.Article
	if err := r.articleDB.Raw(getArticlesRandomlySQL, num).Scan(&articles).Error; err != nil {
		log.Printf("failed to get articles randomly, %v", err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return articles, nil
}
