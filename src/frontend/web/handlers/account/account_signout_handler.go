package account

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/auth"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// SignoutHandler 处理退出登录。
type SignoutHandler struct {
	handlers.BaseHandler
}

// NewSignoutHandler 创建处理注册页面的handler。
func NewSignoutHandler(ctx interfaces.Context) *SignoutHandler {
	handler := SignoutHandler{
		BaseHandler: handlers.BaseHandler{
			AuthService: ctx.GetAuthService(),
		},
	}
	return &handler
}

// Post 处理退出登录请求。
func (h *SignoutHandler) Post(c *fiber.Ctx) error {
	sessionCookie := c.Cookies(common.CookieNameSession)
	if sessionCookie == "" {
		return h.RenderBadRequestPage(c, "尚未登录！")
	}

	signOutRequest := &pb.SignOutRequest{
		AuthInfo: &pb.AuthInfo{
			Session: &pb.SessionInfo{
				Session: sessionCookie,
			},
		},
	}
	_, err := h.AuthService.SignOut(c.UserContext(), signOutRequest)
	if err != nil {
		log.Printf("failed to sign out, %v", err)
		return h.RenderError500(c)
	}
	// 设置cookie
	c.ClearCookie(common.CookieNameSession)

	return c.Render("show_result_redirect", fiber.Map{
		"result_title":  "退出登录",
		"result_detail": "已退出登录，如果未自动跳转，请点击右侧链接！",
		"redirect_url":  "/",
	})
}
