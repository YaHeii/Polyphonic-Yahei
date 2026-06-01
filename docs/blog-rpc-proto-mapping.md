# blog RPC proto 映射关系

## 目标

明确 `service/rpc/blog/proto` 目录中：

- 每个 proto 文件的职责
- `proto package` 与预期 `Go package` 的对应关系
- `blog.proto` 与各子 proto 的依赖关系
- 当前结构更接近哪一种“真相源”

## 当前目录结构

```text
service/rpc/blog/proto/
├── blog.proto
└── blog/
    ├── account.proto
    ├── article.proto
    ├── config.proto
    ├── news.proto
    ├── notice.proto
    ├── permission.proto
    ├── resource.proto
    ├── social.proto
    ├── syslog.proto
    └── website.proto
```

## 映射表

| proto 文件 | proto package | 当前 go_package | 预期 Go 包职责 | 被谁引用 |
| --- | --- | --- | --- | --- |
| `blog.proto` | `blogrpc` | `./blogrpc` | 聚合总 service，生成 `BlogRpc` 服务定义 | RPC server/client 主入口 |
| `blog/account.proto` | `accountrpc` | `./accountrpc` | 账号域 message 包 | `blog.proto` |
| `blog/article.proto` | `articlerpc` | `./articlerpc` | 文章域 message 包 | `blog.proto` |
| `blog/config.proto` | `configrpc` | `./configrpc` | 配置域 message 包 | `blog.proto` |
| `blog/news.proto` | `newsrpc` | `./newsrpc` | 留言评论等消息域 message 包 | `blog.proto` |
| `blog/notice.proto` | `noticerpc` | `./noticerpc` | 通知域 message 包 | `blog.proto` |
| `blog/permission.proto` | `permissionrpc` | `./permissionrpc` | 权限域 message 包 | `blog.proto` |
| `blog/resource.proto` | `resourcerpc` | `./resourcerpc` | 资源域 message 包 | `blog.proto` |
| `blog/social.proto` | `socialrpc` | `./socialrpc` | 社交域 message 包 | `blog.proto` |
| `blog/syslog.proto` | `syslogrpc` | `./syslogrpc` | 日志域 message 包 | `blog.proto` |
| `blog/website.proto` | `websiterpc` | `./websiterpc` | 站点统计域 message 包 | `blog.proto` |

## `blog.proto` 的角色

`blog.proto` 不是“公共 message 文件”，而是“总 RPC service 聚合文件”。

它做了两件事：

1. `import "blog/*.proto"`
2. 在 `service BlogRpc` 中直接引用各子 proto 的类型，例如：
   - `accountrpc.LoginReq`
   - `articlerpc.AnalysisArticleReq`
   - `websiterpc.AnalysisVisitReq`

因此它的真实职责是：

- 统一暴露一个 `BlogRpc` 服务
- 把 message 定义按业务域拆在多个子 proto 中

## 这个结构对应的真相源

当前 proto 设计最接近下面这一种真相源：

- proto 层：`blog.proto` 聚合多个子 proto
- Go 层：每个子 proto 生成独立 Go 包
- service 层：`blog.proto` 生成总 server/client 入口

这不是“单一 pb 包真相源”，而是“多 proto、多 package、单 service 聚合真相源”。

## 为什么不能全部改成 `option go_package = "./pb";`

如果全部改成 `./pb`，含义是：

- 所有 proto 都生成到同一个 Go 包
- 所有 `*.pb.go` 都会尝试落到同一个包命名空间

这和当前 proto 内容冲突，原因是多个子 proto 内部存在大量重名 message，例如：

- `PageReq`
- `PageResp`

这些名字在各自独立 proto package 中可以共存，但在同一个 Go 包里不能共存。

所以当前结构天然要求：

- 子 proto 保持独立包边界
- 不能简单扁平化到单一 `pb` 包

## 预期生成关系

如果按当前 proto 结构继续走，合理的生成关系应当是：

| proto package | 预期 Go 包 |
| --- | --- |
| `blogrpc` | `.../internal/pb/blogrpc` |
| `accountrpc` | `.../internal/pb/accountrpc` |
| `articlerpc` | `.../internal/pb/articlerpc` |
| `configrpc` | `.../internal/pb/configrpc` |
| `newsrpc` | `.../internal/pb/newsrpc` |
| `noticerpc` | `.../internal/pb/noticerpc` |
| `permissionrpc` | `.../internal/pb/permissionrpc` |
| `resourcerpc` | `.../internal/pb/resourcerpc` |
| `socialrpc` | `.../internal/pb/socialrpc` |
| `syslogrpc` | `.../internal/pb/syslogrpc` |
| `websiterpc` | `.../internal/pb/websiterpc` |

对应含义是：

- 每个子 proto 的 message 留在各自独立 Go 包
- `blogrpc` 只负责总 service 定义
- server/client 生成代码再 import 这些独立包

## 当前冲突点

当前问题不是 proto 的业务拆分错了，而是“生成规则”和“目录真相”还没有完全对齐，主要体现在：

1. `blog.proto` 的 import 根路径依赖 `service/rpc/blog/proto`
2. `goctl rpc protoc` 既要直接打开主 proto，又要生成 zrpc 代码
3. 当前 `go_package = "./xxxrpc"` 是相对导入语义，容易把生成代码推向相对 import

## 当前结论

先按下面这个结论继续是最稳的：

- `blog.proto` 保持总 service 聚合文件角色
- 各子 proto 保持独立 `proto package`
- Go 侧也保持独立子包，不收敛成单一 `pb`

下一步真正要确认的，不是“要不要一个 `pb` 包”，而是：

- `goctl rpc protoc` 是否适合直接消费这套“聚合 service + 多子 package”的结构
- 如果适合，正确的 `go_package` 和输出目录应该如何对齐
- 如果不适合，是否需要拆成“先 protoc 生成 pb，再单独生成 zrpc”的两段式流程
