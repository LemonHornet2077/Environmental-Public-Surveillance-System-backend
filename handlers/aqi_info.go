package handlers

import (
	"database/sql"
	"epss-backend/database"
	"epss-backend/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllConfirmedAQI 获取所有网格员确认后的AQI信息列表
func GetAllConfirmedAQI(c *fiber.Ctx) error {
	// 查询所有已确认的AQI信息，包括关联的省市、网格员和监督员信息
	query := `
		SELECT 
			s.id, s.province_id, s.city_id, s.address, 
			s.so2_value, s.so2_level, s.co_value, s.co_level, 
			s.spm_value, s.spm_level, s.aqi_id, 
			s.confirm_date, s.confirm_time, s.gm_id, s.fd_id, 
			s.information, IFNULL(s.remarks, '') as remarks,
			p.province_name, c.city_name, 
			IFNULL(gm.gm_name, '') as grid_member_name,
			IFNULL(sup.real_name, '') as supervisor_name,
			a.chinese_explain, a.aqi_explain, a.color, a.health_impact, a.take_steps
		FROM 
			statistics s
		LEFT JOIN 
			grid_province p ON s.province_id = p.province_id
		LEFT JOIN 
			grid_city c ON s.city_id = c.city_id
		LEFT JOIN 
			grid_member gm ON s.gm_id = gm.gm_id
		LEFT JOIN 
			supervisor sup ON s.fd_id = sup.tel_id
		LEFT JOIN 
			aqi a ON s.aqi_id = a.aqi_id
		ORDER BY 
			s.id DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取已确认AQI信息列表失败",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	// 构建AQI信息列表
	var aqiList []fiber.Map
	for rows.Next() {
		var statistics models.Statistics
		var provinceName, cityName, gridMemberName, supervisorName string
		var chineseExplain, aqiExplain, color, healthImpact, takeSteps string
		
		// 使用临时变量接收可能为NULL的字段
		var remarks sql.NullString
		
		err := rows.Scan(
			&statistics.ID, &statistics.ProvinceID, &statistics.CityID, &statistics.Address,
			&statistics.SO2Value, &statistics.SO2Level, &statistics.COValue, &statistics.COLevel,
			&statistics.SPMValue, &statistics.SPMLevel, &statistics.AqiID,
			&statistics.ConfirmDate, &statistics.ConfirmTime, &statistics.GmID, &statistics.FdID,
			&statistics.Information, &remarks,
			&provinceName, &cityName, &gridMemberName, &supervisorName,
			&chineseExplain, &aqiExplain, &color, &healthImpact, &takeSteps,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理AQI数据失败",
				"details": err.Error(),
			})
		}

		// 构建返回数据
		aqiList = append(aqiList, fiber.Map{
			"id":                statistics.ID,
			"province_id":       statistics.ProvinceID,
			"city_id":           statistics.CityID,
			"address":           statistics.Address,
			"so2_value":         statistics.SO2Value,
			"so2_level":         statistics.SO2Level,
			"co_value":          statistics.COValue,
			"co_level":          statistics.COLevel,
			"spm_value":         statistics.SPMValue,
			"spm_level":         statistics.SPMLevel,
			"aqi_id":            statistics.AqiID,
			"confirm_date":      statistics.ConfirmDate,
			"confirm_time":      statistics.ConfirmTime,
			"gm_id":             statistics.GmID,
			"fd_id":             statistics.FdID,
			"information":       statistics.Information,
			"remarks":           remarks.String,
			"province_name":     provinceName,
			"city_name":         cityName,
			"grid_member_name":  gridMemberName,
			"supervisor_name":   supervisorName,
			"chinese_explain":   chineseExplain,
			"aqi_explain":       aqiExplain,
			"color":             color,
			"health_impact":     healthImpact,
			"take_steps":        takeSteps,
		})
	}

	return c.JSON(fiber.Map{
		"data": aqiList,
	})
}
