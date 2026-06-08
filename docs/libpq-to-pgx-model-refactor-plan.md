# lib/pq 到 pgx 的 model 重构计划

## 背景

当前仓库的 PostgreSQL 运行时驱动已经部分切到 `pgx`，例如 [`service/model/vars.go`](../service/model/vars.go) 中已经使用：

```go
_ "github.com/jackc/pgx/v5/stdlib"
```

但项目中仍然大量依赖 `github.com/lib/pq`，主要体现在两类场景：

1. model generated file 中的数组字段类型，例如 `pq.StringArray`
2. handwritten model / RPC logic 中的数组参数编码，例如 `pq.Array(...)`

因此，本次重构目标不是“切换数据库 driver”，而是：

- 去掉项目中的 `lib/pq` 类型依赖
- 去掉项目中的 `lib/pq` 参数编码依赖
- 让 `goctl model pg datasource` 的生成链路在 regenerate 后也不再引入 `lib/pq`

## 目标

完成后应满足以下状态：

1. `go.mod` / `go.sum` 不再依赖 `github.com/lib/pq`
2. `service/model/**/*` 中不再出现 `pq.StringArray`
3. `service/model/**/*` 和 `service/rpc/blog/internal/logic/**/*` 中不再出现 `pq.Array(...)`
4. `make goctl-model-*` 执行后，生成结果不会重新引入 `lib/pq`
5. PostgreSQL array 字段统一使用 `pgx` 兼容的原生 Go 类型体系

## 非目标

以下事项不在本次重构范围内：

1. 不修改数据库 schema
2. 不把 PostgreSQL array 列改为 `jsonb` 或关系表
3. 不重写 query builder
4. 不把 `pgtype.Array[T]` / `pgtype.FlatArray[T]` 扩散到业务层
5. 不顺手做无关的 model 重构

## 现状分析

### 1. 运行时驱动层

当前运行时连接 PostgreSQL 的 driver 注册已经通过 `pgx stdlib` 完成，因此 `lib/pq` 已经不是驱动层依赖。

### 2. model 类型体系

当前 `lib/pq` 主要被用于：

1. 数组字段类型
2. 数组参数编码

典型例子：

- [`service/model/t_article_model_gen.go`](../service/model/t_article_model_gen.go)
- [`service/model/t_talk_model_gen.go`](../service/model/t_talk_model_gen.go)
- [`service/model/t_article_model.go`](../service/model/t_article_model.go)
- [`service/model/permission_query_helper.go`](../service/model/permission_query_helper.go)
- [`service/rpc/blog/internal/logic/socialrpc/helper.go`](../service/rpc/blog/internal/logic/socialrpc/helper.go)

### 3. 生成链路

当前 model 通过 [`makefile`](../makefile) 中的 `goctl model pg datasource` 生成。

如果只手改生成产物而不处理模板链路，则下次 regenerate 后 `lib/pq` 会再次被带回。

因此，本次重构必须同时处理：

1. model 类型体系
2. `goctl` 模板生成链路

## 目标类型策略

本次重构建议统一使用原生切片，而不是将 `pgtype` 类型暴露到业务层。

推荐映射：

- `text[]` -> `[]string`
- `bigint[]` -> `[]int64`
- 其他常见 PG array -> 对应原生切片

不推荐的方案：

- `pq.StringArray` -> `pgtype.FlatArray[string]`
- `pq.StringArray` -> `pgtype.Array[string]`

原因：

1. 会把 `pgx` 类型体系扩散到业务层
2. 会增加 handwritten code 的转换成本
3. 不符合当前仓库的简洁边界

## 关键设计决策

### 决策 1：是否修改 `.sql` 真相源

默认 **不修改**。

原因：

1. 本次改的是 Go 类型映射和参数编码方式
2. 数据库列语义保持不变
3. PostgreSQL array 列本身与 `pgx` 完全兼容

