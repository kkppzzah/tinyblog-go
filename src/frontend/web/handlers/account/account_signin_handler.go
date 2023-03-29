// Package account 处理账户相关请求。
package account

import (
	"log"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/auth"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// SigninHandler 处理登录页面的handler。
type SigninHandler struct {
	handlers.BaseHandler
	authService *service.AuthService
}

// NewSigninHandler 创建处理登录页面的handler。
func NewSigninHandler(ctx interfaces.Context) *SigninHandler {
	handler := SigninHandler{
		authService: ctx.GetAuthService(),
	}
	return &handler
}

// Get 处理http method为GET的请求。
func (h *SigninHandler) Get(c *fiber.Ctx) error {
	// 渲染页面。
	return c.Render("account_signin", fiber.Map{
		"title": "登录",
	})
}

// SignInRequest 登录请求的参数。
type SignInRequest struct {
	Name     string `validate:"required,min=2" form:"name"`
	Password string `validate:"required,min=6" form:"password"`
}

// Post 处理登录请求。
func (h *SigninHandler) Post(c *fiber.Ctx) error {
	request := &SignInRequest{}
	if err := c.BodyParser(request); err != nil {
		return h.RenderBadRequestPage(c, err.Error())
	}

	if err := validate.Struct(request); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			log.Printf("failed to sign in, failed to validate request, %v/%s", err, reflect.TypeOf(err))
			return h.RenderError500(c)
		}
		return h.RenderBadRequestPage(c, h.GetValidationErrorDetails(err))
	}

	authRequest := &pb.SignInRequest{
		Name:       request.Name,
		Password:   request.Password,
		AuthMethod: pb.AuthMethod_AUTH_METHOD_SESSION,
	}
	authRespone, err := h.authService.SignIn(c.UserContext(), authRequest)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return h.RenderError404(c)
		}
		log.Printf("failed to sign in, user: %s, %v", request.Name, err)
		return h.RenderError500(c)
	}
	// 设置cookie
	cookie := &fiber.Cookie{
		Name:    common.CookieNameSession,
		Value:   authRespone.AuthInfo.Session.Session,
		Expires: time.Now().Add(120 * 24 * time.Hour),
	}
	c.Cookie(cookie)

	return c.Render("show_result_redirect", fiber.Map{
		"result_title":  "登录成功",
		"result_detail": "已登陆，如果未自动跳转，请点击右侧链接！",
		"redirect_url":  "/",
	})
}
