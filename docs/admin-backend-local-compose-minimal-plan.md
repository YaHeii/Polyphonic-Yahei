# Admin Backend Local Compose Minimal Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 用 `docker compose + .env` 搭起本地最小后端链路，只保证 `admin-api + blog-rpc + postgres + redis + rabbitmq` 可启动、`/ping` 正常、后台基础 CRUD 可联调，并保留现有鉴权链路。

**Architecture:** 这次不做生产化编排，也不引入 etcd。`admin-api` 直连 `blog-rpc`，两者都通过 go-zero 的 `conf.UseEnv()` 从 YAML 中解析 `${ENV}`；数据库初始化拆成 `blog-init.sql`（表结构）和新增 `seed.sql`（最小可登录数据）；上传链路改成本地存储优先，但保留 Qiniu 路径，不改现有业务接口形状。

**Tech Stack:** Go, go-zero, Docker Compose, PostgreSQL, Redis, RabbitMQ, bcrypt, local file storage

---

## 范围和验收边界

- 本计划只覆盖本地最小链路
- 必须保留现有登录鉴权，不做 dev bypass
- 邮件、手机验证码、第三方登录可以保留不可用，但不能阻塞服务启动
- 本地上传要能完成“上传后返回可访问 URL”
- 不改 `.api` / proto 接口协议
- 不修改 `_rpc.go` 这类生成文件

## 关键现状

1. [`service/api/admin/admin.go`](../service/api/admin/admin.go) 和 [`service/rpc/blog/blog.go`](../service/rpc/blog/blog.go) 现在都没有启用 `conf.UseEnv()`，YAML 里的 `${VAR}` 不会展开。
2. [`service/api/admin/etc/admin-api.yaml`](../service/api/admin/etc/admin-api.yaml) 里当前没有可运行的 `BlogRpcConf` 直连配置，且存在明文敏感信息。
3. [`service/rpc/blog/etc/blog.yaml`](../service/rpc/blog/etc/blog.yaml) 里 PostgreSQL / Redis / RabbitMQ / 邮件 / 第三方登录都还是明文配置。
4. [`service/api/admin/internal/svc/service_context.go`](../service/api/admin/internal/svc/service_context.go) 现在硬编码 `oss.NewQiniu(c.UploadConfig)`。
5. [`pkg/oss/local.go`](../pkg/oss/local.go) 当前把“磁盘目录”和“返回 URL 前缀”混在 `dir` 一个字段里；如果直接拿来跑 compose，上传、列举、删除会互相打架。
6. [`service/api/admin/infra/staticx/static.go`](../service/api/admin/infra/staticx/static.go) 虽然有 `/static/*` 路由逻辑，但当前启动链路里没有注册它。
7. 登录链路仍然是真实链路：`admin-api /login -> AccountRpc.Login -> t_user + bcrypt 校验`，所以必须补 `seed.sql`。
8. 后台首屏和权限接口依赖 `t_user / t_role / t_user_role / t_menu / t_role_menu / t_api / t_role_api` 的基础数据。

## 文件结构与职责

### 配置与入口

- `service/api/admin/admin.go`
  - 为 `admin-api` 启用 `conf.UseEnv()`
  - 确保插件 / 静态路由注册接入启动链路

- `service/rpc/blog/blog.go`
  - 为 `blog-rpc` 启用 `conf.UseEnv()`

- `service/api/admin/etc/admin-api.yaml`
  - 把本地 compose 所需配置改成 `${VAR}` 占位
  - 增加 `BlogRpcConf.Endpoints`
  - 给本地上传配置留出显式字段

- `service/rpc/blog/etc/blog.yaml`
  - 把 PostgreSQL / Redis / RabbitMQ / 邮件 / 第三方登录配置改成 `${VAR}` 占位

- `.env.example`
  - 作为本地 compose 的环境变量模板

### 上传链路

- `pkg/oss/oss.go`
  - 补本地模式所需配置字段
  - 明确区分存储目录和对外 URL 前缀

- `pkg/oss/local.go`
  - 调整本地上传器实现
  - 保证 `UploadFile` / `DeleteFile` / `ListFiles` 的路径语义一致

