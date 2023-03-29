// Package repository 数据存储接口。
package repository

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
	"ppzzl.com/tinyblog-go/article/common"
	"ppzzl.com/tinyblog-go/article/interfaces"
	"ppzzl.com/tinyblog-go/article/model"
)

// ArticleRepositoryDB 用来操作文章存储的接口（存储在内存中）。
type ArticleRepositoryDB struct {
	articleDB *gorm.DB
}

// NewArticleRepositoryDB 创建文章存储接口。
func NewArticleRepositoryDB(ctx interfaces.Context) *ArticleRepositoryDB {
	r := &ArticleRepositoryDB{
		articleDB: ctx.GetArticleDB(),
	}
	return r
}

// Create 创建文章。
func (r *ArticleRepositoryDB) Create(article *model.Article) (*model.Article, error) {
	articleMeta := &model.ArticleMeta{
		UserID:  article.UserID,
		Title:   article.Title,
		Tags:    article.Tags,
		Summary: article.Summary,
		ArticleContent: model.ArticleContent{
			Content: article.Content,
		},
	}
	if err := r.articleDB.Create(articleMeta).Error; err != nil {
		log.Printf("failed to save article to db, %v", err)
		return article, common.NewError(common.ErrorCodeNoFound, err)
	}
	article.ID = articleMeta.ID
	return article, nil
}

// Update 更新文章。
func (r *ArticleRepositoryDB) Update(article *model.Article) error {
	err := r.articleDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.ArticleMeta{}).Where("id = ?", article.ID).Updates(map[string]interface{}{
			"title":      article.Title,
			"tags":       article.Tags,
			"summary":    article.Summary,
			"UpdateTime": time.Now(),
		}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.ArticleContent{}).Where("id = ?", article.ID).Update("Content", article.Content).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("failed update article, article id: %d, %v", article.ID, err)
		return common.NewError(common.ErrorCodeDBOpError, err)
	}
	return err
}

// Delete 删除文章。
func (r *ArticleRepositoryDB) Delete(id int64) error {
	if err := r.articleDB.Delete(&model.ArticleMeta{}, id).Error; err != nil {
		log.Printf("failed to delete article, article id: %d, %v", id, err)
		return common.NewError(common.ErrorCodeDBOpError, err)
	}
	return nil
}

// GetArticleMeta 获取文章元数据。
func (r *ArticleRepositoryDB) GetArticleMeta(id int64) (*model.ArticleMeta, error) {
	articleMeta := &model.ArticleMeta{
		ID: id,
	}
	if err := r.articleDB.First(articleMeta, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewError(common.ErrorCodeNoFound, err)
		}
		log.Printf("failed to get article, article id: %d, %v", id, err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return articleMeta, nil
}

var querySingleArticleSQL = "select am.*, u.nickname, ac.content " +
	"from article_meta am join article_content ac on am.id = ac.id " +
	"left join user u on am.user_id = u.id " +
	"where am.id = ?"

// Get 获取文章元数据及文章内容。
func (r *ArticleRepositoryDB) Get(id int64) (*model.Article, error) {
	article := &model.Article{}

	if err := r.articleDB.Raw(querySingleArticleSQL, id).Scan(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewError(common.ErrorCodeNoFound, err)
		}
		log.Printf("failed to get article, article id: %d, %v", id, err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return article, nil
}

var querySingleUsersArticleSQL = "select am.*, u.nickname, ac.content " +
	"from article_meta am join article_content ac on am.id = ac.id " +
	"left join user u on am.user_id = u.id " +
	"where am.user_id = ? " +
	"order by am.publish_time desc"

// GetByUser 获取用户的文章（不含文章内容）。
func (r *ArticleRepositoryDB) GetByUser(userID int64) ([]*model.Article, error) {
	var articles []*model.Article
	if err := r.articleDB.Raw(querySingleUsersArticleSQL, userID).Scan(&articles).Error; err != nil {
		log.Printf("failed to get articles, user id: %d, %v", userID, err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return articles, nil
}

var queryArticlesByIDSQL = "select am.*, u.nickname " +
	"from article_meta am left join user u on am.user_id = u.id " +
	"where am.id in ? "

// GetByIds 获取一组文章ID对应的文章（不含文章内容）。
func (r *ArticleRepositoryDB) GetByIds(ids []int64) ([]*model.Article, error) {
	var articles []*model.Article
	if err := r.articleDB.Raw(queryArticlesByIDSQL, ids).Scan(&articles).Error; err != nil {
		log.Printf("failed to get articles, ids: %d, %v", ids, err)
		return nil, common.NewError(common.ErrorCodeDBOpError, err)
	}
	return articles, nil
}
