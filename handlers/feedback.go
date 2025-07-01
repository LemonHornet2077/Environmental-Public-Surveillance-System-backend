package handlers

import (
	"database/sql"
	"epss-backend/database"
	"epss-backend/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllFeedbacks 获取所有公众反馈数据列表
func GetAllFeedbacks(c *fiber.Ctx) error {
	// 获取查询参数
	provinceID := c.Query("province_id")
	cityID := c.Query("city_id")

	// 构建基础查询
	baseQuery := `
		SELECT 
			af.af_id, af.tel_id, af.province_id, af.city_id, af.address, 
			af.information, af.estimated_grade, af.af_date, af.af_time, 
			af.gm_id, af.assign_date, af.assign_time, af.state, af.remarks,
			p.province_name, c.city_name, s.real_name as supervisor_name,
			IFNULL(gm.gm_name, '') as grid_member_name
		FROM 
			aqi_feedback af
		LEFT JOIN 
			grid_province p ON af.province_id = p.province_id
		LEFT JOIN 
			grid_city c ON af.city_id = c.city_id
		LEFT JOIN 
			supervisor s ON af.tel_id = s.tel_id
		LEFT JOIN 
			grid_member gm ON af.gm_id = gm.gm_id
	`

	// 添加筛选条件
	whereClause := ""
	params := []interface{}{}

	if provinceID != "" {
		whereClause += " WHERE af.province_id = ?"
		params = append(params, provinceID)

		if cityID != "" {
			whereClause += " AND af.city_id = ?"
			params = append(params, cityID)
		}
	}

	// 完整查询
	query := baseQuery + whereClause + " ORDER BY af.af_id DESC"

	var rows *sql.Rows
	var err error

	if len(params) > 0 {
		rows, err = database.DB.Query(query, params...)
	} else {
		rows, err = database.DB.Query(query)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取反馈列表失败",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	// 构建反馈列表
	var feedbackList []fiber.Map
	for rows.Next() {
		var feedback models.AqiFeedback
		var provinceName, cityName, supervisorName, gridMemberName string
		
		// 使用临时变量接收可能为NULL的字段
		var assignDate, assignTime, remarks sql.NullString
		
		err := rows.Scan(
			&feedback.AfID, &feedback.TelID, &feedback.ProvinceID, &feedback.CityID, &feedback.Address,
			&feedback.Information, &feedback.EstimatedGrade, &feedback.AfDate, &feedback.AfTime,
			&feedback.GmID, &assignDate, &assignTime, &feedback.State, &remarks,
			&provinceName, &cityName, &supervisorName, &gridMemberName,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理反馈数据失败",
				"details": err.Error(),
			})
		}

		// 构建返回数据
		feedbackList = append(feedbackList, fiber.Map{
			"id":               feedback.AfID,
			"tel_id":           feedback.TelID,
			"province_id":      feedback.ProvinceID,
			"city_id":          feedback.CityID,
			"address":          feedback.Address,
			"information":      feedback.Information,
			"estimated_grade":  feedback.EstimatedGrade,
			"af_date":          feedback.AfDate,
			"af_time":          feedback.AfTime,
			"gm_id":            feedback.GmID,
			"assign_date":      assignDate.String,
			"assign_time":      assignTime.String,
			"state":            feedback.State,
			"remarks":          remarks.String,
			"province_name":    provinceName,
			"city_name":        cityName,
			"supervisor_name":  supervisorName,
			"grid_member_name": gridMemberName,
		})
	}

	return c.JSON(fiber.Map{
		"data": feedbackList,
	})
}

// GetSupervisorFeedbacks 获取当前登录的公众监督员的所有反馈数据
func GetSupervisorFeedbacks(c *fiber.Ctx) error {
	// 从JWT中获取监督员ID
	telID := c.Locals("tel_id")
	if telID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "未授权访问",
		})
	}

	// 查询该监督员的所有反馈信息
	query := `
		SELECT 
			af.af_id, af.tel_id, af.province_id, af.city_id, af.address, 
			af.information, af.estimated_grade, af.af_date, af.af_time, 
			af.gm_id, af.assign_date, af.assign_time, af.state, af.remarks,
			p.province_name, c.city_name,
			IFNULL(gm.gm_name, '') as grid_member_name
		FROM 
			aqi_feedback af
		LEFT JOIN 
			grid_province p ON af.province_id = p.province_id
		LEFT JOIN 
			grid_city c ON af.city_id = c.city_id
		LEFT JOIN 
			grid_member gm ON af.gm_id = gm.gm_id
		WHERE 
			af.tel_id = ?
		ORDER BY 
			af.af_id DESC
	`

	rows, err := database.DB.Query(query, telID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取反馈列表失败",
		})
	}
	defer rows.Close()

	// 构建反馈列表
	var feedbackList []fiber.Map
	for rows.Next() {
		var feedback models.AqiFeedback
		var provinceName, cityName, gridMemberName string
		
		// 使用临时变量接收可能为NULL的字段
		var assignDate, assignTime, remarks sql.NullString
		
		err := rows.Scan(
			&feedback.AfID, &feedback.TelID, &feedback.ProvinceID, &feedback.CityID, &feedback.Address,
			&feedback.Information, &feedback.EstimatedGrade, &feedback.AfDate, &feedback.AfTime,
			&feedback.GmID, &assignDate, &assignTime, &feedback.State, &remarks,
			&provinceName, &cityName, &gridMemberName,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理反馈数据失败",
				"details": err.Error(),
			})
		}

		// 构建返回数据
		feedbackList = append(feedbackList, fiber.Map{
			"id":               feedback.AfID,
			"tel_id":           feedback.TelID,
			"province_id":      feedback.ProvinceID,
			"city_id":          feedback.CityID,
			"address":          feedback.Address,
			"information":      feedback.Information,
			"estimated_grade":  feedback.EstimatedGrade,
			"af_date":          feedback.AfDate,
			"af_time":          feedback.AfTime,
			"gm_id":            feedback.GmID,
			"assign_date":      assignDate.String,
			"assign_time":      assignTime.String,
			"state":            feedback.State,
			"remarks":          remarks.String,
			"province_name":    provinceName,
			"city_name":        cityName,
			"grid_member_name": gridMemberName,
		})
	}

	return c.JSON(fiber.Map{
		"data": feedbackList,
	})
}