ADMIN_API_DIR := service/api/admin
ADMIN_API_DOCS_DIR := $(ADMIN_API_DIR)/internal/docs
ADMIN_API_SWAGGER_FILENAME := swagger
MODEL_DIR := ./service/model
GOCTL_TEMPLATE_HOME := ./.goctl
ENV_FILE ?= .env
ENV_EXAMPLE_FILE ?= .env.example
ENV_SOURCE := $(if $(wildcard $(ENV_FILE)),$(ENV_FILE),$(ENV_EXAMPLE_FILE))
-include $(ENV_SOURCE)
COMPOSE_FILE ?= docker-compose.yaml
DOCKER_COMPOSE := docker compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE)
PG_DSN := postgres://root:root@127.0.0.1:5432/blog-init?sslmode=disable
MODEL_STYLE := go_zero
MODEL_CORE_TABLES := t_user,t_user_oauth,t_role,t_menu,t_api,t_article,t_category,t_talk,t_album,t_photo,t_friend,t_comment,t_message,t_system_notice,t_website_config
MODEL_RELATION_TABLES := t_role_api,t_role_menu
MODEL_LOG_TABLES := t_login_log,t_operation_log,t_visit_log,t_visit_daily_stats,t_file_log,t_visitor
BLOG_RPC_OUT := service/rpc/blog/internal/pb
BLOG_ZRPC_OUT := service/rpc/blog
ETC_DIR := service/rpc/blog/etc
BLOG_RPC_PROTO_DIR := service/rpc/blog/proto
GO_BIN_DIR := /usr/local/go/bin
GOCTL_BIN_DIR := /root/go/bin
GO_BUILD_CACHE ?= /tmp/go-build
MIGRATE_DIR := service/db/migrations
BOOTSTRAP_SEED_DIR := service/db/seeds/bootstrap
DB_CLI_HOST ?= 127.0.0.1
MIGRATE ?= migrate
PSQL ?= psql
empty :=
space := $(empty) $(empty)
DB_SSLMODE := $(or $(patsubst sslmode=%,%,$(filter sslmode=%,$(POSTGRES_CONFIG))),disable)
HOST_PG_URI := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(DB_CLI_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(DB_SSLMODE)
HOST_PG_PSQL_DSN := host=$(DB_CLI_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=$(DB_SSLMODE)


.PHONY: goctl-api-admin goctl-api-admin-swagger goctl-api-admin-clean-generated goctl-api-admin-reset goctl-model goctl-model-core goctl-model-relation goctl-model-log goctl-model-all goctl-rpc-blog
.PHONY: env-init compose-config compose-build compose-build-admin compose-build-blog compose-up compose-up-build compose-down compose-logs compose-ps compose-restart
.PHONY: migrate-up migrate-down migrate-version migrate-force seed-bootstrap

# 根据 admin.api 真相源生成后台 API 服务代码。
goctl-api-admin:
	goctl api go \
		-api $(ADMIN_API_DIR)/proto/admin.api \
		-dir $(ADMIN_API_DIR) \
		--style go_zero \
		--type-group
	go mod tidy

# 根据各领域 proto 文件生成 blog RPC 的 protobuf 与 zrpc 代码。
goctl-rpc-blog:
	export PATH="$(GOCTL_BIN_DIR):$(GO_BIN_DIR):$$PATH"; \
	export GOCACHE="$(GO_BUILD_CACHE)"; \
	for file in "$(BLOG_RPC_PROTO_DIR)"/blog/*.proto; do \
		if [ -f "$$file" ]; then \
			goctl rpc protoc "$$file" \
			-I $(BLOG_RPC_PROTO_DIR) \
			--go_out=$(BLOG_RPC_OUT) \
			--go-grpc_out=$(BLOG_RPC_OUT) \
			--zrpc_out=$(BLOG_ZRPC_OUT) \
			--style go_zero \
			-m; \
		fi; \
	done; \
	rm -f "$(ETC_DIR)"/*rpc.yaml; \
	rm -f "$(BLOG_ZRPC_OUT)"/*rpc.go
	PATH="$(GOCTL_BIN_DIR):$(GO_BIN_DIR):$$PATH" GOCACHE="$(GO_BUILD_CACHE)" go mod tidy

# 为核心业务表生成 goctl model 代码。
goctl-model-core:
	goctl model pg datasource \
		--home="$(GOCTL_TEMPLATE_HOME)" \
		--url="$(PG_DSN)" \
		--table="$(MODEL_CORE_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--cache \
		--style="$(MODEL_STYLE)"

# 为 RBAC 关联表生成 goctl model 代码。
goctl-model-relation:
	goctl model pg datasource \
		--home="$(GOCTL_TEMPLATE_HOME)" \
		--url="$(PG_DSN)" \
		--table="$(MODEL_RELATION_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--style="$(MODEL_STYLE)"

# 为日志与统计类表生成 goctl model 代码。
goctl-model-log:
	goctl model pg datasource \
		--home="$(GOCTL_TEMPLATE_HOME)" \
		--url="$(PG_DSN)" \
		--table="$(MODEL_LOG_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--style="$(MODEL_STYLE)"

# 一次性执行全部 model 生成目标。
goctl-model-all: goctl-model-core goctl-model-relation goctl-model-log


# 根据 admin.api 真相源生成 Swagger JSON 文档。
goctl-api-admin-swagger:
	mkdir -p $(ADMIN_API_DOCS_DIR)
	goctl api swagger \
		--api $(ADMIN_API_DIR)/proto/admin.api \
		--dir $(ADMIN_API_DOCS_DIR) \
		--filename $(ADMIN_API_SWAGGER_FILENAME)

# 本地缺少 .env 时，从示例文件复制一份。
env-init:
	test -f $(ENV_FILE) || cp $(ENV_EXAMPLE_FILE) $(ENV_FILE)

# 渲染并展示最终生效的 Docker Compose 配置。
compose-config: env-init
	$(DOCKER_COMPOSE) config

# 构建 Compose 中定义的全部服务镜像。
compose-build: env-init
	$(DOCKER_COMPOSE) build

# 仅构建 admin API 服务镜像。
compose-build-admin: env-init
	$(DOCKER_COMPOSE) build admin-api

# 仅构建 blog RPC 服务镜像。
compose-build-blog: env-init
	$(DOCKER_COMPOSE) build blog-rpc

# 以后台模式启动 Compose 服务栈。
compose-up: env-init
	$(DOCKER_COMPOSE) up -d

# 先重建镜像，再以后台模式启动 Compose 服务栈。
compose-up-build: env-init
	$(DOCKER_COMPOSE) up --build -d

# 停止并移除 Compose 服务栈。
compose-down:
	$(DOCKER_COMPOSE) down

# 持续跟踪 Compose 服务栈最近的日志输出。
compose-logs:
	$(DOCKER_COMPOSE) logs -f --tail=200

# 查看 Compose 服务栈当前容器状态。
compose-ps:
	$(DOCKER_COMPOSE) ps

# 重启 Compose 服务栈中的服务。
compose-restart:
	$(DOCKER_COMPOSE) restart

# 对目标数据库执行全部待应用的迁移。
migrate-up: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" up

# 回滚最近一步数据库迁移。
migrate-down: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" down 1

# 输出当前迁移版本及 dirty 状态。
migrate-version: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" version

# 在数据库迁移 dirty 时强制修正版本标记。
migrate-force: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	@test -n "$(VERSION)" || { echo "VERSION is required, for example: make migrate-force VERSION=1"; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" force $(VERSION)

# 按文件名顺序执行 bootstrap seed SQL。
seed-bootstrap: env-init
	@command -v $(PSQL) >/dev/null 2>&1 || { echo "psql CLI not found. Install postgresql client first."; exit 1; }
	@for file in $(sort $(wildcard $(BOOTSTRAP_SEED_DIR)/*.sql)); do \
		echo "Applying bootstrap seed $$file"; \
		$(PSQL) "$(HOST_PG_PSQL_DSN)" -v ON_ERROR_STOP=1 -f "$$file" || exit $$?; \
	done

# 构建本地 admin 和 blog 二进制文件。
go-build: env-init
	go build ./bin/admin
	go build ./bin/blog
