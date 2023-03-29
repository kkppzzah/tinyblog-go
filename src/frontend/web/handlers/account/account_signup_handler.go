package account

import (
	"log"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/auth"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/service"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
)

// SignupHandler 处理注册页面的handler。
type SignupHandler struct {
	handlers.BaseHandler
	authService *service.AuthService
}

// NewSignupHandler 创建处理注册页面的handler。
func NewSignupHandler(ctx interfaces.Context) *SignupHandler {
	handler := SignupHandler{
		authService: ctx.GetAuthService(),
	}
	return &handler
}

// Get 向用户显示注册页面。
func (h *SignupHandler) Get(c *fiber.Ctx) error {
	// 渲染页面。
	return c.Render("account_signup", fiber.Map{
		"title": "注册",
	})
}

// SignUpRequest 注册请求的参数。
type SignUpRequest struct {
	Name     string `validate:"required,min=2" form:"name"`
	Password string `validate:"required,min=6" form:"password"`
}

// 用作请求校验。
var validate = validator.New()

// Post 处理注册请求。
func (h *SignupHandler) Post(c *fiber.Ctx) error {
	request := &SignUpRequest{}
	if err := c.BodyParser(request); err != nil {
		return h.RenderBadRequestPage(c, err.Error())
	}

	if err := validate.Struct(request); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			log.Printf("failed to sign up, failed to validate request, %v/%s", err, reflect.TypeOf(err))
			return h.RenderError500(c)
		}
		return h.RenderBadRequestPage(c, h.GetValidationErrorDetails(err))
	}

	authRequest := &pb.SignUpRequest{
		Name:     request.Name,
		Password: request.Password,
	}
	authRespone, err := h.authService.SignUp(c.UserContext(), authRequest)
	if err != nil {
		log.Printf("failed to sign up, user: %s, %v", request.Name, err)
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
		"result_title":  "注册成功",
		"result_detail": "注册已完成，如果未自动跳转，请点击右侧链接！",
		"redirect_url":  "/",
	})
}
