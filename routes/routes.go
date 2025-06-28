package routes

import (
	"epss-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes 设置所有路由
func SetupRoutes(app *fiber.App) {
	// API版本前缀
	api := app.Group("/api/v1")

	// 公开路由（不需要认证）
	auth := api.Group("/auth")
	auth.Post("/admin/login", handlers.AdminLogin)
	auth.Post("/member/login", handlers.GridMemberLogin)
	auth.Post("/supervisor/register", handlers.SupervisorRegister)
	auth.Post("/supervisor/login", handlers.SupervisorLogin)

	// 需要管理员认证的路由
	adminProtected := api.Group("/admin")
	adminProtected.Use(handlers.JWTMiddleware)
	adminProtected.Use(handlers.AdminOnly)

	// 管理员功能
	adminProtected.Post("/add", handlers.AddAdmin)
	adminProtected.Post("/member/add", handlers.AddGridMember)
	adminProtected.Delete("/delete/:id", handlers.DeleteAdmin)
	adminProtected.Delete("/member/delete/:id", handlers.DeleteGridMember)

	// 监督员自行删除账户的路由
	supervisorProtected := api.Group("/supervisor")
	supervisorProtected.Use(handlers.JWTMiddleware)
	supervisorProtected.Use(handlers.SupervisorOnly)
	supervisorProtected.Delete("/delete", handlers.DeleteSupervisorSelf)

	// 健康检查
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "服务运行正常",
		})
	})
}
