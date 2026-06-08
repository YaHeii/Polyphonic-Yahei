# lib/pq 到 pgx 的 model 重构执行清单

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use markdown checkboxes for tracking.

**Goal:** 彻底移除仓库中的 `github.com/lib/pq` 依赖，并让 `goctl model pg datasource` 的生成链路稳定输出 `pgx` 兼容的原生切片类型。

**Architecture:** 本次重构不修改数据库 schema，只处理 Go 侧的类型映射、参数编码和代码生成模板。生成链路通过仓库内自定义 `goctl` SQL model 模板接管，业务层统一收敛到原生切片，不向外扩散 `pgtype`。

**Tech Stack:** Go, goctl 1.10.1, go-zero sqlx, PostgreSQL, pgx/v5 stdlib

**Status:** 已完成（2026-06-09）。

## 最终结果

- `makefile` 已通过 `GOCTL_TEMPLATE_HOME := ./.goctl` 接入仓库内模板，并在三处 `goctl model pg datasource` 命令上统一增加 `--home="$(GOCTL_TEMPLATE_HOME)"`
- `.goctl/model/field.tpl` 已将 goctl 上游内部类型名 `pq.StringArray` / `pq.Int64Array` / `pq.Float64Array` 映射为原生切片；`import.tpl` 与 `import-no-cache.tpl` 已去掉 `github.com/lib/pq`
- `service/model/*.go` generated file 与 handwritten file、`service/rpc/blog/internal/logic/*` 中的数组字段与数组参数均已迁移为原生切片
- `go.mod` / `go.sum` 中已移除 `github.com/lib/pq`
- 最终验证已通过：
  - `make goctl-model-all`
  - `env GOCACHE=/tmp/go-build go build ./service/model/... ./service/rpc/blog/...`
  - `env GOCACHE=/tmp/go-build go test ./service/api/admin/infra/responsex -count=1`
- 当前代码路径中，`pq.*Array` 只保留在 `.goctl/model/field.tpl` 的字符串判断里，用于兼容 goctl v1.10.1 上游 converter 暴露的内部类型名；这不是项目运行时依赖

---

## 文件结构与职责

### 1. 生成链路文件

- `makefile`
  - 接入仓库内自定义模板目录
  - 确保 `goctl model pg datasource` regenerate 稳定复现

- `.goctl/model/types.tpl`
  - 生成 model struct 外壳
  - 不直接决定字段类型，但属于模板链路必要组成

- `.goctl/model/field.tpl`
  - 输出单个 struct 字段
  - 上游默认模板直接输出 `{{.type}}`

- `.goctl/model/import.tpl`
  - 控制 generated model 的 import
  - 上游默认模板通过 `.containsPQ` 条件引入 `github.com/lib/pq`

- `.goctl/model/import-no-cache.tpl`
  - 无缓存版本 import 模板
  - 需要与 `import.tpl` 保持一致

### 2. generated model 文件

- `service/model/t_article_model_gen.go`
- `service/model/t_talk_model_gen.go`

这两份文件是第一阶段验证对象，成功后再扩散到其他 generated model。

### 3. handwritten model 文件

- `service/model/t_article_model.go`
- `service/model/t_talk_model.go`
- `service/model/t_category_model.go`
- `service/model/t_comment_model.go`
- `service/model/t_message_model.go`
- `service/model/t_photo_model.go`
- `service/model/t_system_notice_model.go`
- `service/model/t_tag_model.go`

这些文件承担：

- 数组字段默认值处理
- `any($1)` 查询参数传递
- 统计与查询聚合逻辑

### 4. query helper / adapter 文件

- `service/model/permission_query_helper.go`
- `service/model/resource_query_helper.go`
- `service/model/social_query_helper.go`
- `service/model/syslog_query_helper.go`

这些文件承担：

- 将 `?` 条件转换为 PostgreSQL 占位符
- 迁移前通过 `pq.Array(...)` 做数组参数编码

### 5. RPC logic 文件

- `service/rpc/blog/internal/logic/articlerpc/deletes_tag_logic.go`
- `service/rpc/blog/internal/logic/articlerpc/helper.go`
- `service/rpc/blog/internal/logic/socialrpc/helper.go`

这些文件承担：

- 构造 model 输入
- 迁移前直接使用 `pq.StringArray(...)` 或 `pq.Array(...)`

## 执行策略

采用四段式执行：

1. 模板链路验证
2. 核心表迁移验证
3. 全仓 handwritten 替换
4. 依赖清理与 regenerate 回归

每一段都要求：

- 先做最小改动
- 先在局部范围验证
- 再扩大替换范围

## Task 1: 建立 goctl 模板接管点

**Files:**
- Modify: `.goctl/model/types.tpl`
- Modify: `.goctl/model/field.tpl`
- Modify: `.goctl/model/import.tpl`
- Modify: `.goctl/model/import-no-cache.tpl`
- Modify: `makefile`

- [x] **Step 1: 复制 goctl 默认 SQL model 模板到仓库内**