- `pkg/oss/qiniu.go`
  - 保持现状，仅按新配置结构做兼容修改

- `service/api/admin/internal/svc/service_context.go`
  - 根据配置选择 `NewLocal(...)` 或 `NewQiniu(...)`

- `service/api/admin/internal/plugins/plugins.go`
  - 注册静态文件路由
  - 保留 swagger/knife4j 路由

- `service/api/admin/infra/staticx/static.go`
  - 如有必要，提炼出一个明确的本地资源路由注册函数

### 容器与初始化

- `docker-compose.yaml`
  - 定义 `postgres / redis / rabbitmq / blog-rpc / admin-api`

- `docker/admin-api/Dockerfile`
  - 容器内 `go build` 构建 `admin-api`

- `docker/blog-rpc/Dockerfile`
  - 容器内 `go build` 构建 `blog-rpc`

- `blog-init.sql`
  - 保持 schema 真相源，不塞种子数据

- `seed.sql`
  - 新增最小后台可登录数据

- `.dockerignore`
  - 限制构建上下文

### 可选脚本 / 文档

- `docs/admin-backend-local-compose-minimal-plan.md`
  - 当前计划文档

## 实施策略

拆成五段：

1. 配置改造
2. 本地上传链路改造
3. 容器构建与 compose 编排
4. 数据库种子初始化
5. 启动验证与联调验收

每段都要求：

- 先写失败验证或最小检查
- 再做最小实现
- 再跑局部验证
- 再提交一个 commit

## Task 1: 启用环境变量配置并收口 YAML

**Files:**
- Modify: `service/api/admin/admin.go`
- Modify: `service/rpc/blog/blog.go`
- Modify: `service/api/admin/etc/admin-api.yaml`
- Modify: `service/rpc/blog/etc/blog.yaml`
- Create: `.env.example`
- Test: `service/api/admin/internal/logic/ping_logic_test.go`

- [ ] **Step 1: 先确认当前入口未启用 `UseEnv`**

Run:

```bash
sed -n '1,80p' service/api/admin/admin.go
sed -n '1,80p' service/rpc/blog/blog.go
```

Expected:

```text
conf.MustLoad(*configFile, &c)
```

- [ ] **Step 2: 为 `admin-api` 启用环境变量展开**

将 [`service/api/admin/admin.go`](../service/api/admin/admin.go) 中：

```go
conf.MustLoad(*configFile, &c)
```

改成：

```go
conf.MustLoad(*configFile, &c, conf.UseEnv())
```

- [ ] **Step 3: 为 `blog-rpc` 启用环境变量展开**

将 [`service/rpc/blog/blog.go`](../service/rpc/blog/blog.go) 中：

```go
conf.MustLoad(*configFile, &c)
```

改成：

```go
conf.MustLoad(*configFile, &c, conf.UseEnv())
```

- [ ] **Step 4: 改写 `admin-api.yaml` 为 compose 可注入形式**

目标配置骨架：

```yaml
Name: admin-api
Host: 0.0.0.0
Port: ${ADMIN_API_PORT}
Mode: dev
Timeout: 60000
DevServer:
  Enabled: true
  Port: ${ADMIN_API_DEVSERVER_PORT}
  HealthPath: "/ping"
  MetricsPath: "/metrics"
  EnableMetrics: true
Log:
  Mode: console
  Encoding: plain
  Path: runtime/admin-api/log
UploadConfig:
  provider: ${UPLOAD_PROVIDER}
  local-dir: ${UPLOAD_LOCAL_DIR}
  local-base-url: ${UPLOAD_LOCAL_BASE_URL}
  zone: ${QINIU_ZONE}
  endpoint: ${QINIU_ENDPOINT}
  access-key-id: ${QINIU_ACCESS_KEY_ID}
  access-key-secret: ${QINIU_ACCESS_KEY_SECRET}
  bucket-name: ${QINIU_BUCKET_NAME}
  bucket-url: ${QINIU_BUCKET_URL}
RedisConf:
  db: ${REDIS_DB}
  host: ${REDIS_HOST}
  port: "${REDIS_PORT}"
  password: ${REDIS_PASSWORD}
BlogRpcConf:
  Endpoints:
    - ${BLOG_RPC_ENDPOINT}
  NonBlock: true
  Timeout: 5000
```

