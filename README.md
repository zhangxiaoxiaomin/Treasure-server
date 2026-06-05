# 藏品拍卖交流平台 - Go 服务端

## 项目概述

为藏品拍卖交流小程序提供后端 API 支持。基于 Go 语言开发，使用 SQLite 作为数据库。

## 技术栈

- **语言**: Go 1.26
- **数据库**: SQLite (modernc.org/sqlite，纯 Go 实现，无需 CGO)
- **HTTP**: 标准库 net/http

## 项目结构

```
Treasure-server/
├── main.go                          # 入口文件
├── database/
│   └── db.go                        # 数据库初始化
├── models/
│   └── collection.go                # 数据模型
├── repository/
│   └── collection_repo.go           # 数据库操作层
├── handlers/
│   └── collection_handler.go        # HTTP 处理器
├── scripts/
│   └── seed.go                      # 数据种子脚本
├── data/
│   └── treasure.db                  # SQLite 数据库文件
└── README.md
```

## 快速开始

### 1. 运行服务端

```bash
cd Treasure-server
go run main.go
```

服务端将在 `http://localhost:8080` 启动。

### 2. 初始化测试数据

```bash
go run scripts/seed.go
```

### 3. 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| PORT | 服务端口 | 8080 |
| DB_PATH | 数据库路径 | ./data/treasure.db |

## API 接口

### 藏品管理 (CRUD)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/health` | 健康检查 |
| GET | `/api/collections` | 获取藏品列表 |
| GET | `/api/collections/:id` | 获取单个藏品 |
| POST | `/api/collections` | 创建藏品 |
| PUT | `/api/collections/:id` | 更新藏品 |
| DELETE | `/api/collections/:id` | 删除藏品 |

### 查询参数 (GET /api/collections)

| 参数 | 类型 | 说明 |
|------|------|------|
| category | string | 分类筛选 (coins/calligraphy/antiques/porcelain/other) |
| keyword | string | 关键词搜索 (标题/ID) |
| page | int | 页码 (默认 1) |
| pageSize | int | 每页数量 (默认 20) |

## 藏品分类

- `coins` - 钱币
- `calligraphy` - 字画
- `antiques` - 文玩
- `porcelain` - 瓷器
- `other` - 杂项

## 数据模型

### 藏品 (Collection)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | string | 唯一标识 |
| titleCN | string | 中文标题 |
| titleEN | string | 英文标题 |
| category | string | 分类 |
| image | string | 封面图 URL |
| detailImages | string[] | 详情图片列表 |
| views | int | 浏览量 (后台配置) |
| likes | int | 点赞数 (后台配置) |
| commentsCount | int | 评论数 (后台配置) |
| badgeCN | string | 中文徽章 |
| badgeEN | string | 英文徽章 |
| dateStrCN | string | 中文日期 |
| dateStrEN | string | 英文日期 |
| descriptionCN | string | 中文描述 |
| descriptionEN | string | 英文描述 |
| detailDescCN | string | 中文详细介绍 |
| detailDescEN | string | 英文详细介绍 |
| comments | Comment[] | 评论列表 |
| createdAt | timestamp | 创建时间 |
| updatedAt | timestamp | 更新时间 |

## 编译部署

```bash
# 编译
go build -o treasure-server.exe .

# 运行
./treasure-server.exe