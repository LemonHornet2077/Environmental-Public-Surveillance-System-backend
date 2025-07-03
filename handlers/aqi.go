package handlers

import (
	"epss-backend/database"
	"epss-backend/models"

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
		// 使用models.Aqi结构体
		var aqi models.Aqi

		err := rows.Scan(
			&aqi.AqiID, &aqi.ChineseExplain, &aqi.AqiExplain, &aqi.Color,
			&aqi.HealthImpact, &aqi.TakeSteps,
			&aqi.SO2Min, &aqi.SO2Max, &aqi.COMin, &aqi.COMax, &aqi.SPMMin, &aqi.SPMMax, &aqi.Remarks,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "处理空气质量指数数据失败",
				"error":   err.Error(),
			})
		}

		aqiList = append(aqiList, fiber.Map{
			"aqi_id":          aqi.AqiID,
			"chinese_explain": aqi.ChineseExplain,
			"aqi_explain":     aqi.AqiExplain,
			"color":           aqi.Color,
			"health_impact":   aqi.HealthImpact,
			"take_steps":      aqi.TakeSteps,
			"so2_min":         aqi.SO2Min,
			"so2_max":         aqi.SO2Max,
			"co_min":          aqi.COMin,
			"co_max":          aqi.COMax,
			"spm_min":         aqi.SPMMin,
			"spm_max":         aqi.SPMMax,
			"remarks":         aqi.Remarks.String, // 使用String属性获取值
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "获取空气质量指数数据成功",
		"data":    aqiList,
	})
}