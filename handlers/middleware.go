package handlers

import (
	"epss-backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWT中间件 - 验证token
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "缺少Authorization头",
		})
	}

	// 检查Bearer前缀
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(401).JSON(fiber.Map{
			"error": "无效的token格式",
		})
	}

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "无效的token",
		})
	}

	// 将claims存储到context中
	if claims, ok := token.Claims.(*Claims); ok {
		c.Locals("user_id", claims.UserID)
		c.Locals("user_tel_id", claims.UserTelID)
		c.Locals("user_type", claims.UserType)
	}

	return c.Next()
}

// 管理员权限中间件
func AdminOnly(c *fiber.Ctx) error {
	userType := c.Locals("user_type")
	if userType != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"error": "需要管理员权限",
		})
	}
	return c.Next()
}

// 监督员权限中间件
func SupervisorOnly(c *fiber.Ctx) error {
	userType := c.Locals("user_type")
	if userType != "supervisor" {
		return c.Status(403).JSON(fiber.Map{
			"error": "需要监督员权限",
		})
	}
	return c.Next()
}

// 网格员权限中间件
func GridMemberOnly(c *fiber.Ctx) error {
	userType := c.Locals("user_type")
	if userType != "member" {
		return c.Status(403).JSON(fiber.Map{
			"error": "需要网格员权限",
		})
	}
	
	// 将用户ID存储为网格员ID
	userID := c.Locals("user_id")
	if userID != nil {
		c.Locals("user_gm_id", userID)
	}
	
	return c.Next()
}