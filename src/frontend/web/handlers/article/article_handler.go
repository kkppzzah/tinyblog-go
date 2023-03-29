// Package article 处理文章相关请求。
package article

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/article"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// Handler 处理登录页面的handler。
type Handler struct {
	handlers.BaseHandler
	articleService *service.ArticleService
	validate       *validator.Validate
}

// NewHandler 创建处理登录页面的handler。
func NewHandler(ctx interfaces.Context) *Handler {
	handler := Handler{
		BaseHandler: handlers.BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
		articleService: ctx.GetArticleService(),
		validate:       validator.New(),
	}
	return &handler
}

// Get 处理获取显示文章页面的请求。
func (h *Handler) Get(c *fiber.Ctx) error {
	userInfo, err := h.GetUserInfo(c)
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
		log.Printf("failed to request article service (GET), %v", err)
		return h.RenderError500(c)
	}

	// 渲染页面。
	return c.Render("article", fiber.Map{
		"title":                "查看文章",
		"article_title":        getResponse.Title,
		"article_username":     getResponse.Nickname,
		"article_publish_time": getResponse.PublishTime,
		"article_content":      getResponse.Content,
		"is_article_owner":     userInfo.UserID == getResponse.UserId,
		"update_article_url":   fmt.Sprintf("/article/%d/edit", id),
		"user_info":            userInfo,
	})
}

// UpdateRequest 更新文章请求的参数。
type UpdateRequest struct {
	Title   string `validate:"required,min=2" form:"title"`
	Summary string `validate:"required,min=5" form:"summary"`
	Content string `validate:"required,min=10" form:"content"`
}

// Post 处理编辑文章请求。
func (h *Handler) Post(c *fiber.Ctx) error {
	userInfo, err := h.GetUserInfo(c)
	fmt.Printf("%v-->", userInfo)
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return h.RenderBadRequestPage(c, fmt.Sprintf("id参数当前值为：%s，值不合法！正确格式：数字。", idParam))
	}

	request := new(UpdateRequest)

	if err := c.BodyParser(request); err != nil {
		return h.RenderBadRequestPage(c, err.Error())
	}

	if err := h.validate.Struct(request); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			return h.RenderError500(c)
		}
		return h.RenderBadRequestPage(c, h.GetValidationErrorDetails(err))
	}

	updateRequest := &pb.UpdateRequest{
		Id:      id,
		Title:   request.Title,
		Summary: request.Summary,
		Content: request.Content,
	}
	_, err = h.articleService.Update(c.UserContext(), updateRequest)
	if err != nil {
		log.Printf("failed to request article service (UPDATE), %v", err)
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return h.RenderError404(c)
		}
		return h.RenderError500(c)
	}

	return c.Render("show_result_redirect", fiber.Map{
		"result_title":  "更新成功",
		"result_detail": "文章已更新完成，如果未自动跳转，请点击右侧链接！",
		"redirect_url":  fmt.Sprintf("/article/%d", id),
		"user_info":     userInfo,
	})
}

// Delete 处理删除文章请求。
func (h *Handler) Delete(c *fiber.Ctx) error {
	userInfo, err := h.GetUserInfo(c)
	fmt.Printf("%v-->", userInfo)
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return h.RenderBadRequestPage(c, fmt.Sprintf("id参数当前值为：%s，值不合法！正确格式：数字。", idParam))
	}

	deleteRequest := &pb.DeleteRequest{
		Id: id,
	}
	_, err = h.articleService.Delete(c.UserContext(), deleteRequest)
	if err != nil {
		log.Printf("failed to request article service, %v", err)
		return h.RenderError500(c)
	}

	return c.Render("show_result_redirect", fiber.Map{
		"result_title":  "删除成功",
		"result_detail": "文章已删除完成，如果未自动跳转，请点击右侧链接！",
		"redirect_url":  fmt.Sprintf("/user/%d/article", userInfo.UserID),
	})
}