- [ ] **Step 5: 改写 `blog.yaml` 为 compose 可注入形式**

目标配置骨架：

```yaml
Name: blog.rpc
ListenOn: 0.0.0.0:${BLOG_RPC_PORT}
Mode: dev
Timeout: 5000
Health: true
Log:
  Mode: console
  Encoding: plain
  Path: runtime/blog-rpc/log
PgsqlConf:
  host: ${POSTGRES_HOST}
  port: ${POSTGRES_PORT}
  config: ${POSTGRES_CONFIG}
  dbname: ${POSTGRES_DB}
  username: ${POSTGRES_USER}
  password: ${POSTGRES_PASSWORD}
RedisConf:
  db: ${REDIS_DB}
  host: ${REDIS_HOST}
  port: "${REDIS_PORT}"
  password: ${REDIS_PASSWORD}
RabbitMQConf:
  host: ${RABBITMQ_HOST}
  port: "${RABBITMQ_PORT}"
  username: ${RABBITMQ_USER}
  password: ${RABBITMQ_PASSWORD}
EmailConf:
  host: ${EMAIL_HOST}
  port: ${EMAIL_PORT}
  username: ${EMAIL_USERNAME}
  password: ${EMAIL_PASSWORD}
  nickname: ${EMAIL_NICKNAME}
  bcc:
    - ${EMAIL_BCC}
ThirdPartyConf:
  blog-web:
    github:
      client_id: ${BLOG_WEB_GITHUB_CLIENT_ID}
      client_secret: ${BLOG_WEB_GITHUB_CLIENT_SECRET}
      redirect_uri: ${BLOG_WEB_GITHUB_REDIRECT_URI}
  admin-web:
    github:
      client_id: ${ADMIN_WEB_GITHUB_CLIENT_ID}
      client_secret: ${ADMIN_WEB_GITHUB_CLIENT_SECRET}
      redirect_uri: ${ADMIN_WEB_GITHUB_REDIRECT_URI}
```

- [ ] **Step 6: 新增 `.env.example`**

建议内容：

```dotenv
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_DB=blog-init
POSTGRES_USER=root
POSTGRES_PASSWORD=root
POSTGRES_CONFIG=sslmode=disable TimeZone=Asia/Shanghai

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest

BLOG_RPC_PORT=9999
BLOG_RPC_ENDPOINT=blog-rpc:9999

ADMIN_API_PORT=9091
ADMIN_API_DEVSERVER_PORT=6061

UPLOAD_PROVIDER=local
UPLOAD_LOCAL_DIR=runtime/resource
UPLOAD_LOCAL_BASE_URL=http://127.0.0.1:9091/static

QINIU_ZONE=
QINIU_ENDPOINT=
QINIU_ACCESS_KEY_ID=
QINIU_ACCESS_KEY_SECRET=
QINIU_BUCKET_NAME=
QINIU_BUCKET_URL=

EMAIL_HOST=smtp.qq.com
EMAIL_PORT=465
EMAIL_USERNAME=
EMAIL_PASSWORD=
EMAIL_NICKNAME=Yahei
EMAIL_BCC=

BLOG_WEB_GITHUB_CLIENT_ID=
BLOG_WEB_GITHUB_CLIENT_SECRET=
BLOG_WEB_GITHUB_REDIRECT_URI=http://127.0.0.1:9420/oauth/login/github

ADMIN_WEB_GITHUB_CLIENT_ID=
ADMIN_WEB_GITHUB_CLIENT_SECRET=
ADMIN_WEB_GITHUB_REDIRECT_URI=http://127.0.0.1:9420/oauth/login/github
```

- [ ] **Step 7: 验证入口仍可编译**

Run:

```bash
env GOCACHE=/tmp/go-build go build ./service/api/admin ./service/rpc/blog
```

Expected:

