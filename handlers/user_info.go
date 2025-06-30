package handlers

import (
	"epss-backend/database"
	"epss-backend/models"

	"github.com/gofiber/fiber/v2"
)

// GetCurrentAdmin 获取当前登录的管理员信息
func GetCurrentAdmin(c *fiber.Ctx) error {
	// 从JWT中获取管理员ID
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "未授权访问",
		})
	}

	// 查询管理员信息
	var admin models.Admin
	query := "SELECT admin_id, admin_code, remarks FROM admins WHERE admin_id = ?"
	err := database.DB.QueryRow(query, userID).Scan(
		&admin.AdminID, &admin.AdminCode, &admin.Remarks,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取管理员信息失败",
		})
	}

	return c.JSON(fiber.Map{
		"id":         admin.AdminID,
		"admin_code": admin.AdminCode,
		"remarks":    admin.Remarks.String,
	})
}

// GetAdminList 获取所有管理员列表
func GetAdminList(c *fiber.Ctx) error {
	// 查询所有管理员
	query := "SELECT admin_id, admin_code, remarks FROM admins"
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取管理员列表失败",
		})
	}
	defer rows.Close()

	// 构建管理员列表
	var adminList []fiber.Map
	for rows.Next() {
		var admin models.Admin
		err := rows.Scan(&admin.AdminID, &admin.AdminCode, &admin.Remarks)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理管理员数据失败",
			})
		}

		adminList = append(adminList, fiber.Map{
			"id":         admin.AdminID,
			"admin_code": admin.AdminCode,
			"remarks":    admin.Remarks.String,
		})
	}

	return c.JSON(fiber.Map{
		"data": adminList,
	})
}

// GetGridMemberList 获取所有网格员列表
func GetGridMemberList(c *fiber.Ctx) error {
	// 查询所有网格员
	query := `
		SELECT gm.gm_id, gm.gm_name, gm.gm_code, gm.province_id, gm.city_id, 
		       gm.tel, gm.state, gm.remarks,
		       p.province_name, c.city_name
		FROM grid_member gm
		LEFT JOIN grid_province p ON gm.province_id = p.province_id
		LEFT JOIN grid_city c ON gm.city_id = c.city_id
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取网格员列表失败",
		})
	}
	defer rows.Close()

	// 构建网格员列表
	var memberList []fiber.Map
	for rows.Next() {
		var member models.GridMember
		var provinceName, cityName string
		err := rows.Scan(
			&member.GmID, &member.GmName, &member.GmCode, &member.ProvinceID, &member.CityID,
			&member.Tel, &member.State, &member.Remarks,
			&provinceName, &cityName,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理网格员数据失败",
			})
		}

		memberList = append(memberList, fiber.Map{
			"id":            member.GmID,
			"member_code":   member.GmCode,
			"real_name":     member.GmName,
			"grid_id":       member.GmID,
			"province_id":   member.ProvinceID,
			"city_id":       member.CityID,
			"province_name": provinceName,
			"city_name":     cityName,
			"tel":           member.Tel,
			"state":         member.State,
			"remarks":       member.Remarks.String,
		})
	}

	return c.JSON(fiber.Map{
		"data": memberList,
	})
}

// GetSupervisorList 获取所有公众监督员列表
func GetSupervisorList(c *fiber.Ctx) error {
	// 查询所有公众监督员，使用 IFNULL 处理可能的 NULL 值
	query := "SELECT tel_id, real_name, birthday, sex, IFNULL(remarks, '') as remarks FROM supervisor"
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取公众监督员列表失败",
		})
	}
	defer rows.Close()

	// 构建公众监督员列表
	var supervisorList []fiber.Map
	for rows.Next() {
		var telID, realName, birthday, remarks string
		var sex int
		err := rows.Scan(&telID, &realName, &birthday, &sex, &remarks)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理公众监督员数据失败",
			})
		}

		supervisorList = append(supervisorList, fiber.Map{
			"tel_id":     telID,
			"real_name":  realName,
			"birthday":   birthday,
			"sex":        sex,
			"remarks":    remarks,
		})
	}

	return c.JSON(fiber.Map{
		"data": supervisorList,
	})
}

// DeleteSupervisor 管理员删除公众监督员
func DeleteSupervisor(c *fiber.Ctx) error {
	telID := c.Params("tel_id")
	if telID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "缺少监督员手机号",
		})
	}

	// 检查要删除的监督员是否存在
	var count int
	checkQuery := "SELECT COUNT(*) FROM supervisor WHERE tel_id = ?"
	err := database.DB.QueryRow(checkQuery, telID).Scan(&count)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "数据库查询失败",
		})
	}

	if count == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "监督员不存在",
		})
	}

	// 删除监督员
	deleteQuery := "DELETE FROM supervisor WHERE tel_id = ?"
	_, err = database.DB.Exec(deleteQuery, telID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "删除监督员失败",
		})
	}

	return c.JSON(fiber.Map{
		"message": "监督员删除成功",
	})
}
