# 管理后台前端搭建方案

状态：设计已收敛，可作为第一阶段实现基线；第 4 节到第 8 节为本次一次性补全内容，尚未逐段审阅

## 1. 文档目的

本文用于收敛 `Polyphonic-Yahei` 在当前仓库内新增管理后台前端的第一阶段方案。

本文关注以下问题：

- 是否需要调整现有仓库目录
- 前端应用应该放在哪个目录
- `React + Ant Design + ProComponents + Swagger` 的边界如何划分
- 第一阶段先做哪些页面、哪些能力
- Swagger 生成链路、请求适配层、登录态流转如何落位

本文描述的是目标方案与实施基线，不代表仓库当前已经实现这些内容。

## 2. 方案结论

### 2.1 技术路线

- 前端技术栈采用 `React + TypeScript + Vite`
- UI 采用 `antd + @ant-design/pro-components`
- 接口类型与 client 采用 Swagger 生成方案
- 推荐生成器采用 `Orval`
- 请求实现采用手写 `request.ts` 统一适配项目级 header、响应解包和错误处理

### 2.2 工程形态

- 在仓库根目录新增独立前端应用：`web/admin`
- 不引入 `Ant Design Pro / Umi / Max` 这类重后台脚手架
- 只使用 `Vite` 作为工程初始化脚手架
- 第一阶段不接入现有 `docker compose` 联调链路
- 第一阶段不引入根级 Node workspace，不新增根级 `package.json`

### 2.3 后端与真相源边界

- 后端 API 真相源保持为 `service/api/admin/proto/**/*.api`
- Swagger 产物保持由后端生成，位置保持为 [swagger.json](/root/Polyphonic-Yahei/service/api/admin/internal/docs/swagger.json:1)
- 前端不手写业务接口类型，不长期维护第二份接口定义
- 前端只消费当前仓库中由后端生成的 Swagger 产物

### 2.4 仓库目录边界

- 不重构 `service/api`、`service/rpc`、`service/model`
- 不因为前端接入而重构 `pkg/`、`common/`
- 不因为前端接入而调整当前 SQL 文件位置
- 前端不直接依赖 Go 侧 `pkg/`、`common/` 的常量或类型

## 3. 非目标

第一阶段明确不做以下内容：

- 不做全量管理后台页面铺设
- 不做动态菜单拉取与动态路由
- 不做按钮级权限控制
- 不做 Docker 化前端运行与容器联调
- 不做 Swagger 到页面 schema 的自动 CRUD 页面生成
- 不做完整设计系统或品牌视觉重构
- 不为了前端生成器反向调整当前后端 API 设计

## 4. 应用结构与目录职责

### 4.1 仓库级目录

前端应用新增为：

```text
.
├── service/
│   ├── api/
│   ├── model/
│   └── rpc/
├── web/
│   └── admin/
└── docs/
```

该结构的约束如下：

- `service/` 继续只承载 Go 后端服务、模型与生成链路
- `web/admin` 作为独立前端应用，不嵌入 `service/api/admin`
- `docs/` 继续承载方案、规范、计划等文档

### 4.2 `web/admin` 顶层结构

建议前端应用采用如下结构：

```text
web/admin
├── package.json
├── tsconfig.json
├── vite.config.ts
├── index.html
├── orval.config.ts
├── public/
├── scripts/
└── src/
    ├── app/
    ├── layouts/
    ├── pages/
    ├── router/
    ├── services/
    ├── access/
    ├── components/
    ├── utils/
    └── main.tsx
```

各目录职责如下：

- `src/app`
  - 应用级装配
  - 放全局 Provider、初始化逻辑、全局样式入口
- `src/layouts`
  - 后台布局组件
  - 以 `ProLayout` 为主
- `src/pages`
  - 页面级组件
  - 按模块拆目录，例如 `auth/`、`account/`、`category/`
- `src/router`
  - 静态路由定义
  - 第一阶段不做后端菜单驱动
- `src/services`
  - 接口访问层
  - 包含生成代码与手写适配层
- `src/access`
  - 路由守卫、登录校验、权限占位逻辑
- `src/components`
  - 可复用轻量业务组件
- `src/utils`
  - 与业务弱相关的工具函数
- `scripts`
  - 仅放前端本地辅助脚本
  - 能用标准 `npm/pnpm script` 解决的问题不额外造脚本

### 4.3 `src/services` 内部边界

建议结构如下：

```text
src/services
├── generated/
├── request.ts
├── client.ts
├── session.ts
└── auth.ts
```

边界约束如下：

- `generated/`
  - 只放 Swagger 自动生成产物
  - 不手改
- `request.ts`
  - 创建统一请求实例
  - 负责 `baseURL`、headers、错误处理、响应解包
- `client.ts`
  - 作为 `Orval` 的 `mutator` 桥接层
  - 把生成 client 统一接到 `request.ts`
- `session.ts`
  - 管理 token、用户 ID、应用名等前端会话数据
