package handlers

import (
	"epss-backend/database"

	"github.com/gofiber/fiber/v2"
)

// GetAQIList 获取所有空气质量指数级别信息
// 该接口可供所有角色使用
func GetAQIList(c *fiber.Ctx) error {
	// 查询所有AQI数据
	query := `
		SELECT aqi_id, chinese_explain, aqi_explain, color, 
		       health_impact, take_steps, 
		       so2_min, so2_max, co_min, co_max, spm_min, spm_max, remarks
		FROM aqi
		ORDER BY aqi_id ASC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "获取空气质量指数数据失败",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	// 构建AQI列表
	var aqiList []fiber.Map
	for rows.Next() {
		var (
			aqiID          int
			chineseExplain string
			aqiExplain     string
			color          string
			healthImpact   string
			takeSteps      string
			so2Min         int
			so2Max         int
			coMin          int
			coMax          int
			spmMin         int
			spmMax         int
			remarks        string
		)

		err := rows.Scan(
			&aqiID, &chineseExplain, &aqiExplain, &color,
			&healthImpact, &takeSteps,
			&so2Min, &so2Max, &coMin, &coMax, &spmMin, &spmMax, &remarks,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "处理空气质量指数数据失败",
				"error":   err.Error(),
			})
		}

		aqiList = append(aqiList, fiber.Map{
			"aqi_id":          aqiID,
			"chinese_explain": chineseExplain,
			"aqi_explain":     aqiExplain,
			"color":           color,
			"health_impact":   healthImpact,
			"take_steps":      takeSteps,
			"so2_min":         so2Min,
			"so2_max":         so2Max,
			"co_min":          coMin,
			"co_max":          coMax,
			"spm_min":         spmMin,
			"spm_max":         spmMax,
			"remarks":         remarks,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "获取空气质量指数数据成功",
		"data":    aqiList,
	})
}
