package handlers

import (
	"epss-backend/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 获取省份分组AQI超标统计
func GetProvinceAQIStats(c *fiber.Ctx) error {
	db := database.DB

	type ProvinceStats struct {
		ProvinceName    string `json:"province_name"`
		ProvinceID      uint   `json:"province_id"`
		SO2ExceedCount  int    `json:"so2_exceed_count"`
		COExceedCount   int    `json:"co_exceed_count"`
		PM25ExceedCount int    `json:"pm25_exceed_count"`
		AQIExceedCount  int    `json:"aqi_exceed_count"`
	}

	var results []ProvinceStats

	// 查询各省份的AQI超标统计
	query := `
		SELECT 
			p.province_name,
			p.province_id,
			SUM(CASE WHEN s.so2_value > 150 THEN 1 ELSE 0 END) as so2_exceed_count,
			SUM(CASE WHEN s.co_value > 4 THEN 1 ELSE 0 END) as co_exceed_count,
			SUM(CASE WHEN s.spm_value > 75 THEN 1 ELSE 0 END) as pm25_exceed_count,
			SUM(CASE WHEN s.aqi_id > 2 THEN 1 ELSE 0 END) as aqi_exceed_count
		FROM 
			statistics s
		JOIN 
			grid_province p ON s.province_id = p.province_id
		GROUP BY 
			p.province_id, p.province_name
		ORDER BY 
			p.province_name
	`

	rows, err := db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "获取省份AQI统计数据失败",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var stat ProvinceStats
		if err := rows.Scan(&stat.ProvinceName, &stat.ProvinceID, &stat.SO2ExceedCount, &stat.COExceedCount, &stat.PM25ExceedCount, &stat.AQIExceedCount); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "解析省份AQI统计数据失败",
				"error":   err.Error(),
			})
		}
		results = append(results, stat)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    results,
	})
}

// 获取AQI指数分布统计
func GetAQILevelStats(c *fiber.Ctx) error {
	db := database.DB

	type AQILevelStats struct {
		Level      string `json:"level"`
		LevelValue int    `json:"level_value"`
		Count      int    `json:"count"`
	}

	var results []AQILevelStats

	// 查询AQI指数分布统计
	// 使用aqi表和statistics表关联查询
	query := `
		SELECT 
			a.chinese_explain as level,
			a.aqi_id as level_value,
			COUNT(*) as count
		FROM 
			statistics s
		JOIN
			aqi a ON s.aqi_id = a.aqi_id
		GROUP BY 
			a.aqi_id, a.chinese_explain
		ORDER BY 
			a.aqi_id
	`

	rows, err := db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "获取AQI指数分布统计失败",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var stat AQILevelStats
		if err := rows.Scan(&stat.Level, &stat.LevelValue, &stat.Count); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "解析AQI指数分布统计失败",
				"error":   err.Error(),
			})
		}
		results = append(results, stat)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    results,
	})
}

// 获取AQI指数趋势统计
func GetAQITrendStats(c *fiber.Ctx) error {
	db := database.DB

	type MonthlyStats struct {
		Month       string `json:"month"`
		ExceedCount int    `json:"exceed_count"`
	}

	var results []MonthlyStats

	// 获取查询参数
	timeRange := c.Query("timeRange", "12months") // 默认查询过去12个月

	// 构建SQL查询
	query := `
		SELECT 
			CONCAT(SUBSTRING(confirm_date, 1, 7)) as month,
			SUM(CASE WHEN aqi_id > 2 THEN 1 ELSE 0 END) as exceed_count
		FROM 
			statistics
	`

	// 根据时间范围参数添加WHERE子句
	var startDate time.Time
	now := time.Now()

	if timeRange == "all" {
		// 不添加时间限制，查询所有数据
		query += `
		GROUP BY 
			month
		ORDER BY 
			month
		`
		
		rows, err := db.Query(query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "获取AQI指数趋势统计失败",
				"error":   err.Error(),
			})
		}
		defer rows.Close()

		for rows.Next() {
			var stat MonthlyStats
			if err := rows.Scan(&stat.Month, &stat.ExceedCount); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": "解析AQI指数趋势统计失败",
					"error":   err.Error(),
				})
			}
			results = append(results, stat)
		}
	} else {
		// 默认查询过去12个月
		startDate = now.AddDate(-1, 0, 0)
		
		query += `
		WHERE 
			STR_TO_DATE(confirm_date, '%Y-%m-%d') >= ?
		GROUP BY 
			month
		ORDER BY 
			month
		`
		
		rows, err := db.Query(query, startDate.Format("2006-01-02"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "获取AQI指数趋势统计失败",
				"error":   err.Error(),
			})
		}
		defer rows.Close()

		for rows.Next() {
			var stat MonthlyStats
			if err := rows.Scan(&stat.Month, &stat.ExceedCount); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": "解析AQI指数趋势统计失败",
					"error":   err.Error(),
				})
			}
			results = append(results, stat)
		}

		// 如果没有数据，生成过去12个月的空数据
		if len(results) == 0 && timeRange == "12months" {
			for i := 0; i < 12; i++ {
				month := now.AddDate(0, -i, 0).Format("2006-01")
				results = append(results, MonthlyStats{Month: month, ExceedCount: 0})
			}
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    results,
	})
}

// 获取空气质量检测数量实时统计
func GetAQIRealtimeStats(c *fiber.Ctx) error {
	db := database.DB

	type RealtimeStats struct {
		TotalCount     int `json:"total_count"`
		GoodCount      int `json:"good_count"`
		ExceedingCount int `json:"exceeding_count"`
	}

	var stats RealtimeStats

	// 查询总检测数量
	row := db.QueryRow("SELECT COUNT(*) FROM statistics")
	if err := row.Scan(&stats.TotalCount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "获取检测总数量失败",
			"error":   err.Error(),
		})
	}

	// 查询良好检测数量 (AQI <= 2, 对应优和良)
	row = db.QueryRow("SELECT COUNT(*) FROM statistics WHERE aqi_id <= 2")
	if err := row.Scan(&stats.GoodCount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "获取良好检测数量失败",
			"error":   err.Error(),
		})
	}

	// 查询超标检测数量 (AQI > 2, 对应轻度污染及以上)
	row = db.QueryRow("SELECT COUNT(*) FROM statistics WHERE aqi_id > 2")
	if err := row.Scan(&stats.ExceedingCount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "获取超标检测数量失败",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    stats,
	})
}