复制来源：

- `/root/go/pkg/mod/github.com/zeromicro/go-zero/tools/goctl@v1.10.1/model/sql/template/tpl/types.tpl`
- `/root/go/pkg/mod/github.com/zeromicro/go-zero/tools/goctl@v1.10.1/model/sql/template/tpl/field.tpl`
- `/root/go/pkg/mod/github.com/zeromicro/go-zero/tools/goctl@v1.10.1/model/sql/template/tpl/import.tpl`
- `/root/go/pkg/mod/github.com/zeromicro/go-zero/tools/goctl@v1.10.1/model/sql/template/tpl/import-no-cache.tpl`

目标目录：

```text
.goctl/model/
```

- [x] **Step 2: 在 `makefile` 中增加模板目录变量**

在现有变量区增加：

```make
GOCTL_TEMPLATE_HOME := ./.goctl
```

- [x] **Step 3: 给所有 `goctl model pg datasource` 命令增加 `--home`**

目标命令：

- `goctl-model-core`
- `goctl-model-relation`
- `goctl-model-log`

统一增加：

```make
--home="$(GOCTL_TEMPLATE_HOME)"
```

- [x] **Step 4: 验证模板接管已生效**

Run:

```bash
rg -n -- '--home=' makefile
```

Expected:

- `goctl model pg datasource` 三处命令都带有 `--home="$(GOCTL_TEMPLATE_HOME)"`

- [x] **Step 5: Commit**

```bash
git add makefile .goctl/model
git commit -m "chore: wire local goctl model templates"
```

## Task 2: 验证模板能否控制数组字段输出

**Files:**
- Modify: `.goctl/model/field.tpl`
- Modify: `.goctl/model/import.tpl`
- Modify: `.goctl/model/import-no-cache.tpl`
- Regenerate: `service/model/t_article_model_gen.go`
- Regenerate: `service/model/t_talk_model_gen.go`

- [x] **Step 1: 检查模板上下文是否只暴露 `.type`**

当前 `field.tpl` 默认内容：

```tpl
{{.name}} {{.type}} {{.tag}} {{if .hasComment}}// {{.comment}}{{end}}
```

判断标准：

- 若模板只能拿到 `.type`，则后续需要基于 `.type` 条件替换
- 若模板还能拿到列原始类型信息，则优先按原始类型判断

- [x] **Step 2: 在模板中为数组类型增加受控输出**

目标输出策略：

- `pq.StringArray` -> `[]string`

如果模板只能看到 `.type`，则先做最小兼容判断，确保 `text[]` 至少能转成 `[]string`。

- [x] **Step 3: 从 import 模板中去掉 `lib/pq` 引入**

修改：

- `import.tpl`
- `import-no-cache.tpl`

要求：

- 生成结果中不再出现 `github.com/lib/pq`

- [x] **Step 4: 重新生成 `t_article` / `t_talk` 做第一轮验证**

Run:

```bash
goctl model pg datasource --home="./.goctl" --url="$(PG_DSN)" --table="t_article,t_talk" --dir="./service/model" --cache --style="go_zero"
```

Expected:

- `service/model/t_article_model_gen.go` 中 `Tags` 为 `[]string`
- `service/model/t_talk_model_gen.go` 中 `Images` 为 `[]string`
- generated import 中无 `github.com/lib/pq`

- [x] **Step 5: Commit**

```bash
git add .goctl/model service/model/t_article_model_gen.go service/model/t_talk_model_gen.go
git commit -m "chore: make goctl pg model arrays use slices"
```

## Task 3: 迁移核心表 handwritten 代码

**Files:**
- Modify: `service/model/t_article_model.go`
- Modify: `service/model/t_talk_model.go`
- Modify: `service/rpc/blog/internal/logic/articlerpc/helper.go`
- Modify: `service/rpc/blog/internal/logic/socialrpc/helper.go`

- [x] **Step 1: 将 `pq.StringArray` 构造改为原生切片**

目标替换：

- `pq.StringArray{}` -> `[]string{}`
- `pq.StringArray(in.ImgList)` -> `append([]string(nil), in.ImgList...)`

- [x] **Step 2: 将 `pq.Array(...)` 参数改为直接切片参数**

目标替换：

- `pq.Array(ids)` -> `ids`
- `pq.Array(names)` -> `names`

注意：

- 保持 SQL 仍使用 `any($1)` / `@>` 等 PostgreSQL 原生数组写法

- [x] **Step 3: 清理 import**

要求：

- 上述四个文件中不再引用 `github.com/lib/pq`

- [x] **Step 4: 运行局部编译验证**

Run:

```bash
env GOCACHE=/tmp/go-build go build ./service/model ./service/rpc/blog/internal/logic/articlerpc ./service/rpc/blog/internal/logic/socialrpc
```

Expected:

- 编译通过

- [x] **Step 5: Commit**

