# Environmental-Public-Surveillance-System-backend

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

### 数据可视化大屏

系统提供了完整的数据可视化大屏功能，支持以下数据的实时监控和统计分析：

- **省份分组AQI超标统计**
  - 按省份统计AQI总体超标数量
  - 分项统计SO2、PM2.5、CO三种污染物的超标数量
  - 支持数据可视化展示，使用柱状图直观呈现

- **AQI指数分布统计**
  - 统计各AQI级别（优、良、轻度污染、中度污染、重度污染、严重污染）的数量分布
  - 使用饼图展示各级别占比情况

- **AQI指数趋势统计**
  - 按月份统计AQI超标数量的历史趋势
  - 支持查看过去12个月或全部历史数据
  - 使用折线图展示趋势变化

- **实时检测统计**
  - 统计总检测数量、良好检测数量和超标检测数量
  - 计算良好率和超标率
  - 数据每5秒自动刷新，确保实时性

### 用户认证与管理

系统包含三种用户角色：**管理员 (Admin)**、**网格员 (Grid Member)** 和 **公众监督员 (Supervisor)**，并实现了完整的认证和权限管理。

- **管理员**
  - 不能自行注册，只能由已有管理员添加。
  - 可以登录、添加新管理员、添加新网格员。
  - 可以删除管理员和网格员。

- **网格员**
  - 不能自行注册，只能由管理员添加。
  - 可以登录。
  - 可以查看分配给自己的公众反馈任务。
  - 可以提交实测的AQI数据，包括二氧化硫、一氧化碳和悬浮颗粒物的浓度值。

- **公众监督员**
  - 可以自行注册和登录。
  - 可以删除自己的账户。

### 公众反馈与任务指派

系统实现了完整的公众反馈和任务指派流程：

- **反馈提交**
  - 公众监督员可以提交环保相关反馈，包含地理位置、详细信息和预估空气质量等级。
  - 提交的反馈初始状态为“未指派”(state=0)。

- **任务指派**
  - 管理员可以将未处理的反馈指派给网格员处理。
  - 支持本地指派：优先将反馈指派给同一地区的网格员。
  - 支持异地指派：当本地网格员人数不足时，可指派给其他地区的网格员。
  - 指派时会记录指派日期、时间和备注信息。
  - 指派后反馈状态变为“已指派”(state=1)。

- **数据完整性**
  - 使用数据库事务确保指派过程的原子性和数据一致性。
  - 指派前进行多重验证，包括反馈和网格员存在性、状态检查、区域匹配等。

### 安全
- 使用 JWT (JSON Web Token) 进行无状态认证。
- 通过中间件实现严格的路由权限控制。
- 敏感配置（如数据库连接字符串、JWT密钥）通过 `.env` 文件管理，并已加入 `.gitignore`。

## API 端点

### 健康检查
- `GET /api/v1/health`: 检查服务是否正常运行。

### 公共接口（所有角色可访问）

- `GET /api/v1/health`: 系统健康状态检查
- `GET /api/v1/public/aqi/list`: 获取所有空气质量指数级别数据
- `GET /api/v1/public/aqi/confirmed/list`: 获取所有已确认的AQI信息
- `GET /api/v1/public/location/provinces`: 获取所有省份列表
- `GET /api/v1/public/location/cities/:province_id`: 获取指定省份的城市列表

### 认证相关

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
- `GET /api/v1/admin/feedback/list`: 获取所有公众反馈数据列表，支持通过province_id和city_id参数筛选
- `POST /api/v1/admin/feedback/assign`: 将公众反馈任务指派给网格员，支持本地和异地指派
- `GET /api/v1/admin/aqi/confirmed/list`: 获取所有网格员确认后的AQI信息列表，支持通过province_id和city_id参数筛选
- `GET /api/v1/admin/location/provinces`: 获取所有省份列表
- `GET /api/v1/admin/location/cities/:province_id`: 获取指定省份的城市列表

### 统计数据路由 (需要管理员JWT认证)
- `GET /admin/stats/province`: 获取按省份分组的AQI超标统计数据，包括总体AQI、SO2、PM2.5、CO三种污染物的超标数量
- `GET /admin/stats/aqi-level`: 获取AQI指数级别分布统计数据，统计各级别（优、良、轻度污染等）的数量
- `GET /admin/stats/aqi-trend`: 获取AQI指数趋势统计数据，支持timeRange参数（12months或all）
- `GET /admin/stats/aqi-realtime`: 获取空气质量检测数量实时统计数据，包括总检测数量、良好检测数量和超标检测数量

### 监督员路由 (需要监督员JWT认证)
- `DELETE /api/v1/supervisor/delete`: 监督员自行删除账户
- `GET /api/v1/supervisor/feedback/list`: 监督员查看自己的所有反馈数据
- `POST /api/v1/supervisor/feedback/submit`: 监督员提交反馈数据

### 网格员路由 (需要网格员JWT认证)
- `GET /api/v1/member/info`: 获取当前登录的网格员信息
- `GET /api/v1/member/feedback/list`: 网格员查看分配给自己的反馈任务，支持通过state参数筛选任务状态
- `POST /api/v1/member/aqi/submit`: 网格员提交实测的AQI数据，包括二氧化硫、一氧化碳和悬浮颗粒物的浓度值

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
- **数据可视化大屏**: 数据大屏接口支持高频率调用（每5秒一次），在生产环境中可能需要添加缓存机制以减轻数据库压力。
- **统计数据API**: 统计数据API返回的是JSON格式，前端需要进行适当的数据处理和格式化才能在图表中正确显示。

## 未来计划

### 待实现功能

- **分项污染物浓度超标统计优化**: 优化现有的省分组统计接口，提供更详细的分项数据，如按时间段统计。
- **统计数据缓存机制**: 为高频调用的统计接口添加缓存机制，减少数据库查询压力。
- **历史数据归档**: 实现数据归档功能，优化大数据量下的查询性能。