```text
command exits 0
```

- [ ] **Step 8: Commit**

```bash
git add service/api/admin/admin.go service/rpc/blog/blog.go service/api/admin/etc/admin-api.yaml service/rpc/blog/etc/blog.yaml .env.example
git commit -m "chore: env-enable admin and blog configs"
```

## Task 2: 重写本地上传配置与实现

**Files:**
- Modify: `pkg/oss/oss.go`
- Modify: `pkg/oss/local.go`
- Modify: `pkg/oss/qiniu.go`
- Modify: `service/api/admin/internal/svc/service_context.go`
- Modify: `service/api/admin/internal/plugins/plugins.go`
- Modify: `service/api/admin/infra/staticx/static.go`
- Test: `service/api/admin/internal/logic/upload/upload_logic_test.go`

- [ ] **Step 1: 先写一组失败测试，锁定本地上传语义**

在 [`service/api/admin/internal/logic/upload/upload_logic_test.go`](../service/api/admin/internal/logic/upload/upload_logic_test.go) 增加一个本地 uploader 语义测试，示例：

```go
func TestLocalUploaderSeparatesDiskPathAndPublicURL(t *testing.T) {
	uploader := oss.NewLocal("runtime/resource", "http://127.0.0.1:9091/static")
	got, err := uploader.UploadFile(strings.NewReader("hello"), "image", "a.txt")
	if err != nil {
		t.Fatalf("UploadFile returned error: %v", err)
	}
	if got != "http://127.0.0.1:9091/static/image/a.txt" {
		t.Fatalf("unexpected public url: %s", got)
	}
	if _, err := os.Stat("runtime/resource/image/a.txt"); err != nil {
		t.Fatalf("uploaded file not found: %v", err)
	}
}
```

- [ ] **Step 2: 扩展 `oss.Config`，显式支持 provider/local**

将 [`pkg/oss/oss.go`](../pkg/oss/oss.go) 中配置结构调整为：