只有在验证阶段证明原生切片在当前 `pgx stdlib + go-zero sqlx` 组合下无法稳定读写时，才重新讨论是否需要调整 SQL 或 schema。

### 决策 2：是否通过 goctl 模板解决 regenerate 回滚

是。

推荐做法：

1. 在仓库内维护自定义 `goctl` SQL model 模板目录
2. 在 `makefile` 中通过 `--home` 显式接入

不推荐仅通过“生成后 patch”解决，因为它会增加隐藏流程，并降低 regenerate 的可维护性。

### 决策 3：是否一次性全量迁移

否。

推荐按阶段推进：

1. 先验证单表和核心数组类型
2. 确认链路稳定后再批量迁移

## 实施阶段

## 阶段一：梳理并冻结当前使用面

### 目标

在动模板前，明确当前所有 `lib/pq` 使用点，并将其分组为：

1. generated file
2. handwritten model
3. handwritten RPC logic
4. query helper / parameter adapter

### 操作

1. 搜索 `github.com/lib/pq`
2. 搜索 `pq.StringArray`
3. 搜索 `pq.Array(`
4. 将命中结果按“generated / handwritten”分类

### 验证

形成完整迁移清单，避免后期遗漏导致 `go mod tidy` 失败。

## 阶段二：建立自定义 goctl 模板目录

### 目标

让 `goctl model pg datasource` 生成结果不再默认输出 `lib/pq` 相关代码。

### 操作

1. 从 `goctl` 默认 SQL model 模板复制出仓库内模板目录
2. 建议目录位置：

```text
hack/goctl-template/model/sql/template/tpl/
```

3. 重点修改以下模板：
   - `types.tpl`
   - `import.tpl`
   - `import-no-cache.tpl`
   - 必要时 `field.tpl`

### 预期改动

1. 数组字段输出为原生切片
2. 删除 `github.com/lib/pq` import

### 验证

以单表生成结果验证：

1. `t_article`
2. `t_talk`

确认生成 struct 中：

- `Tags []string`
- `Images []string`

而不是 `pq.StringArray`

## 阶段三：接入 makefile 生成链路

### 目标

让仓库内模板成为 model 生成的唯一稳定来源。

### 操作

修改 [`makefile`](../makefile)：

1. 增加模板目录变量，例如：

```make
GOCTL_TEMPLATE_HOME := ./hack/goctl-template
```

2. 给 `goctl model pg datasource` 增加：

```make
--home="$(GOCTL_TEMPLATE_HOME)"
```

### 验证

执行：

1. `goctl-model-core`
2. `goctl-model-relation`
3. `goctl-model-log`

确认 regenerate 后不会重新出现 `lib/pq`

## 阶段四：小范围迁移核心表

### 目标

先验证最典型的数组字段和数组参数链路。

### 范围

建议先迁移：

1. `t_article`
2. `t_talk`

### 涉及文件

generated：

- [`service/model/t_article_model_gen.go`](../service/model/t_article_model_gen.go)
- [`service/model/t_talk_model_gen.go`](../service/model/t_talk_model_gen.go)

handwritten：

- [`service/model/t_article_model.go`](../service/model/t_article_model.go)
- [`service/model/t_talk_model.go`](../service/model/t_talk_model.go)
- [`service/rpc/blog/internal/logic/socialrpc/helper.go`](../service/rpc/blog/internal/logic/socialrpc/helper.go)
- [`service/rpc/blog/internal/logic/articlerpc/helper.go`](../service/rpc/blog/internal/logic/articlerpc/helper.go)

### 操作

1. `pq.StringArray` -> `[]string`
2. `pq.StringArray{}` -> `[]string{}`
3. `pq.StringArray(in.ImgList)` -> `append([]string(nil), in.ImgList...)` 或等价写法
4. `pq.Array(ids/names)` -> 直接传 `ids/names`

### 验证

1. 编译 `service/model/...`
2. 编译 `service/rpc/blog/...`
3. 重点验证 array 字段插入、查询、更新链路

