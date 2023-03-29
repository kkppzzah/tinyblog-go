package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/auth"
	"ppzzl.com/tinyblog-go/frontend/service"
)

// GetHandler 处理Get请求
type GetHandler interface {
	Get(c *fiber.Ctx) error
}

// PostHandler 处理Post请求
type PostHandler interface {
	Post(c *fiber.Ctx) error
}

// PutHandler 处理Put请求
type PutHandler interface {
	Put(c *fiber.Ctx) error
}

// DeleteHandler 处理Delete请求
type DeleteHandler interface {
	Delete(c *fiber.Ctx) error
}

// OptionsHandler 处理Options请求
type OptionsHandler interface {
	Options(c *fiber.Ctx) error
}

// PatchHandler 处理Patch请求
type PatchHandler interface {
	Patch(c *fiber.Ctx) error
}

// UserInfo 通过鉴权服务获取到的用户信息。
type UserInfo struct {
	UserID     int64  `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Role       string `json:"role"`
	IsLoggedIn bool   `json:"is_logged_in"`
}

// BaseHandler Handler的通用功能。
type BaseHandler struct {
	AuthService *service.AuthService
}

// GetUserInfo 获取用户信息。
func (h *BaseHandler) GetUserInfo(c *fiber.Ctx) (*UserInfo, error) {
	userInfo := UserInfo{}
	// 如果网关已经获取用户信息，则从头部解析用户信息。
	header := c.Get(common.UserInfoHeader)
	var err error
	if header != "" {
		err = json.Unmarshal([]byte(header), &userInfo)
		userInfo.IsLoggedIn = err == nil
	}
	if !userInfo.IsLoggedIn {
		// 如果网关处没有获取用户信息，则调用鉴权服务获取用户信息。
		sessionCookie := c.Cookies(common.CookieNameSession)
		if sessionCookie != "" && h.AuthService != nil {
			req := &pb.GetAuthUserInfoRequest{
				AuthInfo: &pb.AuthInfo{
					Session: &pb.SessionInfo{
						Session: sessionCookie,
					},
				},
			}
			rsp, err := h.AuthService.GetAuthUserInfo(c.UserContext(), req)
			if err == nil && rsp != nil {
				userInfo.UserID = rsp.Id
				userInfo.Name = rsp.Name
				userInfo.NickName = rsp.Nickname
				userInfo.Avatar = rsp.Avatar
				userInfo.IsLoggedIn = true
			}
		}
	}
	return &userInfo, err
}

// RenderError500 返回500错误页面。
func (h *BaseHandler) RenderError500(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendFile("./static/html/500.html")
}

// RenderError404 返回404错误页面。
func (h *BaseHandler) RenderError404(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendFile("./static/html/404.html")
}

// RenderGenericError 返回通用的错误页面。
func (h *BaseHandler) RenderGenericError(c *fiber.Ctx, status int, params fiber.Map) error {
	return c.Status(status).Render("generic_error_page", params)
}

// GetValidationErrorDetails 获取验证错误的详细信息。
func (h *BaseHandler) GetValidationErrorDetails(err error) string {
	var sb strings.Builder
	// TODO 这地方应该转换为给人看的信息。
	for _, err := range err.(validator.ValidationErrors) {
		if sb.Len() > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strings.ToLower(err.Field()))
		sb.WriteString(":")
		sb.WriteString(err.Tag())
		sb.WriteString("/")
		sb.WriteString(err.Param())
		value := fmt.Sprintf("%v", err.Value())
		if value != "" {
			sb.WriteString(":")
			sb.WriteString(value)
		}
	}
	return sb.String()
}

// RenderBadRequestPage 渲染错误请求应答页面。
func (h *BaseHandler) RenderBadRequestPage(c *fiber.Ctx, errorDetail string) error {
	return h.RenderGenericError(c, fiber.StatusBadRequest, fiber.Map{
		"error_code":    fiber.StatusBadRequest,
		"error_title":   "请求错误",
		"action_prompt": "您可以重新提交请求或返回首页！",
		"error_detail":  errorDetail,
	})
}
