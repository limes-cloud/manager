# 统一应用管理平台

百度知道内部统一应用管理平台，提供应用、用户、角色、菜单、权限、OC、OAuth 登录、数据字典等统一管理能力，采用 proto-first 的微服务架构。

## 功能特性

| 模块 | 说明 |
|---|---|
| 应用管理 | 应用基本信息、登录方式（AppField / AppOAuther）、示例应用 |
| 用户管理 | 用户信息维护、手机号邮箱绑定、用户-部门关系 |
| 组织架构 | 部门、岗位、部门角色、岗位角色、用户部门 |
| 权限体系 | 角色、角色菜单、角色实体、数据作用域（Scope），当前为无租户版本 |
| 认证授权 | 账密登录、图形验证码、邮箱验证码、OAuth 登录（微信/企业微信等）及绑定 |
| 菜单管理 | 菜单维护，支持业务方通过 `RegisterMenu` 自注册菜单 |
| 表单与字段 | 动态表单、字段/实体注册，支持业务方注册实体（`RegisterEntity`） |
| 系统能力 | 字典、公告、操作日志、反馈、系统设置（标题/Logo/水印等） |

## 技术栈

| 维度 | 选型 |
|---|---|
| 语言 | Go 1.24.6 |
| 框架 | go-kratos/kratos v2.8.4 + kratosx v1.3.2 |
| 协议 | gRPC + HTTP（kratos 网关） |
| 数据库 | MySQL（GORM，默认库名 `ns_manager_v2`，支持自动建库） |
| 缓存 | Redis（go-redis/v9） |
| 认证 | JWT + 图形/邮箱验证码 + HMAC-SHA256 签名（signature） |
| 其他 | wechat/v2（OAuth）、base64Captcha、resty、gomail、miniredis |

## 目录结构

```
manager/
├── main.go                        # 入口：初始化 App、注册中间件、自注册菜单/实体
├── api/                           # 独立 Go 模块（github.com/limes-cloud/manager/api）
│   ├── <module>/proto/            # proto 定义（消息 + 服务）
│   ├── <module>/*.pb.go           # 生成代码（grpc/http/validate/errors）
│   └── go.mod                     # 独立模块，可被其他服务引用
├── internal/
│   ├── app/                       # App 层：gRPC/HTTP handler，自注册模式
│   ├── core/                      # 核心上下文与配置（core.Context、InitApp）
│   ├── domain/                    # 领域层：entity + repository 接口 + service
│   ├── infra/                     # 基础设施层：dbs 仓储实现、缓存等
│   ├── middleware/                # 中间件（jwt、自定义校验等）
│   ├── pkg/                       # 工具包（jwt、cache）
│   ├── types/                     # 请求/响应 DTO
│   └── initiator/                 # 启动初始化：自注册菜单、实体注册
├── conf/conf.yaml                 # 本地配置文件
├── deploy/                        # 数据库迁移/初始化脚本
├── static/                        # 静态资源（证书、模板）
└── Dockerfile                     # 多阶段构建镜像
```

## 快速开始

### 环境要求

- Go 1.24+
- MySQL 5.7+/8.0
- Redis

### 本地启动

```bash
# 1. 按需修改配置（数据库/Redis 连接、端口等）
vim conf/conf.yaml

# 2. 启动（首次启动会自动建库建表）
go run main.go
```

启动后：

- HTTP 服务：`http://127.0.0.1:7010`（路由前缀 `/manager/api/...`）
- gRPC 服务：`127.0.0.1:8010`

### 默认账号

- 超级管理员：`superadmin`（可通过配置 `business.superAdminUsernames` 调整）

## 配置说明

`conf/conf.yaml` 为本地开发配置，主要节点：

| 配置节点 | 说明 |
|---|---|
| `server.http` / `server.grpc` | 服务监听地址与端口（HTTP 7010 / gRPC 8010） |
| `database` | MySQL 连接，`autoCreate: true` 自动建库，`initializer` 可执行初始化 SQL |
| `redis` | Redis 实例（验证码、缓存、Token 黑名单） |
| `jwt` | Token 签名与白名单（`business.jwt.whitelist`） |
| `signature` | 服务间调用签名（ak/sk），白名单见 `onlyList` |
| `captcha` | 登录/邮箱验证码配置 |
| `business` | 超级管理员、系统设置（标题/Logo/水印）等业务配置 |

## 构建与部署

### Docker 构建

```bash
docker build -t manager:latest .
```

多阶段构建：编译阶段使用 `golang:alpine`，运行阶段使用 `alpine` 并复制 `static/`、`deploy/`、`conf/`。

### 数据库迁移

首次部署或从旧版升级时，按需在数据库执行 `deploy/` 下的脚本：

| 脚本 | 用途 |
|---|---|
| `init_no_tenant.sql` | 初始化 `ns_manager_v2`（无租户版） |
| `migrate_from_ns_manager.sql` | 从旧 `ns_manager` 迁移数据 |
| `migration_app_id_to_app.sql` | 应用 ID 字段迁移（app_id → app） |
| `migration_app_loginable.sql` | 应用登录方式迁移 |
| `sync_field_userinfo.sql` | 同步 field / userinfo 表结构 |

## 二次开发

- **proto-first**：新增接口先在 `api/<module>/proto/` 定义 proto，用 `kratosx proto client` 生成代码，再实现 handler。
- **自注册模式**：app 层通过 `init() + register()` 自动注册 gRPC/HTTP handler，新增模块无需改动中央注册文件。
- **业务方接入**：菜单、实体可通过 `POST /manager/api/menu/register`、`POST /manager/api/entity/register` 自注册（已在签名/JWT 白名单中）。
- **依赖方向**：`app → domain ← infra`，domain 层不依赖任何具体实现。

## License

[MIT](LICENSE)