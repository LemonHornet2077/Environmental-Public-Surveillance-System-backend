package routes

import (
	"epss-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes 设置所有路由
func SetupRoutes(app *fiber.App) {
	// API版本前缀
	api := app.Group("/api/v1")

	// 公开路由
	api.Get("/health", handlers.HealthCheck)

	// 所有角色共用的公共路由
	public := api.Group("/public")
	{
		// AQI相关
		public.Get("/aqi/list", handlers.GetAQIList)
		public.Get("/aqi/confirmed/list", handlers.GetAllConfirmedAQI)

		// 位置信息相关
		public.Get("/location/provinces", handlers.GetProvinces)
		public.Get("/location/cities/:province_id", handlers.GetCities)
	}

	// 认证相关路由
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
	
	// 获取用户信息接口
	adminProtected.Get("/info", handlers.GetCurrentAdmin)
	adminProtected.Get("/list", handlers.GetAdminList)
	adminProtected.Get("/member/list", handlers.GetGridMemberList)
	adminProtected.Get("/supervisor/list", handlers.GetSupervisorList)
	adminProtected.Delete("/supervisor/delete/:tel_id", handlers.DeleteSupervisor)
	
	// 管理员路由组
	adminGroup := adminProtected.Group("")
	{
		// AQI相关
		adminGroup.Get("/aqi/confirmed/list", handlers.GetAllConfirmedAQI)

		// 反馈相关
		adminGroup.Get("/feedback/list", handlers.GetAllFeedbacks)
		adminGroup.Post("/feedback/assign", handlers.AssignFeedback)

		// 位置信息相关
		adminGroup.Get("/location/provinces", handlers.GetProvinces)
		adminGroup.Get("/location/cities/:province_id", handlers.GetCities)

		// 统计数据相关
		adminGroup.Get("/stats/province", handlers.GetProvinceAQIStats)
		adminGroup.Get("/stats/aqi-level", handlers.GetAQILevelStats)
		adminGroup.Get("/stats/aqi-trend", handlers.GetAQITrendStats)
		adminGroup.Get("/stats/aqi-realtime", handlers.GetAQIRealtimeStats)
	}

	// 监督员相关路由
	supervisorProtected := api.Group("/supervisor")
	supervisorProtected.Use(handlers.JWTMiddleware)
	supervisorProtected.Use(handlers.SupervisorOnly)
	supervisorProtected.Get("/info", handlers.GetCurrentSupervisor) // 获取当前登录的监督员信息
	supervisorProtected.Delete("/delete", handlers.DeleteSupervisorSelf)
	supervisorProtected.Get("/feedback/list", handlers.GetSupervisorFeedbacks)
	supervisorProtected.Post("/feedback/submit", handlers.SubmitFeedback)

	// 网格员相关路由
	memberProtected := api.Group("/member")
	memberProtected.Use(handlers.JWTMiddleware)
	memberProtected.Use(handlers.GridMemberOnly)
	memberProtected.Get("/info", handlers.GetCurrentGridMember) // 获取当前登录的网格员信息
	memberProtected.Get("/feedback/list", handlers.GetGridMemberFeedbacks) // 获取分配给当前网格员的反馈任务
	memberProtected.Post("/aqi/submit", handlers.SubmitAQIMeasurement) // 提交实测的AQI数据

	// 健康检查
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "服务运行正常",
		})
	})
}
