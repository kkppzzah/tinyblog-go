package article

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/article"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// EditHandler 处理编辑页面的handler。
type EditHandler struct {
	handlers.BaseHandler
	articleService *service.ArticleService
}

// NewEditHandler 创建编辑页面的handler。
func NewEditHandler(ctx interfaces.Context) *EditHandler {
	handler := EditHandler{
		BaseHandler: handlers.BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
		articleService: ctx.GetArticleService(),
	}
	return &handler
}

// Get 处理获取发布文章页面的请求。
func (h *EditHandler) Get(c *fiber.Ctx) error {
	userInfo, err := h.GetUserInfo(c)
	if err != nil || !userInfo.IsLoggedIn {
		return h.RenderGenericError(c, fiber.StatusForbidden, fiber.Map{
			"error_code":    fiber.StatusForbidden,
			"error_title":   "无法执行编辑",
			"action_prompt": "请确保已经登录！如果已登录，请重试！",
			"error_detail":  "鉴权失败",
		})
	}

	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return h.RenderBadRequestPage(c, fmt.Sprintf("id参数当前值为：%s，值不合法！正确格式：数字。", idParam))
	}

	getResponse, err := h.articleService.Get(c.UserContext(), &pb.GetRequest{
		Id: id,
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return h.RenderError404(c)
		}
		log.Printf("failed to request article service, %v", err)
		return h.RenderError500(c)
	}

	if getResponse.UserId != userInfo.UserID {
		return h.RenderGenericError(c, fiber.StatusForbidden, fiber.Map{
			"error_code":    fiber.StatusForbidden,
			"error_title":   "无法执行编辑",
			"action_prompt": "请确保已经登录！如果已登录，请重试！",
			"error_detail":  "非文章所有者",
		})
	}

	// 渲染页面。
	return c.Render("article_edit", fiber.Map{
		"title":                 "编辑文章",
		"article_title":         base64.StdEncoding.EncodeToString([]byte(getResponse.Title)),
		"article_content":       base64.StdEncoding.EncodeToString([]byte(getResponse.Content)),
		"article_summary":       base64.StdEncoding.EncodeToString([]byte(getResponse.Summary)),
		"atricle_edit_post_url": fmt.Sprintf("/article/%d", id),
		"user_info":             userInfo,
	})
}