```bash
git add service/model/t_article_model.go service/model/t_talk_model.go service/rpc/blog/internal/logic/articlerpc/helper.go service/rpc/blog/internal/logic/socialrpc/helper.go
git commit -m "refactor: migrate article and talk arrays to slices"
```

## Task 4: 迁移 query helper 与其他 handwritten model

**Files:**
- Modify: `service/model/permission_query_helper.go`
- Modify: `service/model/resource_query_helper.go`
- Modify: `service/model/social_query_helper.go`
- Modify: `service/model/syslog_query_helper.go`
- Modify: `service/model/t_category_model.go`
- Modify: `service/model/t_comment_model.go`
- Modify: `service/model/t_message_model.go`
- Modify: `service/model/t_photo_model.go`
- Modify: `service/model/t_system_notice_model.go`
- Modify: `service/model/t_tag_model.go`
- Modify: `service/rpc/blog/internal/logic/articlerpc/deletes_tag_logic.go`

- [x] **Step 1: 逐文件替换 `pq.Array(...)`**

统一策略：

- 所有 `[]int64` / `[]string` 直接作为参数传入
- 不修改 SQL 语义，只处理 Go 侧参数编码

- [x] **Step 2: 清理 `github.com/lib/pq` import**

要求：

- 本任务涉及文件中不再出现 `lib/pq`

- [x] **Step 3: 运行局部搜索确认 handwritten 使用面归零**

Run:

```bash
rg -n "pq\\.Array\\(|pq\\.StringArray|github.com/lib/pq" service/model service/rpc/blog/internal/logic -g'*.go'
```

Expected:

- 只允许剩余 generated 文件命中；如果 handwritten 还有命中，继续清理

- [x] **Step 4: 运行 model 与 RPC 编译验证**

Run:

```bash
env GOCACHE=/tmp/go-build go build ./service/model/... ./service/rpc/blog/...
```

Expected:

- 编译通过

- [x] **Step 5: Commit**

```bash
git add service/model service/rpc/blog/internal/logic/articlerpc/deletes_tag_logic.go
git commit -m "refactor: remove pq array helpers from handwritten code"
```

## Task 5: 全量 regenerate 并清理依赖

**Files:**
- Regenerate: `service/model/*.go`
- Modify: `go.mod`
- Modify: `go.sum`

- [x] **Step 1: 执行全量 model regenerate**

Run:

```bash
make goctl-model-all
```

Expected:

- 所有 model generated file 使用仓库模板生成
- 不重新引入 `lib/pq`

- [x] **Step 2: 运行全仓搜索确认 `lib/pq` 使用归零**

Run:

```bash
rg -n "github.com/lib/pq|pq\\.Array\\(|pq\\.StringArray" service pkg go.mod go.sum
```

Expected:

- 无结果

- [x] **Step 3: 执行依赖整理**

Run:

```bash
env GOCACHE=/tmp/go-build go mod tidy
```

Expected:

- `go.mod` / `go.sum` 中不再包含 `github.com/lib/pq`

- [x] **Step 4: 运行回归编译验证**

Run:

```bash
env GOCACHE=/tmp/go-build go build ./service/model/... ./service/rpc/blog/...
```

Expected:

- 编译通过

- [x] **Step 5: Commit**

```bash
git add go.mod go.sum service/model makefile .goctl/model
git commit -m "refactor: remove libpq from postgres model stack"
```

## Task 6: 最终验证与文档回填

**Files:**
- Modify: `docs/libpq-to-pgx-model-refactor-plan.md`
- Modify: `docs/libpq-to-pgx-model-refactor-checklist.md`

- [x] **Step 1: 记录最终模板接入方式**

补充内容：

- 模板目录位置
- `makefile` 中的 `--home` 接法
- regenerate 命令

- [x] **Step 2: 记录迁移后的类型约定**

补充内容：

- PostgreSQL array 列统一使用原生切片
- 不再允许在业务层使用 `pq.StringArray` / `pq.Array`

- [x] **Step 3: 运行最终验证命令**

Run:

```bash
env GOCACHE=/tmp/go-build go build ./service/model/... ./service/rpc/blog/...
env GOCACHE=/tmp/go-build go test ./service/api/admin/infra/responsex -count=1
```

Expected:

- 编译通过
- 不因本次重构影响已改过的局部包

- [x] **Step 4: Commit**

```bash
git add docs/libpq-to-pgx-model-refactor-plan.md docs/libpq-to-pgx-model-refactor-checklist.md
git commit -m "docs: finalize libpq to pgx migration guidance"
```

## 补充说明

### 关于 `.sql` 真相源

本计划默认不修改 SQL 真相源，因为：

1. 目标是更换 Go 侧类型体系
2. PostgreSQL array 列语义不变
3. 当前问题是生成器和参数编码，而不是 schema

### 关于风险控制

每个任务都要求：

1. 先局部验证，再扩大范围
2. generated 改动必须通过模板链路落地
3. handwritten 改动只处理 `lib/pq` 相关依赖，不顺手改邻近逻辑
