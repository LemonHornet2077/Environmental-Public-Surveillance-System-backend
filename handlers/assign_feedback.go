package handlers

import (
	"database/sql"
	"epss-backend/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// AssignFeedback 将反馈任务指派给网格员
func AssignFeedback(c *fiber.Ctx) error {
    // 解析请求体
    var req struct {
        FeedbackID    int    `json:"feedback_id"`
        GridMemberID  int    `json:"grid_member_id"`
        Remarks       string `json:"remarks"`
    }
    
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "无效的请求数据",
            "details": err.Error(),
        })
    }
    
    // 参数验证
    if req.FeedbackID <= 0 || req.GridMemberID <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "反馈ID和网格员ID必须为正整数",
        })
    }
    
    // 获取当前日期和时间
    now := time.Now()
    assignDate := now.Format("2006-01-02")
    assignTime := now.Format("15:04:05")
    
    // 开始数据库事务
    tx, err := database.DB.Begin()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "启动数据库事务失败",
            "details": err.Error(),
        })
    }
    defer tx.Rollback()
    
    // 1. 检查反馈信息是否存在且未指派
    var feedback struct {
        AfID       int
        ProvinceID int
        CityID     int
        State      int
    }
    
    err = tx.QueryRow(
        "SELECT af_id, province_id, city_id, state FROM aqi_feedback WHERE af_id = ? AND state = 0",
        req.FeedbackID,
    ).Scan(&feedback.AfID, &feedback.ProvinceID, &feedback.CityID, &feedback.State)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "反馈信息不存在或已被指派",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "查询反馈信息失败",
            "details": err.Error(),
        })
    }
    
    // 2. 检查网格员是否存在且处于工作状态
    var gridMember struct {
        GmID       int
        ProvinceID int
        CityID     int
        State      int
    }
    
    err = tx.QueryRow(
        "SELECT gm_id, province_id, city_id, state FROM grid_member WHERE gm_id = ? AND state = 0",
        req.GridMemberID,
    ).Scan(&gridMember.GmID, &gridMember.ProvinceID, &gridMember.CityID, &gridMember.State)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "网格员不存在或不处于工作状态",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "查询网格员信息失败",
            "details": err.Error(),
        })
    }
    
    // 3. 检查网格员负责区域是否与反馈信息区域匹配
    if gridMember.ProvinceID != feedback.ProvinceID || gridMember.CityID != feedback.CityID {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "网格员负责区域与反馈信息区域不匹配",
        })
    }
    
    // 4. 更新反馈信息
    _, err = tx.Exec(
        "UPDATE aqi_feedback SET gm_id = ?, assign_date = ?, assign_time = ?, state = 1, remarks = ? WHERE af_id = ?",
        req.GridMemberID, assignDate, assignTime, req.Remarks, req.FeedbackID,
    )
    
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "更新反馈信息失败",
            "details": err.Error(),
        })
    }
    
    // 提交事务
    if err := tx.Commit(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "提交数据库事务失败",
            "details": err.Error(),
        })
    }
    
    // 返回成功响应
    return c.JSON(fiber.Map{
        "success": true,
        "message": "任务已成功指派给网格员",
        "data": fiber.Map{
            "feedback_id": req.FeedbackID,
            "grid_member_id": req.GridMemberID,
            "assign_date": assignDate,
            "assign_time": assignTime,
        },
    })
}
