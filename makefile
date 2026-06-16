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
MODEL_CORE_TABLES := t_user,t_user_oauth,t_role,t_menu,t_api,t_article,t_category,t_tag,t_talk,t_page,t_album,t_photo,t_friend,t_comment,t_message,t_system_notice,t_website_config
MODEL_RELATION_TABLES := t_role_api,t_role_menu,t_user_role
MODEL_LOG_TABLES := t_login_log,t_operation_log,t_visit_log,t_visit_daily_stats,t_file_log,t_visitor
BLOG_RPC_OUT := service/rpc/blog/internal/pb
BLOG_ZRPC_OUT := service/rpc/blog
ETC_DIR := service/rpc/blog/etc
BLOG_RPC_PROTO_DIR := service/rpc/blog/proto
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

goctl-api-admin:
	goctl api go \
		-api $(ADMIN_API_DIR)/proto/admin.api \
		-dir $(ADMIN_API_DIR) \
		--style go_zero \
		--type-group
	go mod tidy

goctl-rpc-blog:
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
	go mod tidy

goctl-model-core:
	goctl model pg datasource \
		--home="$(GOCTL_TEMPLATE_HOME)" \
		--url="$(PG_DSN)" \
		--table="$(MODEL_CORE_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--cache \
		--style="$(MODEL_STYLE)"

goctl-model-relation:
	goctl model pg datasource \
		--home="$(GOCTL_TEMPLATE_HOME)" \
		--url="$(PG_DSN)" \
		--table="$(MODEL_RELATION_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--style="$(MODEL_STYLE)"

goctl-model-log:
	goctl model pg datasource \
		--home="$(GOCTL_TEMPLATE_HOME)" \
		--url="$(PG_DSN)" \
		--table="$(MODEL_LOG_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--style="$(MODEL_STYLE)"

goctl-model-all: goctl-model-core goctl-model-relation goctl-model-log


goctl-api-admin-swagger:
	mkdir -p $(ADMIN_API_DOCS_DIR)
	goctl api swagger \
		--api $(ADMIN_API_DIR)/proto/admin.api \
		--dir $(ADMIN_API_DOCS_DIR) \
		--filename $(ADMIN_API_SWAGGER_FILENAME)

env-init:
	test -f $(ENV_FILE) || cp $(ENV_EXAMPLE_FILE) $(ENV_FILE)

compose-config: env-init
	$(DOCKER_COMPOSE) config

compose-build: env-init
	$(DOCKER_COMPOSE) build

compose-build-admin: env-init
	$(DOCKER_COMPOSE) build admin-api

compose-build-blog: env-init
	$(DOCKER_COMPOSE) build blog-rpc

compose-up: env-init
	$(DOCKER_COMPOSE) up -d

compose-up-build: env-init
	$(DOCKER_COMPOSE) up --build -d

compose-down:
	$(DOCKER_COMPOSE) down

compose-logs:
	$(DOCKER_COMPOSE) logs -f --tail=200

compose-ps:
	$(DOCKER_COMPOSE) ps

compose-restart:
	$(DOCKER_COMPOSE) restart

migrate-up: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" up

migrate-down: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" down 1

migrate-version: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" version

migrate-force: env-init
	@command -v $(MIGRATE) >/dev/null 2>&1 || { echo "migrate CLI not found. Install golang-migrate first."; exit 1; }
	@test -n "$(VERSION)" || { echo "VERSION is required, for example: make migrate-force VERSION=1"; exit 1; }
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(HOST_PG_URI)" force $(VERSION)

seed-bootstrap: env-init
	@command -v $(PSQL) >/dev/null 2>&1 || { echo "psql CLI not found. Install postgresql client first."; exit 1; }
	@for file in $(sort $(wildcard $(BOOTSTRAP_SEED_DIR)/*.sql)); do \
		echo "Applying bootstrap seed $$file"; \
		$(PSQL) "$(HOST_PG_PSQL_DSN)" -v ON_ERROR_STOP=1 -f "$$file" || exit $$?; \
	done