```go
type Config struct {
	Provider        string `json:"provider" yaml:"provider"`
	LocalDir        string `json:"local-dir" yaml:"local-dir"`
	LocalBaseURL    string `json:"local-base-url" yaml:"local-base-url"`
	Zone            string `json:"zone" yaml:"zone"`
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret" yaml:"access-key-secret"`
	BucketName      string `json:"bucket-name" yaml:"bucket-name"`
	BucketUrl       string `json:"bucket-url" yaml:"bucket-url"`
}
```

- [ ] **Step 3: 修正 `Local` 上传器的路径语义**

把 [`pkg/oss/local.go`](../pkg/oss/local.go) 结构改成：

```go
type Local struct {
	rootDir  string
	baseURL  string
}
```

核心逻辑改成：

```go
target := filepath.Join(s.rootDir, filepath.FromSlash(key))
publicURL := strings.TrimRight(s.baseURL, "/") + "/" + key
```

删除时用磁盘路径：

```go
target := filepath.Join(s.rootDir, filepath.FromSlash(filepathValue))
```

列举时返回相对 `prefix` 的文件路径和 `baseURL`：

```go
FilePath: relPath,
FileUrl:  strings.TrimRight(s.baseURL, "/") + "/" + filepath.ToSlash(relPath),
```

构造函数改成：

```go
func NewLocal(rootDir, baseURL string) *Local
```

- [ ] **Step 4: 保持 `Qiniu` 路径兼容**

[`pkg/oss/qiniu.go`](../pkg/oss/qiniu.go) 只做字段兼容，不改行为。确保仍然使用：

```go
s.cfg.BucketUrl + "/" + key
```

- [ ] **Step 5: 在 `ServiceContext` 中按 provider 选择 uploader**

把 [`service/api/admin/internal/svc/service_context.go`](../service/api/admin/internal/svc/service_context.go) 里的：

```go
uploader := oss.NewQiniu(c.UploadConfig)
```

改成类似：

```go
var uploader oss.Uploader
switch strings.ToLower(strings.TrimSpace(c.UploadConfig.Provider)) {
case "", "qiniu":
	uploader = oss.NewQiniu(c.UploadConfig)
case "local":
	uploader = oss.NewLocal(c.UploadConfig.LocalDir, c.UploadConfig.LocalBaseURL)
default:
	panic(fmt.Sprintf("unsupported upload provider: %s", c.UploadConfig.Provider))
}
```

- [ ] **Step 6: 把静态资源路由接入启动链路**

在 [`service/api/admin/internal/plugins/plugins.go`](../service/api/admin/internal/plugins/plugins.go) 增加显式注册，例如：

```go
func RegisterPluginHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	staticx.RegisterLocalStatic(server, "/static/", "runtime/resource")

	var knife4jPrefix = "/admin-api/v1/swagger"
	server.AddRoutes(staticx.PrefixRoutes(knife4jPrefix, http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		knife4j.NewKnife4jPlugin(docs.Docs).Handler(knife4jPrefix).ServeHTTP(w, r)
	}))
}
```

并在 [`service/api/admin/infra/staticx/static.go`](../service/api/admin/infra/staticx/static.go) 增加一个小函数：

```go
func RegisterLocalStatic(server *rest.Server, prefix, dir string) {
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    path.Join(strings.TrimSuffix(prefix, "/"), ":file"),
		Handler: http.StripPrefix(prefix, http.FileServer(http.Dir(dir))).ServeHTTP,
	})
	server.AddRoutes(PrefixRoutes(prefix, http.MethodGet, http.StripPrefix(prefix, http.FileServer(http.Dir(dir))).ServeHTTP))
}
```

这里要注意：已有 `PrefixRoutes` 只能覆盖五层目录；如果不想继续沿用这个限制，就在同一任务里把本地静态路由改成更直接的匹配方式，但不要顺手扩展无关静态站点逻辑。

- [ ] **Step 7: 在 `admin.go` 里接入插件注册**

在 [`service/api/admin/admin.go`](../service/api/admin/admin.go) 中注册：

```go
plugins.RegisterPluginHandlers(server, ctx)
```

位置放在 `handler.RegisterHandlers(server, ctx)` 后面即可。

- [ ] **Step 8: 跑上传逻辑测试**

Run:

```bash
env GOCACHE=/tmp/go-build go test ./service/api/admin/internal/logic/upload -count=1
```

Expected:

```text
ok
```

- [ ] **Step 9: Commit**

```bash
git add pkg/oss/oss.go pkg/oss/local.go pkg/oss/qiniu.go service/api/admin/internal/svc/service_context.go service/api/admin/internal/plugins/plugins.go service/api/admin/infra/staticx/static.go service/api/admin/admin.go service/api/admin/internal/logic/upload/upload_logic_test.go
git commit -m "feat: support local upload for admin compose"
```

## Task 3: 补齐容器构建与 compose 编排

**Files:**
- Create: `docker/admin-api/Dockerfile`
- Create: `docker/blog-rpc/Dockerfile`
- Create: `docker-compose.yaml`
- Create: `.dockerignore`

- [ ] **Step 1: 新增 `.dockerignore`，收缩构建上下文**

建议内容：

```dockerignore
.git
.idea
.vscode
tmp
runtime
dist
node_modules
```

- [ ] **Step 2: 新增 `blog-rpc` Dockerfile**

`docker/blog-rpc/Dockerfile`：

```dockerfile
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/blog-rpc ./service/rpc/blog

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /out/blog-rpc /app/blog-rpc
COPY service/rpc/blog/etc/blog.yaml /app/etc/blog.yaml
CMD ["/app/blog-rpc", "-f", "/app/etc/blog.yaml"]
```

- [ ] **Step 3: 新增 `admin-api` Dockerfile**

`docker/admin-api/Dockerfile`：

```dockerfile
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/admin-api ./service/api/admin

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /out/admin-api /app/admin-api
COPY service/api/admin/etc/admin-api.yaml /app/etc/admin-api.yaml
RUN mkdir -p /app/runtime/resource
CMD ["/app/admin-api", "-f", "/app/etc/admin-api.yaml"]
```

- [ ] **Step 4: 新增 `docker-compose.yaml`**

建议起步版本：

```yaml
services:
  postgres:
    image: postgres:16
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      TZ: Asia/Shanghai
    ports:
      - "${POSTGRES_PORT}:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data
      - ./blog-init.sql:/docker-entrypoint-initdb.d/01-blog-init.sql:ro
      - ./seed.sql:/docker-entrypoint-initdb.d/02-seed.sql:ro

  redis:
    image: redis:7
    ports:
      - "${REDIS_PORT}:6379"

  rabbitmq:
    image: rabbitmq:3-management
    environment:
      RABBITMQ_DEFAULT_USER: ${RABBITMQ_USER}
      RABBITMQ_DEFAULT_PASS: ${RABBITMQ_PASSWORD}
    ports:
      - "${RABBITMQ_PORT}:5672"
      - "15672:15672"

  blog-rpc:
    build:
      context: .
      dockerfile: docker/blog-rpc/Dockerfile
    depends_on:
      - postgres
      - redis
      - rabbitmq
    env_file:
      - .env
    ports:
      - "${BLOG_RPC_PORT}:${BLOG_RPC_PORT}"

  admin-api:
    build:
      context: .
      dockerfile: docker/admin-api/Dockerfile
    depends_on:
      - blog-rpc
      - redis
    env_file:
      - .env
    ports:
      - "${ADMIN_API_PORT}:${ADMIN_API_PORT}"
      - "${ADMIN_API_DEVSERVER_PORT}:${ADMIN_API_DEVSERVER_PORT}"
    volumes:
      - admin-runtime:/app/runtime

