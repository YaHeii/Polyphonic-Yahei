# AGENTS.md

本文件作用域覆盖 `service/rpc/blog/proto/blog` 及其子目录，只补充 blog RPC proto 这一层的局部约束；未提到的事项继续遵循仓库根目录 [AGENTS.md](/root/Polyphonic-Yahei/AGENTS.md)。

## 1. 这一层的职责

- 这里维护的是 blog RPC 的领域 proto 真相源，不是通用 gRPC 教程或后端分层规范。
- 当前 blog RPC 的输入集就是 `service/rpc/blog/proto/blog/*.proto`，每个领域文件各自声明 `service XxxRpc`。
- 改 blog RPC 契约时，优先改这里的 `.proto` 文件；不要先改 `internal/pb`、`client/*_rpc.go`、`internal/server/*` 或 `blog.go`。

## 2. 真相源与生成边界

- `service/rpc/blog/proto/blog/*.proto` 是 blog RPC 的契约真相源。
- `service/rpc/blog/internal/pb/*`、`service/rpc/blog/client/*/*_rpc.go`、`service/rpc/blog/internal/server/*` 中由 goctl 生成或 scaffold 的部分，都不是手写真相源。
- 当前生成入口是 `make goctl-rpc-blog`，它会循环扫描 `service/rpc/blog/proto/blog/*.proto` 并逐个执行 `goctl rpc protoc`。
- 不要直接 patch 生成 client wrapper 去补类型暴露；下次 regenerate 会覆盖。

常用生成命令：

```bash
make goctl-rpc-blog
```

## 3. proto 书写规则

- 每个 proto 文件只表达一个领域 service 及其相关 message，继续沿用当前按领域拆分的结构，例如 `account.proto`、`article.proto`、`permission.proto`。
- 先匹配当前仓库既有模式：每个文件都显式声明自己的 `package`、`option go_package`、通用消息段、基础消息段、领域 RPC 段。
- 当前仓库允许不同 proto 文件各自拥有同名消息，例如 `PageReq`、`PageResp`；不要为了“去重”强行把它们压成跨领域共享包，避免生成后的 Go 类型冲突。
- 新增 RPC 时，`rpc` 名称、请求消息、响应消息、server 实现目录应保持稳定映射，便于 goctl 生成与人工排查。
- 修改字段时优先做兼容性判断；已经被 API、client alias 或其他领域 logic 依赖的字段，不要随意改名、换 tag、改语义。

## 4. client alias 规则

- goctl 生成的 `client/*_rpc.go` 只暴露 RPC 方法以及每个方法直接 input/output 对应的 alias，不会自动暴露所有嵌套 pb 类型。
- 像 `PageReq`、`PageResp`、`UserInfo`、`ArticleDetails` 这类“pb 里存在但 client wrapper 没直接导出”的类型，正确补法是 hand-written `service/rpc/blog/client/<domain>/types.go`。
- `service/rpc/blog/client/*/types.go` 是 regeneration-safe 扩展层；需要额外暴露嵌套消息时，只在这里补 alias，不改生成的 `*_rpc.go`。
- 当你在 proto 中新增会被外部直接构造或返回后需要直接引用的嵌套消息时，要同步检查对应 `client/<domain>/types.go` 是否需要补 alias。

## 5. `blog.go` 聚合入口规则

- [blog.go](/root/Polyphonic-Yahei/service/rpc/blog/blog.go) 是 blog RPC 的手写聚合启动入口，不是生成产物。
- 每个 proto 文件中的 `service XxxRpc`，最终都需要在 `blog.go` 中完成两类同步：
  - 注册对应的 pb server：`internal/pb/<domain>rpc`
  - 引入对应的 server 实现：`internal/server/<domain>rpc`
- 新增、删除、重命名某个 proto service 后，必须同步检查 `blog.go` 的 import 和 `Register*Server(...)` 调用；不要假设 goctl 会自动替你更新聚合入口。
- `blog.go` 只应导入当前 checkout 中真实存在的 `internal/server/<domain>` 包；不要引用不存在的 umbrella package 或空目录。

## 6. account proto 的 verification 边界

- `account.proto` 既是账号领域 RPC 契约，也是最容易把 API verification 语义重新塞回 RPC 的地方；改这个文件时必须额外检查边界。
- 当前约束是：验证码、短信码、邮件码、OAuth code 换取外部身份，属于 API 边界；RPC 只接 post-verification business fields。
- `AccountRpc` 请求中不应重新引入这类字段：
  - `verify_code`
  - `captcha_key`
  - `captcha_code`
  - 其他同类前置验证字段
- 密码、邮箱、手机号、`platform`、`open_id` 这类业务或凭证字段可以继续存在于 RPC；但“如何验证它们是否合法”不应回流到 RPC。
- 若需求会改变这条边界，先更新约束文档或 spec，再改 proto，不要直接把验证字段加回 `account.proto`。

可参考现有边界说明：

- [docs/spec/2026-06-12-identity-auth-boundary.md](/root/Polyphonic-Yahei/docs/spec/2026-06-12-identity-auth-boundary.md)

## 7. 改 proto 时的联动检查单

- 改了 `.proto`：先确认是否只影响当前领域，还是会影响 admin API、client alias、`blog.go` 聚合入口。
- 新增/删除/重命名 `message`：检查 `client/<domain>/types.go` 是否需要补或删 alias。
- 新增/删除/重命名 `service` 或 `rpc`：检查 `blog.go` 是否需要同步注册或调整 import。
- 改 account 相关请求字段：检查是否把 verification 语义错误地下沉回 RPC。
- 改公共分页或值对象：检查下游 logic、API convert、client 调用方是否依赖旧字段名或旧语义。

## 8. 验证方式

- 文档改动可不补业务测试，但若改了 proto，至少跑生成链路与受影响构建/测试。
- 常用命令：

```bash
make goctl-rpc-blog
GOCACHE=/tmp/go-build /usr/local/go/bin/go test ./service/rpc/blog/... -count=1
GOCACHE=/tmp/go-build /usr/local/go/bin/go build ./service/rpc/blog
```

- 如果 `go mod tidy` 在当前环境里出现只读缓存噪音，按退出码判断是否真失败；最终以 `go test` / `go build` 的结果为准。

## 9. 什么时候需要先停下来收敛

- 如果需求要把多个 proto 合并成统一包、抽公共 proto、改 `go_package` 策略、或重做 `blog.go` 聚合结构，先在 `docs/` 写短 spec，再开始改代码。
- 如果只是给现有领域补一个 RPC、补一个 message、或补一个 client alias，按现有模式直接改并验证即可，不要扩成整套生成链路重构。
