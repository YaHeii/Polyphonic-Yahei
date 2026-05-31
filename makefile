ADMIN_API_DIR := service/api/admin
MODEL_DIR := ./service/model
PG_DSN := postgres://root:root@127.0.0.1:5432/blog-init?sslmode=disable
MODEL_STYLE := go_zero
MODEL_CORE_TABLES := t_user,t_user_oauth,t_role,t_menu,t_api,t_article,t_category,t_tag,t_talk,t_page,t_album,t_photo,t_friend,t_comment,t_message,t_system_notice,t_website_config
MODEL_RELATION_TABLES := t_role_api,t_role_menu,t_user_role
MODEL_LOG_TABLES := t_login_log,t_operation_log,t_visit_log,t_visit_daily_stats,t_file_log,t_visitor

.PHONY: goctl-api-admin goctl-api-admin-clean-generated goctl-api-admin-reset goctl-model goctl-model-core goctl-model-relation goctl-model-log goctl-model-all

goctl-api-admin:
	goctl api go \
		-api $(ADMIN_API_DIR)/proto/admin.api \
		-dir $(ADMIN_API_DIR) \
		--style go_zero \
		--type-group
	go mod tidy

goctl-api-admin-clean-generated:
	rm -rf $(ADMIN_API_DIR)/etc
	rm -f $(ADMIN_API_DIR)/admin.go

goctl-api-admin-reset:
	rm -rf $(ADMIN_API_DIR)/internal
	rm -rf $(ADMIN_API_DIR)/etc
	rm -f $(ADMIN_API_DIR)/admin.go

goctl-model: goctl-model-core

goctl-model-core:
	goctl model pg datasource \
		--url="$(PG_DSN)" \
		--table="$(MODEL_CORE_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--cache \
		--style="$(MODEL_STYLE)"

goctl-model-relation:
	goctl model pg datasource \
		--url="$(PG_DSN)" \
		--table="$(MODEL_RELATION_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--style="$(MODEL_STYLE)"

goctl-model-log:
	goctl model pg datasource \
		--url="$(PG_DSN)" \
		--table="$(MODEL_LOG_TABLES)" \
		--dir="$(MODEL_DIR)" \
		--style="$(MODEL_STYLE)"

goctl-model-all: goctl-model-core goctl-model-relation goctl-model-log
