// Package web 用来处理http请求。
package web

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/template/html"

	"ppzzl.com/tinyblog-go/frontend/common"
	"ppzzl.com/tinyblog-go/frontend/interfaces"
	"ppzzl.com/tinyblog-go/frontend/web/handlers"
	"ppzzl.com/tinyblog-go/frontend/web/handlers/account"
	"ppzzl.com/tinyblog-go/frontend/web/handlers/article"
	"ppzzl.com/tinyblog-go/frontend/web/handlers/my"
)

// Service 用来管理web服务。
type Service struct {
	address string             // 监听的地址
	app     *fiber.App         // 用来启停服务
	ctx     interfaces.Context // 用来获取应用context中的各个组件
}

// NewWebService 用来创建WebServer实例。
func NewWebService(ctx interfaces.Context) *Service {
	templateEngine := html.New("./templates", ".html")
	templateEngine.AddFunc("safe", func(s string) template.HTML {
		return template.HTML(s)
	})
	templateEngine.AddFunc("formatTs", func(s int64, fmt string) template.HTML {
		ts := time.Unix(s, 0)
		return template.HTML(ts.Format(fmt))
	})
	templateEngine.Reload(true) // 仅调试用。
	ws := Service{
		address: common.MustGetEnv(common.EnvVarNameListenAddress, ""),
		app: fiber.New(fiber.Config{
			Views: templateEngine,
			// Override default error handler
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError

				// Retrieve the custom status code if it's a *fiber.Error
				var e *fiber.Error
				if errors.As(err, &e) {
					code = e.Code
				}

				// Send custom error page
				err = ctx.Status(fiber.StatusInternalServerError).SendFile(fmt.Sprintf("./static/html/%d.html", code))
				if err != nil {
					// In case the SendFile fails
					return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}

				// Return from handler
				return nil
			},
		}),
		ctx: ctx,
	}
	ws.registerHandlers()
	return &ws
}

// Run 用来启动监听。
func (ws *Service) Run() {
	log.Printf("listen %s", ws.address)
	ws.app.Listen(ws.address)
}

// Stop 用来停止监听。
func (ws *Service) Stop() {
	ws.app.Shutdown()
}

// registerHandlers 用来初始各个URL对应的handler。
func (ws *Service) registerHandlers() {
	// 网站首页。
	ws.registerHandler("/", handlers.NewHomeHandler(ws.ctx))
	// 搜索。
	ws.registerHandler("/search", handlers.NewSearchHandler(ws.ctx))
	// 账号相关。
	ws.registerHandler("/account/signin", account.NewSigninHandler(ws.ctx))
	ws.registerHandler("/account/signup", account.NewSignupHandler(ws.ctx))
	ws.registerHandler("/account/signout", account.NewSignoutHandler(ws.ctx))
	// 我的。
	ws.registerHandler("/my/article", my.NewArticleHandler(ws.ctx))
	// 用户相关。
	// 文章相关。
	ws.registerHandler("/article/post", article.NewUserArticlePostHandler(ws.ctx))
	ws.registerHandler("/article/:id", article.NewHandler(ws.ctx))
	ws.registerHandler("/article/:id/edit", article.NewEditHandler(ws.ctx))
	//
	ws.app.Get("/_healthz", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	// 静态资源。
	ws.app.Static("/static", "./static")
	ws.app.Use(favicon.New(favicon.Config{
		File: "./static/img/favicon.ico",
		URL:  "/favicon.ico",
	}))
	// 不存在的链接处理。
	ws.app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./static/html/404.html")
	})
}

func (ws *Service) registerHandler(url string, handler any) {
	if getHandler, ok := handler.(handlers.GetHandler); ok {
		ws.app.Get(url, getHandler.Get)
	}
	if postHandler, ok := handler.(handlers.PostHandler); ok {
		ws.app.Post(url, postHandler.Post)
	}
	if putHandler, ok := handler.(handlers.PutHandler); ok {
		ws.app.Put(url, putHandler.Put)
	}
	if patchHandler, ok := handler.(handlers.PatchHandler); ok {
		ws.app.Patch(url, patchHandler.Patch)
	}
	if deleteHandler, ok := handler.(handlers.DeleteHandler); ok {
		ws.app.Delete(url, deleteHandler.Delete)
	}
}
