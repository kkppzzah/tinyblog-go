// Package repository 数据存储接口。
package repository

import (
	"sort"

	"ppzzl.com/tinyblog-go/article/common"
	"ppzzl.com/tinyblog-go/article/model"
)

// ArticleRepositoryMemory 用来操作文章存储的接口（存储在内存中）。
type ArticleRepositoryMemory struct {
	id    int64
	store map[int64]*model.Article
}

// NewArticleRepositoryMemory 创建文章存储接口。
func NewArticleRepositoryMemory() *ArticleRepositoryMemory {
	r := &ArticleRepositoryMemory{
		store: make(map[int64]*model.Article),
	}
	return r
}

// Create 创建文章。
func (r *ArticleRepositoryMemory) Create(article *model.Article) (*model.Article, error) {
	r.id++
	article.ID = r.id
	r.store[r.id] = article
	return article, nil
}

// Update 更新文章。
func (r *ArticleRepositoryMemory) Update(newArticle *model.Article) error {
	article := r.store[newArticle.ID]
	if article == nil {
		return common.NewError(common.ErrorCodeNoFound, nil)
	}
	article.Title = newArticle.Title
	article.Summary = newArticle.Summary
	article.Content = newArticle.Content
	return nil
}

// Delete 删除文章。
func (r *ArticleRepositoryMemory) Delete(id int64) error {
	delete(r.store, id)
	return nil
}

// GetArticleMeta 获取文章元数据。
func (r *ArticleRepositoryMemory) GetArticleMeta(id int64) (*model.ArticleMeta, error) {
	if article := r.store[id]; article != nil {
		return &model.ArticleMeta{
			ID:          article.ID,
			UserID:      article.UserID,
			Title:       article.Title,
			Tags:        article.Tags,
			Summary:     article.Summary,
			PublishTime: article.PublishTime,
		}, nil
	}
	return nil, nil
}

// Get 获取文章元数据及文章内容。
func (r *ArticleRepositoryMemory) Get(id int64) (*model.Article, error) {
	return r.store[id], nil
}

// GetByUser 获取用户的文章（不含文章内容）。
func (r *ArticleRepositoryMemory) GetByUser(userID int64) ([]*model.Article, error) {
	articles := []*model.Article{}
	for _, article := range r.store {
		if article.UserID == userID {
			articles = append(articles, article)
		}
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishTime.After(articles[j].PublishTime)
	})
	return articles, nil
}

// GetByIds 获取一组文章ID对应的文章（不含文章内容）。
func (r *ArticleRepositoryMemory) GetByIds(ids []int64) ([]*model.Article, error) {
	articles := []*model.Article{}
	for _, articleID := range ids {
		if article := r.store[articleID]; article != nil {
			articles = append(articles, article)
		}
	}
	return articles, nil
}
