// main
package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// WebServer 用来管理web服务。
type WebServer struct {
	host string     // 监听的主机名或IP地址
	port int        // 监听的端口号
	app  *fiber.App // 用来启停服务
}

// NewWebServer 用来创建WebServer实例。
func NewWebServer(host string, port int) *WebServer {
	ws := WebServer{
		host: host,
		port: port,
		app:  fiber.New(),
	}
	ws.initializeHandlers()
	return &ws
}

// Run 用来启动监听。
func (ws *WebServer) Run() {
	addr := fmt.Sprintf("%s:%d", ws.host, ws.port)
	ws.app.Listen(addr)
}

// Stop 用来停止监听。
func (ws *WebServer) Stop() {
	ws.app.Shutdown()
}

// initializeHandlers 用来初始各个URL对应的handler。
func (ws *WebServer) initializeHandlers() {
	ws.app.Get("/direct", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	ws.app.Get("/simple/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("'name' is required")
		}
		content := fmt.Sprintf("Hello, %s!", name)
		return c.SendString(content)
	})
}

func main() {
	ws := NewWebServer("0.0.0.0", 8001)
	ws.initializeHandlers()
	ws.Run()
}