- `auth.ts`
  - 提供少量高层认证相关调用与清理逻辑

页面层原则：

- 页面不自己拼鉴权头
- 页面不自己判断 `code === 0`
- 页面不直接依赖 Swagger 文件路径
- 页面优先调用 `generated client + 薄包装 service`

### 4.4 状态管理边界

第一阶段不引入重状态管理方案。

约束如下：

- 不引入 `redux`、`mobx`
- 登录态以 `session.ts + React Context` 或同等级简单方案维护
- 请求层拿 token 时优先走同步可读的 `session.ts`
- 页面级临时状态维持在组件内部

## 5. 页面范围与路由组织

### 5.1 第一阶段页面范围

第一阶段只覆盖能证明后台壳子成立的最小页面集：

- `登录页`
- `Dashboard` 占位页
- `账号列表示例页`
- `分类管理示例页`
- `404 / 无权限占位页`

第一阶段成功标准不是“后台完整”，而是：

- 前端工程能独立启动
- 能消费 Swagger 生成类型和 client
- 能完成登录态流转
- 能用列表页与表单页跑通核心链路

### 5.2 示例模块选择

第一阶段样板模块采用：

- `account`
  - 作为列表页样板
  - 对应接口以 [account.api](/root/Polyphonic-Yahei/service/api/admin/proto/admin/account.api:1) 为准
- `category`
  - 作为列表 + 新增/编辑表单样板
  - 对应接口以 [category.api](/root/Polyphonic-Yahei/service/api/admin/proto/admin/category.api:1) 为准

选择理由：

- `account` 列表查询参数明确，适合验证 `ProTable + 分页 + 查询`
- `category` 提供 `list/add/update/delete`，适合验证基础 CRUD 页面范式
- 这两个模块都不要求第一阶段先处理复杂上传、树结构、富文本或联动子表

### 5.3 路由组织

第一阶段采用静态路由，建议如下：

```text
/login
/
  /dashboard
  /account/list
  /category/list
  /403
  /*
```

约束如下：

- 第一阶段不接动态菜单接口
- 第一阶段不做按钮级权限
- `category` 的新增/编辑优先采用列表页内 `Modal/Drawer + ProForm`
- 第一阶段不强制拆出独立 `/category/edit/:id` 页面

## 6. Swagger 生成链路与请求适配

### 6.1 真相源与生成顺序

前端接口生成链路采用以下顺序：

1. 后端修改 `.api`
2. 后端执行 `make goctl-api-admin-swagger`
3. Swagger 更新 [swagger.json](/root/Polyphonic-Yahei/service/api/admin/internal/docs/swagger.json:1)
4. 前端执行 `pnpm gen:api`
5. 前端根据生成类型与编译错误修正页面代码

约束如下：

- 不手改 `swagger.json`
- 不把 Swagger 复制成第二份长期维护副本
- 不在前端维护脱离 Swagger 的自定义业务 DTO 真相源

### 6.2 为什么选 `Orval`

`Orval` 作为推荐生成器，主要因为：

- 可以直接消费当前 Swagger/OpenAPI 文档
- 可以生成 TypeScript 类型与请求函数
- 可以通过 `mutator` 接入项目自己的请求实例
- 适合把“生成代码”和“项目级请求策略”清晰分层

本方案不要求 `Orval` 自动生成页面，只要求其承担：

- 请求/响应类型生成
- 基础 client 生成
- 与项目 `request.ts` 的桥接

### 6.3 `request.ts` 的职责

基于当前 Swagger 产物，前端必须保留手写请求适配层。

原因：

- 当前 Swagger 响应普遍采用统一包装：`code/data/msg`
- 当前 Swagger 安全定义包含多个头部语义：`App-Name`、`Timestamp`、`Uid`、`Authorization`

因此 `request.ts` 负责：

- 设置 `baseURL`
- 注入 `Authorization`
- 注入 `App-Name`
- 注入 `Timestamp`
- 在可用时注入 `Uid`
- 解包统一响应
- 统一错误提示
- 统一处理未登录或登录失效

页面层目标：

- 页面默认拿到的是解包后的业务 `data`
- 页面不重复解析 `code/msg/data`

### 6.4 安全定义与运行时 header 的边界

虽然 [swagger.json](/root/Polyphonic-Yahei/service/api/admin/internal/docs/swagger.json:1) 中存在安全定义，但前端运行时 header 真相源仍以项目请求层约定为准。

约束如下：

- 不依赖生成器自动决定所有安全 header 注入行为
- header 注入规则统一写在 `request.ts` / `session.ts`
- 若后端后续新增或调整公共 header，优先改请求适配层，而不是逐页修补

### 6.5 生成代码目录约束

生成代码固定放在：

```text
web/admin/src/services/generated
```

并遵守以下规则：

- 生成目录视为生成产物，不手改
- 需要改生成行为时，优先改 `orval.config.ts`
- 需要补项目逻辑时，写在 `request.ts`、`client.ts` 或薄包装 service 中

## 7. 第一阶段实施步骤

