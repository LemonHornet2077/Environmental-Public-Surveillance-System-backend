package database

import (
	"database/sql"
	"epss-backend/config"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect initializes the database connection.
func Connect() {
	dsn := config.Config("DB_DSN")
	if dsn == "" {
		log.Fatal("错误: DB_DSN 未在 .env 文件中设置")
	}

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法打开数据库连接: %v", err)
	}

	// 设置连接池参数
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	// 验证数据库连接
	err = DB.Ping()
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	fmt.Println("成功连接到数据库.")
}
