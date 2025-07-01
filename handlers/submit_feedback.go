package handlers

import (
	"epss-backend/database"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SubmitFeedback 公众监督员提交反馈数据
func SubmitFeedback(c *fiber.Ctx) error {
	// 从JWT中获取监督员ID
	telID := c.Locals("tel_id")
	if telID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "未授权访问",
		})
	}

	// 解析请求体
	type FeedbackRequest struct {
		ProvinceID     int64  `json:"province_id"`
		CityID         int64  `json:"city_id"`
		Address        string `json:"address"`
		Information    string `json:"information"`
		EstimatedGrade int    `json:"estimated_grade"`
	}

	var request FeedbackRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "请求数据格式错误",
			"details": err.Error(),
		})
	}

	// 验证请求数据
	if request.ProvinceID <= 0 || request.CityID <= 0 || request.Address == "" || request.Information == "" || request.EstimatedGrade <= 0 || request.EstimatedGrade > 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "请求数据不完整或无效",
		})
	}

	// 获取当前日期和时间
	now := time.Now()
	afDate := now.Format("2006-01-02")
	afTime := now.Format("15:04:05")

	// 插入反馈数据
	query := `
		INSERT INTO aqi_feedback 
		(tel_id, province_id, city_id, address, information, estimated_grade, af_date, af_time, gm_id, state) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, 0, 0)
	`

	result, err := database.DB.Exec(
		query,
		telID,
		request.ProvinceID,
		request.CityID,
		request.Address,
		request.Information,
		request.EstimatedGrade,
		afDate,
		afTime,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "提交反馈数据失败",
			"details": err.Error(),
		})
	}

	// 获取插入的ID
	afID, err := result.LastInsertId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取反馈ID失败",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "反馈数据提交成功",
		"feedback_id": afID,
		"tel_id": telID,
		"submit_time": fmt.Sprintf("%s %s", afDate, afTime),
	})
}