以下步骤作为进入实现时的默认顺序：

1. 新建 `web/admin` 工程骨架
   - 使用 `Vite + React + TypeScript`
   - 验证：`pnpm dev` 可启动空应用
2. 安装基础依赖
   - 运行时：`antd`、`@ant-design/pro-components`、`@ant-design/icons`、`react-router-dom`、`axios`、`dayjs`
   - 开发依赖：`orval`
   - 验证：`pnpm build` 无依赖缺失
3. 接入 Swagger 生成链路
   - 配置 `orval.config.ts`
   - 读取仓库内 `service/api/admin/internal/docs/swagger.json`
   - 验证：`pnpm gen:api` 能输出 `src/services/generated`
4. 搭建请求适配层与会话层
   - 编写 `request.ts`、`client.ts`、`session.ts`
   - 验证：生成 client 能通过适配层发请求
5. 搭建基础布局和静态路由
   - 使用 `ProLayout`
   - 补 `/login`、`/dashboard`、`/403`、`/404`
   - 验证：未登录重定向、登录后进入主布局
6. 接入登录页
   - 对接现有登录接口
   - 验证：token 存储、刷新页面后仍能恢复登录态
7. 完成 `account` 列表页
   - 使用 `ProTable`
   - 验证：分页、查询、错误提示链路可用
8. 完成 `category` 列表与新增/编辑表单
   - 使用 `ProTable + Modal/Drawer + ProForm`
   - 验证：列表、创建、更新可用
9. 收尾最小化规范
   - 整理脚本、环境变量、README 或补充文档
   - 验证：新成员按文档可本地启动

## 8. 测试与验证策略

### 8.1 默认验证命令

后端 Swagger 生成验证：

```bash
make goctl-api-admin-swagger
```

前端生成与构建验证：

```bash
pnpm install
pnpm gen:api
pnpm build
```

本地开发验证：

```bash
pnpm dev
```

### 8.2 第一阶段必做验证

至少覆盖以下验证点：

- Swagger 能生成前端 client
- 构建通过
- 登录页可用
- 登录态失效时能跳回登录页
- `account` 列表可查询并展示分页结果
- `category` 列表可新增、编辑

### 8.3 第一阶段测试边界

第一阶段不强求完整页面测试体系，但建议补最小必要的前端单元测试：

- `request.ts` 的响应解包逻辑
- 未登录或错误码分支的通用处理
- `session.ts` 的 token 读写

第一阶段不要求：

- 端到端自动化测试
- 大量 UI 快照测试
- 为每个页面补齐完整组件测试

## 9. 风险、假设与待确认项

### 9.1 已知风险

- Swagger 当前表达的是接口层，不是页面 schema；不能期待自动生成完整 CRUD 页面
- 当前公共 header 语义虽然在 Swagger 中有定义，但运行时规则仍需以前端请求层实测为准
- 若部分接口返回结构在 Swagger 中描述不充分，前端生成类型可能仍需通过实际联调修正
- 登录、刷新 token、登出链路可能受当前后端实现细节影响，需要在第一阶段联调时校准

### 9.2 当前假设

- 登录接口可作为第一阶段前端接入入口
- `account` 和 `category` 接口已足够支撑列表/表单样板
- 当前后台 API 启动方式由后端现有流程负责，前端第一阶段不负责容器化
- 前端第一阶段可以接受静态菜单与静态路由

### 9.3 待后续确认但不阻塞第一阶段启动的项

- 是否需要接入刷新 token 自动续期
- `App-Name` 的固定取值与来源
- `Uid` 是否应始终由前端注入，或仅在部分接口注入
- 后续是否把更多模块扩展到统一页面范式
- 第二阶段是否接动态菜单与权限矩阵

## 10. 审阅状态

### 10.1 已在讨论中明确确认的内容

- 采用轻量路线，而不是全量 `Ant Design Pro` 脚手架
- 前端应用新增为 `web/admin`
- 不重构 `service/api`、`service/rpc`、`service/model`
- 不因为前端接入而调整 `pkg/`、`common/`、SQL 文件位置
- 采用 `React + Ant Design + ProComponents + Swagger`
- UI 第一阶段以现成后台组件范式为主，不先做重视觉定制

### 10.2 本次一次性补全、待后续逐段审阅的内容

- `web/admin` 具体目录职责细化
- `src/services` 分层约束
- `Orval` 作为推荐生成器的落位
- `account + category` 作为第一阶段样板模块
- 第一阶段实施顺序
- 最小测试与验证策略
- 风险、假设与待确认项

## 11. 实施入口建议

如果基于本文继续进入实现，默认从以下顺序开始：

1. 在 `web/admin` 初始化 `Vite + React + TypeScript`
2. 建立 `Orval + request.ts` 生成与适配链路
3. 先跑通登录、布局、静态路由
4. 再实现 `account` 列表和 `category` 列表/表单样板

以上顺序优先保证“架子先立住”，而不是一开始铺满页面或接入所有后台能力。
