package my

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/article"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// ArticleHandler 处理我的文章页面的handler。
type ArticleHandler struct {
	handlers.BaseHandler
	articleService *service.ArticleService
}

// NewArticleHandler 创建ArticleHandler。
func NewArticleHandler(ctx interfaces.Context) *ArticleHandler {
	handler := ArticleHandler{
		BaseHandler: handlers.BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
		articleService: ctx.GetArticleService(),
	}
	return &handler
}

// Get 处理显示文章列表的页面。
func (h *ArticleHandler) Get(c *fiber.Ctx) error {
	// 获取用户信息。
	userInfo, err := h.GetUserInfo(c)
	if err != nil || !userInfo.IsLoggedIn {
		return h.RenderGenericError(c, fiber.StatusForbidden, fiber.Map{
			"error_code":    fiber.StatusForbidden,
			"error_title":   "无法执行操作",
			"action_prompt": "请确保已经登录！如果已登录，请重试！",
			"error_detail":  "鉴权失败",
		})
	}

	// 获取文章列表。
	// TODO 此处应该考虑分页。
	req := &pb.GetByUserRequest{
		UserId: userInfo.UserID,
	}
	rsp, err := h.articleService.GetByUser(c.UserContext(), req)
	if err != nil {
		log.Printf("failed to get articles for user %d, %v", userInfo.UserID, err)
		return h.RenderError500(c)
	}

	// 渲染页面。
	return c.Render("my_article_list", fiber.Map{
		"title":     "我的文章",
		"articles":  rsp.Articles,
		"user_info": userInfo,
	})
}
