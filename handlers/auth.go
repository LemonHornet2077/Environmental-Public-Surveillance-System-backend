package handlers

import (
	"epss-backend/config"
	"epss-backend/database"
	"epss-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWT Claims 结构
type Claims struct {
	UserID   int64  `json:"user_id"`
	UserType string `json:"user_type"` // "admin", "member", "supervisor"
	jwt.RegisteredClaims
}

// 生成JWT Token
func generateToken(userID int64, userType string) (string, error) {
	claims := Claims{
		UserID:   userID,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config("JWT_SECRET")))
}

// 管理员登录
func AdminLogin(c *fiber.Ctx) error {
	type LoginRequest struct {
		AdminCode string `json:"admin_code"`
		Password  string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "无效的请求格式",
		})
	}

	// 验证输入
	if req.AdminCode == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "管理员编码和密码不能为空",
		})
	}

	// 查询管理员
	var admin models.Admin
	query := "SELECT admin_id, admin_code, password, remarks FROM admins WHERE admin_code = ?"
	err := database.DB.QueryRow(query, req.AdminCode).Scan(
		&admin.AdminID, &admin.AdminCode, &admin.Password, &admin.Remarks,
	)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "管理员编码或密码错误",
		})
	}

	// 验证密码（注意：在生产环境中应该使用哈希密码）
	if admin.Password != req.Password {
		return c.Status(401).JSON(fiber.Map{
			"error": "管理员编码或密码错误",
		})
	}

	// 生成JWT token
	token, err := generateToken(admin.AdminID, "admin")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "生成token失败",
		})
	}

	return c.JSON(fiber.Map{
		"message": "登录成功",
		"token":   token,
		"admin": fiber.Map{
			"admin_id":   admin.AdminID,
			"admin_code": admin.AdminCode,
		},
	})
}

// 添加新管理员（需要管理员权限）
func AddAdmin(c *fiber.Ctx) error {
	type AddAdminRequest struct {
		AdminCode string `json:"admin_code"`
		Password  string `json:"password"`
		Remarks   string `json:"remarks"`
	}

	var req AddAdminRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "无效的请求格式",
		})
	}

	// 验证输入
	if req.AdminCode == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "管理员编码和密码不能为空",
		})
	}

	// 检查管理员编码是否已存在
	var count int
	checkQuery := "SELECT COUNT(*) FROM admins WHERE admin_code = ?"
	err := database.DB.QueryRow(checkQuery, req.AdminCode).Scan(&count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "数据库查询失败",
		})
	}

	if count > 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "管理员编码已存在",
		})
	}

	// 插入新管理员
	insertQuery := "INSERT INTO admins (admin_code, password, remarks) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(insertQuery, req.AdminCode, req.Password, req.Remarks)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "添加管理员失败",
		})
	}

	adminID, _ := result.LastInsertId()

	return c.JSON(fiber.Map{
		"message":  "管理员添加成功",
		"admin_id": adminID,
	})
}

// 网格员登录
func GridMemberLogin(c *fiber.Ctx) error {
	type LoginRequest struct {
		GmCode   string `json:"gm_code"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效的请求格式"})
	}

	if req.GmCode == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "网格员编码和密码不能为空"})
	}

	var member models.GridMember
	query := "SELECT gm_id, gm_code, password FROM grid_member WHERE gm_code = ?"
	err := database.DB.QueryRow(query, req.GmCode).Scan(&member.GmID, &member.GmCode, &member.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "网格员编码或密码错误"})
	}

	if member.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "网格员编码或密码错误"})
	}

	token, err := generateToken(member.GmID, "member")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "生成token失败"})
	}

	return c.JSON(fiber.Map{
		"message": "登录成功",
		"token":   token,
		"member": fiber.Map{
			"gm_id":   member.GmID,
			"gm_code": member.GmCode,
		},
	})
}

// 添加新网格员（需要管理员权限）
func AddGridMember(c *fiber.Ctx) error {
	var req models.GridMember
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效的请求格式"})
	}

	if req.GmCode == "" || req.Password == "" || req.GmName == "" || req.Tel == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "姓名、编码、密码和电话为必填项"})
	}

	// 检查网格员编码是否已存在
	var count int
	checkQuery := "SELECT COUNT(*) FROM grid_member WHERE gm_code = ?"
	err := database.DB.QueryRow(checkQuery, req.GmCode).Scan(&count)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "数据库查询失败"})
	}

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "网格员编码已存在"})
	}

	insertQuery := `INSERT INTO grid_member 
		(gm_name, gm_code, password, province_id, city_id, tel, state, remarks) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(insertQuery,
		req.GmName, req.GmCode, req.Password, req.ProvinceID, req.CityID, req.Tel, req.State, req.Remarks)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "添加网格员失败"})
	}

	gmID, _ := result.LastInsertId()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "网格员添加成功",
		"gm_id":   gmID,
	})
}

// 公众监督员注册
func SupervisorRegister(c *fiber.Ctx) error {
	var req models.Supervisor
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效的请求格式"})
	}

	if req.TelID == "" || req.Password == "" || req.RealName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "手机号、密码和姓名为必填项"})
	}

	// 检查手机号是否已注册
	var count int
	checkQuery := "SELECT COUNT(*) FROM supervisor WHERE tel_id = ?"
	err := database.DB.QueryRow(checkQuery, req.TelID).Scan(&count)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "数据库查询失败"})
	}

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "该手机号已被注册"})
	}

	insertQuery := `INSERT INTO supervisor 
		(tel_id, password, real_name, birthday, sex, remarks) 
		VALUES (?, ?, ?, ?, ?, ?)`
	
	// 处理 remarks 字段，如果为空字符串则插入 NULL
	var remarks interface{}
	if req.Remarks == "" {
		remarks = nil
	} else {
		remarks = req.Remarks
	}
	
	_, err = database.DB.Exec(insertQuery,
		req.TelID, req.Password, req.RealName, req.Birthday, req.Sex, remarks)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "注册失败"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "注册成功"})
}

// 公众监督员登录
func SupervisorLogin(c *fiber.Ctx) error {
	type LoginRequest struct {
		TelID    string `json:"tel_id"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效的请求格式"})
	}

	if req.TelID == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "手机号和密码不能为空"})
	}

	var supervisor models.Supervisor
	// 注意：tel_id是varchar类型，在数据库中是主键，不能自增，所以没有supervisor_id
	query := "SELECT tel_id, password FROM supervisor WHERE tel_id = ?"
	err := database.DB.QueryRow(query, req.TelID).Scan(&supervisor.TelID, &supervisor.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "手机号或密码错误"})
	}

	if supervisor.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "手机号或密码错误"})
	}

	// 由于supervisor的ID是tel_id(string)，而我们的token生成函数需要int64，这里我们暂时传入0。
	// 在实际应用中，可能需要调整token的payload或supervisor表结构。
	token, err := generateToken(0, "supervisor")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "生成token失败"})
	}

	return c.JSON(fiber.Map{
		"message": "登录成功",
		"token":   token,
		"supervisor": fiber.Map{
			"tel_id": supervisor.TelID,
		},
	})
}
