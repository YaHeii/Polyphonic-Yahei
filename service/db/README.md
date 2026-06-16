# service/db

`service/db` 是当前仓库的数据库真相源目录。

## 目录职责

- `migrations/`
  - 只放 schema / index / constraint / trigger / function 变更
  - 使用 `golang-migrate` 的 `*.up.sql` / `*.down.sql` 约定
- `seeds/bootstrap/`
  - 只放系统启动必需的 bootstrap 数据
  - 这些 SQL 必须幂等，可重复执行

## 命名规则

- baseline schema：`000001_blog_init.up.sql`
- baseline rollback：`000001_blog_init.down.sql`
- 后续 schema 变更：`000002_xxx.up.sql` / `000002_xxx.down.sql`
- bootstrap seed：按职责递增编号，例如：
  - `001_auth_bootstrap.sql`
  - `002_permission_bootstrap.sql`
  - `003_site_bootstrap.sql`

## 执行顺序

本地、测试、线上统一走显式命令：

```bash
make migrate-up
make seed-bootstrap
```

回滚与重放：

```bash
make migrate-down
make migrate-force VERSION=1
make seed-bootstrap
```

## 工具依赖

`make` 目标默认依赖宿主机安装：

- `migrate`
- `psql`

默认通过宿主机地址连接数据库，因此 `make` 侧会优先使用 `127.0.0.1` 作为 CLI 连接地址；如需覆盖，可传：

```bash
make migrate-up DB_CLI_HOST=postgres
```

## Bootstrap 边界

允许放入 `seeds/bootstrap/` 的数据：

- seed admin 用户
- 基础角色
- 基础菜单
- 基础 API 权限资源
- 角色与用户、角色与菜单、角色与 API 的绑定
- 网站基础配置

不应放入 `seeds/bootstrap/` 的数据：

- 示例文章
- 示例评论/留言
- 相册、照片、友链等演示内容
- 只服务于本地演示的样例数据
