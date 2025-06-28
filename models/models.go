package models

import "database/sql"

// Admin 对应 'admins' 表
type Admin struct {
	AdminID   int64          `json:"admin_id"`
	AdminCode string         `json:"admin_code"`
	Password  string         `json:"password"`
	Remarks   sql.NullString `json:"remarks"`
}

// Aqi 对应 'aqi' 表
type Aqi struct {
	AqiID          int64          `json:"aqi_id"`
	ChineseExplain string         `json:"chinese_explain"`
	AqiExplain     string         `json:"aqi_explain"`
	Color          string         `json:"color"`
	HealthImpact   string         `json:"health_impact"`
	TakeSteps      string         `json:"take_steps"`
	SO2Min         int            `json:"so2_min"`
	SO2Max         int            `json:"so2_max"`
	COMin          int            `json:"co_min"`
	COMax          int            `json:"co_max"`
	SPMMin         int            `json:"spm_min"`
	SPMMax         int            `json:"spm_max"`
	Remarks        sql.NullString `json:"remarks"`
}

// AqiFeedback 对应 'aqi_feedback' 表
type AqiFeedback struct {
	AfID           int64          `json:"af_id"`
	TelID          string         `json:"tel_id"`
	ProvinceID     int64          `json:"province_id"`
	CityID         int64          `json:"city_id"`
	Address        string         `json:"address"`
	Information    string         `json:"information"`
	EstimatedGrade int            `json:"estimated_grade"`
	AfDate         string         `json:"af_date"`
	AfTime         string         `json:"af_time"`
	GmID           int64          `json:"gm_id"`
	AssignDate     sql.NullString `json:"assign_date"`
	AssignTime     sql.NullString `json:"assign_time"`
	State          int            `json:"state"`
	Remarks        sql.NullString `json:"remarks"`
}

// GridCity 对应 'grid_city' 表
type GridCity struct {
	CityID     int64          `json:"city_id"`
	CityName   string         `json:"city_name"`
	ProvinceID int64          `json:"province_id"`
	Remarks    sql.NullString `json:"remarks"`
}

// GridMember 对应 'grid_member' 表
type GridMember struct {
	GmID       int64          `json:"gm_id"`
	GmName     string         `json:"gm_name"`
	GmCode     string         `json:"gm_code"`
	Password   string         `json:"password"`
	ProvinceID int64          `json:"province_id"`
	CityID     int64          `json:"city_id"`
	Tel        string         `json:"tel"`
	State      int            `json:"state"`
	Remarks    sql.NullString `json:"remarks"`
}

// GridProvince 对应 'grid_province' 表
type GridProvince struct {
	ProvinceID   int64          `json:"province_id"`
	ProvinceName string         `json:"province_name"`
	ProvinceAbbr string         `json:"province_abbr"`
	Remarks      sql.NullString `json:"remarks"`
}

// Statistics 对应 'statistics' 表
type Statistics struct {
	ID          int64          `json:"id"`
	ProvinceID  int64          `json:"province_id"`
	CityID      int64          `json:"city_id"`
	Address     string         `json:"address"`
	SO2Value    int            `json:"so2_value"`
	SO2Level    int            `json:"so2_level"`
	COValue     int            `json:"co_value"`
	COLevel     int            `json:"co_level"`
	SPMValue    int            `json:"spm_value"`
	SPMLevel    int            `json:"spm_level"`
	AqiID       int64          `json:"aqi_id"`
	ConfirmDate string         `json:"confirm_date"`
	ConfirmTime string         `json:"confirm_time"`
	GmID        int64          `json:"gm_id"`
	FdID        string         `json:"fd_id"`
	Information string         `json:"information"`
	Remarks     sql.NullString `json:"remarks"`
}

// Supervisor 对应 'supervisor' 表
type Supervisor struct {
	TelID    string         `json:"tel_id"`
	Password string         `json:"password"`
	RealName string         `json:"real_name"`
	Birthday string         `json:"birthday"`
	Sex      int            `json:"sex"` // 性别，约定 0 为女性，1 为男性
	Remarks  string         `json:"remarks"`
}