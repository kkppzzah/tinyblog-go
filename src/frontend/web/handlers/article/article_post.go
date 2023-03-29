package article

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/article"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// UserArticlePostHandler 处理登录页面的handler。
type UserArticlePostHandler struct {
	handlers.BaseHandler
	articleService *service.ArticleService
}

// UserArticlePostRequest 发布文章请求的参数。
type UserArticlePostRequest struct {
	Title   string `validate:"required,min=2" form:"title"`
	Summary string `validate:"required,min=5" form:"summary"`
	Content string `validate:"required,min=10" form:"content"`
}

// NewUserArticlePostHandler 创建处理登录页面的handler。
func NewUserArticlePostHandler(ctx interfaces.Context) *UserArticlePostHandler {
	handler := UserArticlePostHandler{
		BaseHandler: handlers.BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
		articleService: ctx.GetArticleService(),
	}
	return &handler
}

// Get 处理获取发布文章页面的请求。
func (h *UserArticlePostHandler) Get(c *fiber.Ctx) error {
	userInfo, err := h.GetUserInfo(c)
	if err != nil || !userInfo.IsLoggedIn {
		return h.RenderGenericError(c, fiber.StatusForbidden, fiber.Map{
			"error_code":    fiber.StatusBadRequest,
			"error_title":   "无法执行发布",
			"action_prompt": "请确保已经登录！如果已登录，请重试！",
			"error_detail":  "鉴权失败",
		})
	}

	// 渲染页面。
	return c.Render("user_article_post", fiber.Map{
		"title":     "发表文章",
		"user_info": userInfo,
	})
}

// 用作请求校验。
var validate = validator.New()

// Post 处理发布文章请求。
func (h *UserArticlePostHandler) Post(c *fiber.Ctx) error {
	userInfo, err := h.GetUserInfo(c)
	if err != nil || !userInfo.IsLoggedIn {
		return h.RenderGenericError(c, fiber.StatusForbidden, fiber.Map{
			"error_code":    fiber.StatusForbidden,
			"error_title":   "无法执行发布",
			"action_prompt": "请确保已经登录！如果已登录，请重试！",
			"error_detail":  "鉴权失败",
		})
	}

	request := new(UserArticlePostRequest)
	if err := c.BodyParser(request); err != nil {
		return h.RenderBadRequestPage(c, err.Error())
	}

	if err := validate.Struct(request); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			return h.RenderError500(c)
		}
		return h.RenderBadRequestPage(c, h.GetValidationErrorDetails(err))
	}

	pulishRequest := &pb.PublishRequest{
		UserId:  userInfo.UserID,
		Title:   request.Title,
		Summary: request.Summary,
		Content: request.Content,
	}
	publishResponse, err := h.articleService.Publish(c.UserContext(), pulishRequest)
	if err != nil {
		log.Printf("failed to request article service, %v", err)
		return h.RenderError500(c)
	}

	fmt.Println(request.Content)
	return c.Render("show_result_redirect", fiber.Map{
		"result_title":  "发布成功",
		"result_detail": "文章已成功发布，如果未自动跳转，请点击右侧链接！",
		"redirect_url":  fmt.Sprintf("/article/%d", publishResponse.Id),
	})
}
