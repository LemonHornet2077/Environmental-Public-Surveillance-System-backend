package main

import (
	"epss-backend/config"
	"epss-backend/database"
	"epss-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	// 连接数据库
	database.Connect()

	app := fiber.New()

	// 启用CORS
	app.Use(cors.New())

	// 设置路由
	routes.SetupRoutes(app)

	// 从 .env 文件获取服务器端口
	port := config.Config("SERVER_PORT")
	if port == "" {
		port = ":3000" // 提供一个默认端口
	}

	log.Printf("服务器启动在端口 %s", port)
	log.Fatal(app.Listen(port))
}
