# AGENTS.md

本文件是仓库根目录的开发约定，作用域覆盖整个 `Polyphonic-Yahei`。

如果某个子目录需要更细的规则，在更靠近目标目录的位置新增 `AGENTS.md` 或 `AGENTS.override.md`，并只写该子树特有的约束。

## 1. 沟通与文档

- 默认使用中文交流、中文 Markdown 文档。
- 代码注释默认使用英文，且只在关键逻辑、协议边界、生成链路等不自解释的位置添加。
- 纯展示、计划、说明文档默认放在 `docs/`。
- 不要把设计态讨论写成“已实现事实”；实现、计划、待验证三者要明确区分。

## 2. 修改原则

- 编码前先明确成功标准、影响范围和验证方式；不要靠猜。
- 只修改完成当前任务所必需的文件，不顺手清理无关代码。
- 优先复用现有模式、现有目录结构、现有 helper，不要平白加抽象层。
- 不要为一次性需求做“未来可扩展性”设计。
- 如果发现无关死代码，可以指出，但除非用户明确要求，不要顺手删除。
- 搜索优先使用 `rg` / `rg --files`。

## 3. 生成代码与真相源

### 3.1 不要直接改生成产物

以下内容默认视为生成产物，除非任务明确要求验证或修生成器，否则不要直接手改：

- `service/rpc/blog/client/*_rpc.go`
- `service/rpc/blog/internal/pb/*`
- `service/model/*_gen.go`
- `service/api/admin/internal/docs/swagger.json`
- `service/api/admin/internal/docs/docs.go` 中由生成器覆盖的部分

如果生成代码缺少补充能力，应在手写文件中补，不要改生成文件本体。例如：

- RPC client 缺少别名时，补到同目录 `types.go`
- model 类型或输出不符合要求时，优先改 `.goctl` 模板或生成链路

### 3.2 先改真相源，再重新生成

- Admin API 的真相源是 `service/api/admin/proto/*.api`
- Blog RPC 的真相源是 `service/rpc/blog/proto/**/*.proto`
- Model 的真相源是 PostgreSQL schema、`.goctl` 模板和 `goctl model` 命令
- Swagger 文档从 `.api` 生成，不要把 `swagger.json` 当手写源

常用生成命令：

```bash
make goctl-api-admin
make goctl-api-admin-swagger
make goctl-rpc-blog
make goctl-model-all
```

## 4. 测试规范

### 4.1 默认策略

- 默认采用 TDD 或至少“先补失败测试再修复”的方式推进。
- 测试库统一使用 `github.com/stretchr/testify`。
- 新增或修改测试时，优先使用：
  - `require.*` 处理前置条件、错误、空值、长度等必须立即中断的检查
  - `assert.*` 处理可以成组展示的值比较
- 新测试里不要把 `t.Fatalf` / `t.Errorf` 当作常规断言手段；它们只适合极少量测试辅助构造失败分支。

推荐写法：

```go
require.NoError(t, err)
require.Len(t, items, 1)
assert.Equal(t, "expected", items[0].Name)
```

### 4.2 测试范围

- 纯文档改动、纯注释改动、纯配置搬运、纯生成文件刷新，可以不补业务测试，但要说明为什么不需要。
- 狭窄逻辑改动：优先补包级单元测试。
- HTTP/logic 层改动：优先直接测 logic 或 handler，不要默认拉整套容器。
- 跨模块、协议、生成链路改动：补最小必要的集成验证。

### 4.3 验证要求

至少运行与改动范围匹配的验证命令，常用命令如下：

```bash
GOCACHE=/tmp/go-build /usr/local/go/bin/go test ./... -count=1
GOCACHE=/tmp/go-build /usr/local/go/bin/go build ./service/api/admin
GOCACHE=/tmp/go-build /usr/local/go/bin/go build ./service/rpc/blog
```

做局部改动时，可以只跑受影响包，但最终说明里要写清楚跑了哪些命令。

## 5. 本地联调与运行

当前本地最小后端链路以 `docker compose` 为准，目标是保证以下服务可启动并联通：

- `admin-api`
- `blog-rpc`
- `postgres`
- `redis`
- `rabbitmq`

常用命令：

```bash
make compose-config
make compose-build
make compose-up
make compose-down
make compose-logs
```

冒烟检查至少覆盖：

- `GET /ping`
- 登录拿 token
- 一个带鉴权的受保护接口

系统级接口联调与回归分组，参考 `docs/admin-backend-system-test-plan.md`。

## 6. 后端约束

- 保留现有鉴权链路，不要为了“方便本地调试”绕开 token 逻辑。
- 上传当前以 Linux 本地存储为主；如果改上传逻辑，必须同时检查存储路径和静态文件路由是否一致。
- 配置中的敏感信息优先放 `.env`，Compose 负责编排注入；不要把明文凭证重新写回 YAML。

## 7. 提交与交付

- 一个 commit 只做一类逻辑变更，提交信息要能说明范围。
- 如果任务跨多个相互独立的 API 面或模块，优先按可验证的小步提交。
- 最终交付要说明：
  - 改了哪些文件
  - 为什么这样改
  - 运行了哪些验证
  - 哪些东西没有验证