## 阶段五：批量迁移所有 handwritten 参数编码

### 目标

去掉所有 `pq.Array(...)` 依赖。

### 重点区域

- `service/model/*_query_helper.go`
- `service/model/t_*_model.go`
- `service/rpc/blog/internal/logic/articlerpc/*`
- `service/rpc/blog/internal/logic/socialrpc/*`

### 操作

1. 所有 `pq.Array(v)` 统一改为直接切片参数
2. 保持 SQL 仍使用 `= any($1)` 或其他 PG 原生数组写法
3. 必要时仅做最小显式 cast 修正

### 验证

1. `rg -n "pq\\.Array\\(" ...` 归零
2. 相关包编译通过

## 阶段六：批量迁移所有字段类型构造

### 目标

去掉所有 `pq.StringArray` 依赖。

### 操作

1. 所有 generated model field 改为原生切片
2. 所有 handwritten 默认值与构造逻辑改为原生切片
3. 清理 import

### 验证

1. `rg -n "pq\\.StringArray" ...` 归零
2. 编译通过

## 阶段七：移除 lib/pq 依赖

### 目标

从依赖树中真正删除 `github.com/lib/pq`

### 操作

1. 全仓确认无引用
2. 执行 `go mod tidy`

### 验证

以下搜索结果应为空：

```text
github.com/lib/pq
pq.StringArray
pq.Array(
```

## 阶段八：全量 regenerate 与回归验证

### 目标

证明这次重构不是“手工修生成结果”，而是生成链路已经稳定。

### 操作

1. 重新执行 model 生成命令
2. 做最小但覆盖关键链路的编译与测试验证

### 最低验证建议

1. 生成验证
   - `make goctl-model-core`
   - `make goctl-model-relation`
   - `make goctl-model-log`

2. 编译验证
   - `go build ./service/model/...`
   - `go build ./service/rpc/blog/...`

3. 重点测试
   - article 相关
   - talk 相关
   - socialrpc / articlerpc 相关

### 成功判定

1. regenerate 后无 `lib/pq`
2. 编译通过
3. 关键测试通过

## 风险点

### 1. nil slice 与 empty slice 语义差异

当前部分逻辑显式设置 `pq.StringArray{}`，迁移后要明确保留：

- `nil`
或
- `[]string{}`

避免写库语义变化。

### 2. array 操作符参数推断

例如：

- `id = any($1)`
- `tags @> $1`

在直接传切片后，可能需要局部显式 cast。出现问题时只做局部 SQL 修正，不扩大改造范围。

### 3. goctl 模板能力不足

如果模板上下文里拿不到足够的字段原始类型信息，可能需要进一步调整 `goctl` 的类型映射层，而不仅是 `.tpl` 文件。

### 4. generated 与 handwritten 边界混淆

不能手改生成产物作为最终方案。所有会被 regenerate 覆盖的改动，必须回收到模板链路。

## 建议的里程碑

### 里程碑 1

完成模板验证：

- 自定义模板目录可用
- `t_article` / `t_talk` 生成结果不再出现 `pq.StringArray`

### 里程碑 2

完成核心表迁移：

- article / talk 相关链路编译通过
- handwritten `pq.StringArray` / `pq.Array` 用法在核心表范围内清零

### 里程碑 3

完成全仓迁移：

- `lib/pq` 依赖从代码与模块依赖中完全消失
- regenerate 不回退

## 推荐执行顺序

1. 建模板目录
2. 改模板
3. 接 `makefile --home`
4. 小范围 regenerate
5. 迁移 `t_article` / `t_talk`
6. 扩散到全部 handwritten 使用点
7. `go mod tidy`
8. 全量 regenerate + 回归验证

## 后续建议

如果开始实施，建议下一步再补一份更细的执行清单，按以下维度展开：

1. generated 文件列表
2. handwritten 文件列表
3. 模板文件列表
4. 每阶段验证命令
5. 每阶段回滚点

这样可以直接作为执行 checklist 使用。
