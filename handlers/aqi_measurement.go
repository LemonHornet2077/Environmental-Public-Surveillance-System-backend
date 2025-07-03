package handlers

import (
	"epss-backend/database"
	"epss-backend/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SubmitAQIMeasurement 网格员提交实测AQI数据
func SubmitAQIMeasurement(c *fiber.Ctx) error {
	// 从JWT中获取网格员ID
	gmID := c.Locals("user_gm_id")
	if gmID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "未授权访问",
		})
	}

	// 解析请求数据
	type SubmitAQIRequest struct {
		FeedbackID    int    `json:"feedback_id"`    // 关联的反馈ID
		ProvinceID    int64  `json:"province_id"`    // 省份ID
		CityID        int64  `json:"city_id"`        // 城市ID
		Address       string `json:"address"`        // 详细地址
		SO2Value      int    `json:"so2_value"`      // 二氧化硫浓度值
		COValue       int    `json:"co_value"`       // 一氧化碳浓度值
		SPMValue      int    `json:"spm_value"`      // 悬浮颗粒物浓度值
		Information   string `json:"information"`    // 信息描述
		SupervisorTel string `json:"supervisor_tel"` // 反馈者手机号
	}

	var req SubmitAQIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "无效的请求格式",
		})
	}

	// 验证输入
	if req.ProvinceID <= 0 || req.CityID <= 0 || req.Address == "" || 
	   req.SO2Value < 0 || req.COValue < 0 || req.SPMValue < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "请提供有效的测量数据",
		})
	}

	// 根据浓度值确定各项指标的级别
	so2Level, err := getAQILevelForPollutant("so2", req.SO2Value)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("计算二氧化硫级别失败: %v", err),
		})
	}

	coLevel, err := getAQILevelForPollutant("co", req.COValue)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("计算一氧化碳级别失败: %v", err),
		})
	}

	spmLevel, err := getAQILevelForPollutant("spm", req.SPMValue)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("计算悬浮颗粒物级别失败: %v", err),
		})
	}

	// 确定综合AQI级别（取三者中的最高级别）
	aqiID := getMaxLevel(so2Level, coLevel, spmLevel)

	// 获取当前日期和时间
	now := time.Now()
	confirmDate := now.Format("2006-01-02")
	confirmTime := now.Format("15:04:05")

	// 如果提供了反馈ID，则更新对应反馈的状态为已确认(2)
	if req.FeedbackID > 0 {
		updateQuery := "UPDATE aqi_feedback SET state = 2 WHERE af_id = ? AND gm_id = ?"
		_, err = database.DB.Exec(updateQuery, req.FeedbackID, gmID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("更新反馈状态失败: %v", err),
			})
		}
	}

	// 插入实测数据到statistics表
	insertQuery := `
		INSERT INTO statistics (
			province_id, city_id, address, 
			so2_value, so2_level, co_value, co_level, 
			spm_value, spm_level, aqi_id, 
			confirm_date, confirm_time, gm_id, 
			fd_id, information
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := database.DB.Exec(
		insertQuery,
		req.ProvinceID, req.CityID, req.Address,
		req.SO2Value, so2Level, req.COValue, coLevel,
		req.SPMValue, spmLevel, aqiID,
		confirmDate, confirmTime, gmID,
		req.SupervisorTel, req.Information,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("保存AQI数据失败: %v", err),
		})
	}

	// 获取新插入记录的ID
	id, _ := result.LastInsertId()

	// 查询AQI级别信息
	var aqi models.Aqi
	aqiQuery := "SELECT aqi_id, chinese_explain, color FROM aqi WHERE aqi_id = ?"
	err = database.DB.QueryRow(aqiQuery, aqiID).Scan(&aqi.AqiID, &aqi.ChineseExplain, &aqi.Color)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("获取AQI级别信息失败: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"message": "AQI数据提交成功",
		"data": fiber.Map{
			"id":             id,
			"so2_value":      req.SO2Value,
			"so2_level":      so2Level,
			"co_value":       req.COValue,
			"co_level":       coLevel,
			"spm_value":      req.SPMValue,
			"spm_level":      spmLevel,
			"aqi_id":         aqiID,
			"aqi_level":      aqi.ChineseExplain,
			"aqi_color":      aqi.Color,
			"confirm_date":   confirmDate,
			"confirm_time":   confirmTime,
			"feedback_id":    req.FeedbackID,
			"supervisor_tel": req.SupervisorTel,
		},
	})
}

// getAQILevelForPollutant 根据污染物浓度值确定其AQI级别
func getAQILevelForPollutant(pollutantType string, value int) (int, error) {
	var query string
	switch pollutantType {
	case "so2":
		query = "SELECT aqi_id FROM aqi WHERE ? BETWEEN so2_min AND so2_max"
	case "co":
		query = "SELECT aqi_id FROM aqi WHERE ? BETWEEN co_min AND co_max"
	case "spm":
		query = "SELECT aqi_id FROM aqi WHERE ? BETWEEN spm_min AND spm_max"
	default:
		return 0, fmt.Errorf("未知的污染物类型")
	}

	var level int
	err := database.DB.QueryRow(query, value).Scan(&level)
	if err != nil {
		return 0, err
	}

	return level, nil
}

// getMaxLevel 获取多个级别中的最高级别
func getMaxLevel(levels ...int) int {
	max := 0
	for _, level := range levels {
		if level > max {
			max = level
		}
	}
	return max
}
