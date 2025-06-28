package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("警告: 无法加载 .env 文件, 将使用系统环境变量: %s", err)
	}
	return os.Getenv(key)
}
