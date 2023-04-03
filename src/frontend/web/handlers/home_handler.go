package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pbarticle "ppzzl.com/tinyblog-go/frontend/genproto/article"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
)

// HomeHandler 处理首页的handler。
type HomeHandler struct {
	BaseHandler
	recommentService *service.RecommendService
	articleService   *service.ArticleService
}

// NewHomeHandler 创建处理首页的handler。
func NewHomeHandler(ctx interfaces.Context) *HomeHandler {
	handler := HomeHandler{
		BaseHandler: BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
		recommentService: ctx.GetRecommendService(),
		articleService:   ctx.GetArticleService(),
	}
	return &handler
}

// Get 处理http method为GET的请求。
func (h *HomeHandler) Get(c *fiber.Ctx) error {
	// 获取用户信息。
	userInfo, _ := h.GetUserInfo(c)
	// 获取推荐文章ID。
	recommendArticleIds, err := h.recommentService.RecommendForHome(c.UserContext(), userInfo.UserID)
	log.Printf("articleIds: %v, %v", recommendArticleIds, err)
	// 获取推荐文章列表。
	rsp, err := h.articleService.GetByIds(c.UserContext(), recommendArticleIds)
	var articles []*pbarticle.UserArticle
	if rsp != nil {
		articles = rsp.Articles
	}
	// 渲染页面。
	return c.Render("home", fiber.Map{
		"title":              "首页",
		"recommend_articles": articles,
		"user_info":          userInfo,
	})
}
