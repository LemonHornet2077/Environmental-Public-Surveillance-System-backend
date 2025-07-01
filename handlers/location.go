package handlers

import (
	"epss-backend/database"
	"github.com/gofiber/fiber/v2"
)

// GetProvinces 获取所有省份列表
func GetProvinces(c *fiber.Ctx) error {
	// 查询所有省份
	query := `SELECT province_id, province_name, province_abbr FROM grid_province ORDER BY province_id`
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取省份列表失败",
		})
	}
	defer rows.Close()

	// 构建省份列表
	var provinceList []fiber.Map
	for rows.Next() {
		var provinceID int64
		var provinceName, provinceAbbr string
		
		err := rows.Scan(&provinceID, &provinceName, &provinceAbbr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理省份数据失败",
				"details": err.Error(),
			})
		}

		provinceList = append(provinceList, fiber.Map{
			"province_id": provinceID,
			"province_name": provinceName,
			"province_abbr": provinceAbbr,
		})
	}

	return c.JSON(fiber.Map{
		"data": provinceList,
	})
}

// GetCities 获取指定省份的城市列表
func GetCities(c *fiber.Ctx) error {
	// 获取省份ID参数
	provinceID := c.Params("province_id")
	if provinceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "缺少省份ID参数",
		})
	}

	// 查询指定省份的城市
	query := `SELECT city_id, city_name FROM grid_city WHERE province_id = ? ORDER BY city_id`
	rows, err := database.DB.Query(query, provinceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "获取城市列表失败",
		})
	}
	defer rows.Close()

	// 构建城市列表
	var cityList []fiber.Map
	for rows.Next() {
		var cityID int64
		var cityName string
		
		err := rows.Scan(&cityID, &cityName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "处理城市数据失败",
				"details": err.Error(),
			})
		}

		cityList = append(cityList, fiber.Map{
			"city_id": cityID,
			"city_name": cityName,
		})
	}

	return c.JSON(fiber.Map{
		"data": cityList,
	})
}