volumes:
  postgres-data:
  admin-runtime:
```

- [ ] **Step 5: 快速检查 compose 语法**

Run:

```bash
docker compose config
```

Expected:

```text
compose renders successfully
```

- [ ] **Step 6: Commit**

```bash
git add .dockerignore docker/admin-api/Dockerfile docker/blog-rpc/Dockerfile docker-compose.yaml
git commit -m "chore: add local compose runtime"
```

## Task 4: 新增最小 `seed.sql`

**Files:**
- Create: `seed.sql`
- Test: `blog-init.sql`

- [ ] **Step 1: 固定一个最小管理员身份**

建议常量：

```sql
-- username: admin
-- password: 123456
-- user_id: admin-001
-- role_id: 1
```

这里的密码哈希不要写“运行时生成”，直接在 `seed.sql` 中落固定 bcrypt 值。

- [ ] **Step 2: 插入管理员用户**

建议骨架：

```sql
INSERT INTO t_user (
  user_id, username, password, nickname, avatar, email, phone, info, status, register_type, ip_address, ip_source
) VALUES (
  'admin-001',
  'admin',
  '$2a$10$DNSq5UwzRTn8Xco1c8.jku1gfz2YnL1vl4WjT5WASYljPPat7pYiO',
  'Administrator',
  '',
  'admin@example.com',
  '',
  '{}',
  0,
  'username',
  '127.0.0.1',
  'local'
) ON CONFLICT (user_id) DO NOTHING;
```

- [ ] **Step 3: 插入管理员角色与关联**

```sql
INSERT INTO t_role (id, parent_id, role_key, role_label, role_comment, is_default, status)
VALUES (1, 0, 'super_admin', 'Super Admin', 'compose seed role', true, 0)
ON CONFLICT (id) DO NOTHING;

INSERT INTO t_user_role (user_id, role_id)
VALUES ('admin-001', 1)
ON CONFLICT DO NOTHING;
```

- [ ] **Step 4: 插入最小菜单**

至少补一组能让 `get_user_menus` 返回非空数据的菜单，例如：

```sql
INSERT INTO t_menu (id, parent_id, path, name, component, redirect, type, title, icon, rank, perm, params, keep_alive, always_show, visible, status, extra)
VALUES
  (1, 0, '/system', 'System', '/system/index', '', '0', '系统管理', 'Setting', 1, 'system', '', false, true, true, false, '{}'::jsonb),
  (2, 1, '/system/account', 'Account', '/system/account/index', '', '1', '用户管理', 'User', 1, 'account:list', '', false, false, true, false, '{}'::jsonb)
