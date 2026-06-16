# Polyphonic-Yahei

> 复调原指包含两个或两个以上独立且同等重要的旋律同时进行、交织在一起的音乐形式.后来巴赫金用这个词来形容陀氏的小说, 
> 试图描述他书中那个价值观各自独立,声音嘈杂,然而精彩纷呈的那个世界
>
> 复调也存在于一个人的身上, 无论是价值观还是个性特点, 都会交织的在我们身上体现.

# 关于这个项目

这是一个个人博客项目, 目前正在施工🚧. 


## 技术栈

- `Go 1.25`
- `go-zero`
- `gRPC / protobuf`
- `PostgreSQL`
- `Redis`
- `RabbitMQ`
- `Swagger`
- `Docker Compose`
- `goctl`

## 目录速览

```text
.
├── common/                   # 通用常量与基础枚举
├── docs/                     # 设计说明、约束、计划文档
├── pkg/                      # 基础设施与项目级工具能力
├── service/
│   ├── api/admin/            # 管理端 API
│   ├── db/                   # 数据库 migration 与 bootstrap seed 真相源
│   ├── model/                # 数据模型与生成产物
│   └── rpc/blog/             # Blog RPC 服务
├── web/                      # Web 相关目录，前端能力仍在逐步补齐
├── docker-compose.yaml       # 本地联调主入口
└── makefile                  # 生成与联调命令入口
```

## 快速启动

如果你只是想先把这套后端跑起来，可以从这里开始：

```bash
make env-init
docker compose --env-file .env -f docker-compose.yaml up -d postgres
make migrate-up
make seed-bootstrap
make compose-up-build
```

默认会拉起这些服务：

- `postgres`
- `redis`
- `rabbitmq`
- `blog-rpc`
- `admin-api`

启动后可以先做一个最朴素的冒烟检查：

```bash
curl http://127.0.0.1:9091/admin-api/v1/ping
```

常用联调命令：

```bash
make migrate-version
make migrate-down
make compose-logs
make compose-ps
make compose-down
```

数据库真相源已经收口到 `service/db`：

- schema 变更放 `service/db/migrations`
- 必需初始化数据放 `service/db/seeds/bootstrap`

🫡 参考
https://github.com/ve-weiyi/ve-blog-golang
