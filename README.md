# Environmental-Public-Surveillance-System-backend
环保公众监督系统的后端

## 项目简介

本项目是环保公众监督系统的后端服务，基于 Go 和 Fiber 框架构建。它为前端应用提供了用户认证、数据管理和业务逻辑处理等核心功能。

## 技术栈

- **语言**: Go
- **Web 框架**: Fiber v2
- **数据库**: MySQL
- **数据库驱动**: `github.com/go-sql-driver/mysql`
- **认证**: JWT (`github.com/golang-jwt/jwt/v5`)
- **配置管理**: `github.com/joho/godotenv`

## 项目结构

采用简化的扁平化项目结构，便于快速开发和维护：

```
.
├── config/             # 配置加载
├── database/           # 数据库连接初始化
├── handlers/           # HTTP 请求处理器（业务逻辑）
├── models/             # 数据模型（数据库表结构体）
├── routes/             # 路由定义
├── scripts/            # SQL脚本
├── .env                # 环境变量配置文件 (需自行创建)
├── .gitignore          # Git 忽略文件
├── go.mod              # Go 模块依赖
├── go.sum              # Go 模块校验和
└── main.go             # 程序入口
```

## 已实现功能

### 用户认证与管理

系统包含三种用户角色：**管理员 (Admin)**、**网格员 (Grid Member)** 和 **公众监督员 (Supervisor)**，并实现了完整的认证和权限管理。

- **管理员**
  - 不能自行注册，只能由已有管理员添加。
  - 可以登录、添加新管理员、添加新网格员。
  - 可以删除管理员和网格员。

- **网格员**
  - 不能自行注册，只能由管理员添加。
  - 可以登录。

- **公众监督员**
  - 可以自行注册和登录。
  - 可以删除自己的账户。

### 安全
- 使用 JWT (JSON Web Token) 进行无状态认证。
- 通过中间件实现严格的路由权限控制。
- 敏感配置（如数据库连接字符串、JWT密钥）通过 `.env` 文件管理，并已加入 `.gitignore`。

## API 端点

### 健康检查
- `GET /api/v1/health`: 检查服务是否正常运行。

### 公开路由 (无需认证)
- `POST /api/v1/auth/admin/login`: 管理员登录
- `POST /api/v1/auth/member/login`: 网格员登录
- `POST /api/v1/auth/supervisor/register`: 公众监督员注册
- `POST /api/v1/auth/supervisor/login`: 公众监督员登录

### 管理员路由 (需要管理员JWT认证)
- `POST /api/v1/admin/add`: 添加新管理员
- `POST /api/v1/admin/member/add`: 添加新网格员
- `DELETE /api/v1/admin/delete/:id`: 删除管理员
- `DELETE /api/v1/admin/member/delete/:id`: 删除网格员
- `GET /api/v1/admin/info`: 获取当前登录管理员信息
- `GET /api/v1/admin/list`: 获取所有管理员列表
- `GET /api/v1/admin/member/list`: 获取所有网格员列表
- `GET /api/v1/admin/supervisor/list`: 获取所有公众监督员列表
- `DELETE /api/v1/admin/supervisor/delete/:tel_id`: 管理员删除公众监督员
- `GET /api/v1/admin/feedback/list`: 获取所有公众反馈数据列表
- `GET /api/v1/admin/aqi/confirmed/list`: 获取所有网格员确认后的AQI信息列表

### 监督员路由 (需要监督员JWT认证)
- `DELETE /api/v1/supervisor/delete`: 监督员自行删除账户
- `GET /api/v1/supervisor/feedback/list`: 监督员查看自己的所有反馈数据

## 如何运行

### 1. 环境准备
- 安装 Go (建议版本 1.18+)
- 安装 MySQL 数据库

### 2. 数据库设置
- 创建一个名为 `nep` 的数据库。
- 执行 `scripts/nep.sql` 或 `scripts/nep_wtd.sql` 文件来创建所需的表结构。

### 3. 配置环境变量
在项目根目录下创建一个 `.env` 文件，并填入以下内容：

```env
# 数据库连接字符串 (Data Source Name)
DB_DSN="your_user:your_password@tcp(127.0.0.1:3306)/nep?parseTime=true"

# 服务器运行端口
SERVER_PORT=":3000"

# JWT 签名密钥 (请使用一个强随机字符串)
JWT_SECRET="your-super-secret-key"
```

### 4. 安装依赖
```bash
go mod tidy
```

### 5. 启动服务
```bash
go run main.go
```
服务启动后，将监听在 `http://127.0.0.1:3000`。

## 注意事项
- **JWT 密钥**: `.env` 文件中的 `JWT_SECRET` 务必使用一个长且复杂的随机字符串以保证安全。