ON CONFLICT (id) DO NOTHING;
```

- [ ] **Step 5: 插入最小 API**

至少补登录后首批受保护接口会访问到的 API 记录，例如：

```sql
INSERT INTO t_api (id, parent_id, name, path, method, traceable, status)
VALUES
  (1, 0, 'GetUserInfo', '/admin-api/v1/user/get_user_info', 'GET', false, false),
  (2, 0, 'GetUserMenus', '/admin-api/v1/user/get_user_menus', 'GET', false, false),
  (3, 0, 'GetUserRoles', '/admin-api/v1/user/get_user_roles', 'GET', false, false),
  (4, 0, 'GetUserApis', '/admin-api/v1/user/get_user_apis', 'GET', false, false),
  (5, 0, 'FindAccountList', '/admin-api/v1/account/find_account_list', 'POST', true, false)
ON CONFLICT (id) DO NOTHING;
```

- [ ] **Step 6: 插入角色菜单 / 角色接口关联**

```sql
INSERT INTO t_role_menu (role_id, menu_id)
VALUES (1, 1), (1, 2)
ON CONFLICT DO NOTHING;

INSERT INTO t_role_api (role_id, api_id)
VALUES (1, 1), (1, 2), (1, 3), (1, 4), (1, 5)
ON CONFLICT DO NOTHING;
```

- [ ] **Step 7: 用查询验证种子字段能命中现有模型查询**

对照以下查询来源逐项检查：

- [`service/model/t_menu_model.go`](../service/model/t_menu_model.go)
- [`service/model/t_api_model.go`](../service/model/t_api_model.go)
- [`service/model/t_role_model.go`](../service/model/t_role_model.go)

Run:

```bash
rg -n "FindByUserID|FindRolesByUserID" service/model/t_menu_model.go service/model/t_api_model.go service/model/t_role_model.go
```

Expected:

```text
queries join t_user_role / t_role_menu / t_role_api
```

- [ ] **Step 8: Commit**

```bash
git add seed.sql
git commit -m "feat: add local admin seed data"
```

## Task 5: 启动验证与最小联调验收

**Files:**
- Modify: `docker-compose.yaml`（仅在验证中发现必需问题时）
- Modify: `.env`（本地私有文件，不提交）

- [ ] **Step 1: 准备本地 `.env`**

Run:

```bash
cp .env.example .env
```

Expected:

```text
local .env created
```

- [ ] **Step 2: 首次启动编排**

Run:

```bash
docker compose up --build
```

Expected:

```text
postgres started
redis started
rabbitmq started
blog-rpc started
admin-api started
```

- [ ] **Step 3: 验证健康检查**

Run:

```bash
curl -s http://127.0.0.1:${ADMIN_API_PORT}/admin-api/v1/ping
```

Expected body shape:

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "env": "dev",
    "name": "admin-api"
  }
}
```

- [ ] **Step 4: 获取验证码**

Run:

```bash
curl -s -X POST http://127.0.0.1:${ADMIN_API_PORT}/admin-api/v1/get_captcha_code \
  -H 'Content-Type: application/json' \
  -d '{"width":120,"height":40}'
```

Expected:

```json
{
  "code": 0,
  "data": {
    "captcha_key": "...",
    "captcha_code": "...."
  }
}
```

注意：当前接口直接回传 `captcha_code`，所以本地验收可以直接复用，不需要 OCR。

- [ ] **Step 5: 登录拿 token**

Run:

```bash
LOGIN_RESP="$(curl -s -X POST http://127.0.0.1:${ADMIN_API_PORT}/admin-api/v1/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"admin","password":"123456"}')"
printf '%s\n' "$LOGIN_RESP"
```

Expected:

```json
{
  "code": 0,
  "data": {
    "user_id": "admin-001",
    "token": {
      "access_token": "...",
      "refresh_token": "..."
    }
  }
}
```

- [ ] **Step 6: 验证受保护接口**

Run:

