package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/search"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
)

// SearchHandler 处理搜索页面的handler。
type SearchHandler struct {
	BaseHandler
	searchService  *service.SearchService
	articleService *service.ArticleService
}

// NewSearchHandler 创建处理搜索页面的handler。
func NewSearchHandler(ctx interfaces.Context) *SearchHandler {
	handler := SearchHandler{
		BaseHandler: BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
		searchService:  ctx.GetSearchService(),
		articleService: ctx.GetArticleService(),
	}
	return &handler
}

// Get 处理http method为GET的请求。
func (h *SearchHandler) Get(c *fiber.Ctx) error {
	// 获取用户信息。
	userInfo, _ := h.GetUserInfo(c)
	content := c.Query("content")

	if content == "" {
		return h.RenderBadRequestPage(c, "搜索内容不可为空！")
	}

	// 执行搜索。
	searchRsp, err := h.searchService.SimpleSearch(c.UserContext(), &pb.SimpleSearchRequest{
		Content: content,
	})
	if err != nil {
		log.Printf("failed to search articles, content: %s, %v", content, err)
		return h.RenderError500(c)
	}

	// 获取文章详情（当前搜索服务没有返回完整的信息）。
	ids := make([]int64, len(searchRsp.Articles))
	for i, article := range searchRsp.Articles {
		ids[i] = article.Id
	}
	rsp, err := h.articleService.GetByIds(c.UserContext(), ids)
	if err != nil {
		log.Printf("failed to get article details, content: %s, %v", content, err)
		return h.RenderError500(c)
	}

	// 渲染页面。
	return c.Render("search", fiber.Map{
		"title":     "搜索",
		"articles":  rsp.Articles,
		"user_info": userInfo,
	})
}