```bash
UID="$(printf '%s' "$LOGIN_RESP" | sed -n 's/.*"user_id":"\([^"]*\)".*/\1/p')"
ACCESS_TOKEN="$(printf '%s' "$LOGIN_RESP" | sed -n 's/.*"access_token":"\([^"]*\)".*/\1/p')"
curl -s http://127.0.0.1:${ADMIN_API_PORT}/admin-api/v1/user/get_user_menus \
  -H "uid: ${UID}" \
  -H "authorization: ${ACCESS_TOKEN}"
```

Expected:

```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 1
      }
    ]
  }
}
```

- [ ] **Step 7: 验证一个基础 CRUD 列表接口**

Run:

```bash
curl -s -X POST http://127.0.0.1:${ADMIN_API_PORT}/admin-api/v1/account/find_account_list \
  -H 'Content-Type: application/json' \
  -H "uid: ${UID}" \
  -H "authorization: ${ACCESS_TOKEN}" \
  -d '{"page":1,"page_size":10}'
```

Expected:

```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "username": "admin"
      }
    ]
  }
}
```

- [ ] **Step 8: 验证上传链路**

Run:

```bash
curl -s -X POST http://127.0.0.1:${ADMIN_API_PORT}/admin-api/v1/upload/upload_file \
  -H "uid: ${UID}" \
  -H "authorization: ${ACCESS_TOKEN}" \
  -F "file=@README.md" \
  -F "file_path=test"
```

Expected:

```json
{
  "code": 0,
  "data": {
    "file_url": "http://127.0.0.1:9091/static/test/..."
  }
}
```

然后直接访问返回的 `file_url`，确认 200。

- [ ] **Step 9: 收尾修正 compose 细节**

如果启动中出现时序问题，只允许补最小必要改动：

- `depends_on`
- service restart policy
- mount path
- 端口映射

不要在这一步顺手加 nginx、etcd、前端、1Password injector。

- [ ] **Step 10: Commit**

```bash
git add docker-compose.yaml docker/admin-api/Dockerfile docker/blog-rpc/Dockerfile service/api/admin/etc/admin-api.yaml service/rpc/blog/etc/blog.yaml pkg/oss/oss.go pkg/oss/local.go pkg/oss/qiniu.go service/api/admin/internal/svc/service_context.go service/api/admin/internal/plugins/plugins.go service/api/admin/infra/staticx/static.go service/api/admin/admin.go seed.sql .env.example .dockerignore
git commit -m "feat: bootstrap local admin backend compose chain"
```

## 风险与处理

1. **`admin-api` 现在未注册插件 / 静态路由**
   - 这不是文档问题，是启动链路缺口；必须和本地上传一起修。

2. **`pkg/oss/local.go` 路径语义不完整**
   - 不能只把 `NewQiniu` 换成 `NewLocal`，否则返回 URL 不可访问，删除也可能删错路径。

3. **`seed.sql` 过少会导致 UI 登录后空白**
   - 至少要让 `get_user_info / get_user_menus / get_user_roles / get_user_apis` 能返回。

4. **验证码接口当前会返回明文验证码**
   - 这对本地最小链路是利好，可以降低联调成本；后续再讨论是否需要关闭。

5. **邮件 / 第三方登录配置为空时的启动行为**
   - 如果发现 `blog-rpc` 初始化阶段就读取这些配置并 panic，需要在实现时把对应依赖改成惰性使用；但不要提前过度改造。

## 完成标准

完成后必须同时满足：

1. `docker compose up --build` 能启动 `postgres / redis / rabbitmq / blog-rpc / admin-api`
2. `GET /admin-api/v1/ping` 返回成功
3. `POST /admin-api/v1/login` 能用 `admin / 123456` 登录成功
4. 携带 `uid + authorization` 后，`get_user_menus` 能返回非空
5. 至少一个后台列表 CRUD 接口能正常返回数据
6. 上传接口返回的本地 `file_url` 可以直接访问
7. 不要求邮件和第三方登录在这轮可用

## 建议执行顺序

1. Task 1 配置改造
2. Task 2 本地上传链路
3. Task 3 compose 编排
4. Task 4 seed 数据
5. Task 5 启动与验收
