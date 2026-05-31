为了将原有的 MySQL 初始化脚本全面改造为高性能、高度可扩展的 **PostgreSQL 单机多模态架构**，以下为你制定一份详细的类型更改与表重构计划。

---

## 一、 数据类型精密映射规范（MySQL $\rightarrow$ PostgreSQL）

在单服务器部署下，精简类型不仅能大幅节省单机内存与磁盘 I/O，还能避免无谓的类型转换开销。请严格执行以下映射标准：

| MySQL 类型 | PostgreSQL 映射类型 | 改进与重构意图说明 |
| --- | --- | --- |
| `int unsigned AUTO_INCREMENT` | `bigserial` 或 `serial` | PgSQL 不支持 `unsigned` 关键字。主键或关联 ID 统一采用序列类型 `serial`（4字节，足够博客系统使用）或 `bigserial`（8字节）。 |
| `tinyint(1)` | `boolean` | 状态开关、逻辑删除等布尔概念，必须使用标准的 `boolean` 类型（`TRUE` / `FALSE`），消除整型歧义。 |
| `tinyint` / `smallint` | `smallint` | 状态值（如文章状态 1、2、3）在 PgSQL 中统一使用 `smallint`（2字节），比 4字节的 `integer` 更节约存储。 |
| `datetime` | `timestamptz` | 即 `timestamp with time zone`。带时区的时间戳，原生规避由于单机 Docker 容器时区漂移带来的历史断层风险。 |
| `longtext` / `text` | `text` | PgSQL 中的 `text` 类型在底层通过 TOAST 机制管理，性能极高（等同或优于 `varchar`），长度无硬性限制，适合存放文章内容或大文本。 |
| `varchar(N)` | `varchar(N)` | 保持一致，但 PgSQL 的 `varchar` 改变长度通常不需要锁表，扩展性极佳。 |

---

## 二、 核心表重构核心四法（压榨多模态红利）

针对博客系统的复杂业务，我们利用 PostgreSQL 的**数组、JSONB、GIN 索引**特性，对原有表结构进行深度重构：

### 1. 数组化（Arrays）取代传统多对多中间表

* **重构对象**：`t_article` 与 `t_article_tag`。
* **做法**：彻底干掉 `t_article_tag` 物理表。在 `t_article` 表中直接增加 `tags text[]` 字段，配合 **GIN 索引**。
* **收益**：单机部署下少维护一张物理表，查询“包含某个标签的文章”时，由 3 表 `JOIN` 退化为单表 GIN 索引扫描，吞吐量提升数倍。

### 2. JSONB 泛化（Document）取代碎片化字段与多变配置

* **重构对象**：`t_website_config`、`t_menu`、`t_article`。
* **做法**：
* 将 `t_website_config` 的 `config varchar(4096)` 改为 `jsonb`。
* 将 `t_menu` 的 `extra` 字段改为 `jsonb`。
* 在 `t_article` 中增加 `metadata jsonb` 字段，用来承载未来的置顶、打赏、SEO 优化等零碎字段。


* **收益**：AI Agent 或团队在扩展菜单属性或新增网站配置时，**零 DDL 变更**，直接操作 JSONB，且支持对 JSONB 建立 GIN 路径索引。

### 3. 原生全文检索（Full-Text Search）取代重型搜索引擎

* **重构对象**：`t_article` 的检索。
* **做法**：对 `article_title` 和 `article_content` 建立基于 `to_tsvector` 的虚拟组合列或直接创建 GIN 表达式索引。
* **收益**：单机不需要再额外背负 Elasticsearch 这样动辄消耗 1~2G 内存的巨兽，完美守住单服务器内存底线。

### 4. 统一触发器自动化 `updated_at`

* **做法**：PgSQL 不支持 MySQL 的 `ON UPDATE CURRENT_TIMESTAMP`。必须在初始化时声明一个全局 plpgsql 触发器函数，一键挂载到所有带有 `updated_at` 的表上，确保时序数据的准确性。

---

## 三、 全套 PostgreSQL 重构建表脚本（生产可用）

以下是经过完整重构与优化后的 `.sql` 脚本，可以直接用于 PostgreSQL 初始化：

```sql
-- =============================================================================
-- 0. 基础基础设施准备
-- =============================================================================
SET client_encoding = 'UTF8';

-- 创建自动更新 updated_at 的通用触发器函数
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- =============================================================================
-- 1. 用户与权限模块
-- =============================================================================

DROP TABLE IF EXISTS "t_user" CASCADE;
CREATE TABLE "t_user" (
    "id" serial NOT NULL,
    "user_id" varchar(64) NOT NULL DEFAULT '',
    "username" varchar(64) NOT NULL DEFAULT '',
    "password" varchar(128) NOT NULL DEFAULT '',
    "nickname" varchar(64) NOT NULL DEFAULT '',
    "avatar" varchar(255) NOT NULL DEFAULT '',
    "email" varchar(64) NOT NULL DEFAULT '',
    "phone" varchar(64) NOT NULL DEFAULT '',
    "info" varchar(1024) NOT NULL DEFAULT '',
    "status" smallint NOT NULL DEFAULT 0, -- -1删除 0正常 1禁用
    "register_type" varchar(64) NOT NULL DEFAULT '',
    "ip_address" varchar(255) NOT NULL DEFAULT '',
    "ip_source" varchar(255) NOT NULL DEFAULT '',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "uk_user_uid" UNIQUE ("user_id"),
    CONSTRAINT "uk_user_username" UNIQUE ("username")
);

CREATE TRIGGER update_user_modtime BEFORE UPDATE ON "t_user" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

-- =============================================================================
-- 2. 内容管理模块（文章、分类、说说）
-- =============================================================================

DROP TABLE IF EXISTS "t_category" CASCADE;
CREATE TABLE "t_category" (
    "id" serial NOT NULL,
    "category_name" varchar(32) NOT NULL DEFAULT '',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "uk_category_name" UNIQUE ("category_name")
);

CREATE TRIGGER update_category_modtime BEFORE UPDATE ON "t_category" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

DROP TABLE IF EXISTS "t_article" CASCADE;
CREATE TABLE "t_article" (
    "id" serial NOT NULL,
    "user_id" varchar(64) NOT NULL DEFAULT '',
    "category_id" integer NOT NULL DEFAULT 0,
    "article_cover" varchar(1024) NOT NULL DEFAULT '',
    "article_title" varchar(64) NOT NULL DEFAULT '',
    "article_content" text NOT NULL, -- 使用高性能的 text 类型
    "article_type" smallint NOT NULL DEFAULT 0, -- 1原创 2转载 3翻译
    "original_url" varchar(255) NOT NULL DEFAULT '',
    "tags" text[] NOT NULL DEFAULT '{}', -- 【核心改进】原生文本数组存放标签
    "metadata" jsonb NOT NULL DEFAULT '{}', -- 【核心改进】JSONB 存放文章动态扩展元数据
    "is_top" boolean NOT NULL DEFAULT FALSE, -- 使用标准布尔型
    "is_delete" boolean NOT NULL DEFAULT FALSE,
    "status" smallint NOT NULL DEFAULT 1, -- 1公开 2私密 3草稿 4评论可见
    "like_count" integer NOT NULL DEFAULT 0,
    "view_count" integer NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_article_modtime BEFORE UPDATE ON "t_article" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

-- 为文章标签建立强力的 GIN 倒排索引
CREATE INDEX "idx_article_tags" ON "t_article" USING GIN ("tags");

-- 为文章建立原生的中文/英文全文检索 GIN 索引（单机检索防御壁垒）
CREATE INDEX "idx_article_fulltext" ON "t_article" USING GIN (to_tsvector('simple', "article_title" || ' ' || "article_content"));

DROP TABLE IF EXISTS "t_talk" CASCADE;
CREATE TABLE "t_talk" (
    "id" serial NOT NULL,
    "user_id" varchar(64) NOT NULL DEFAULT '',
    "content" varchar(2048) NOT NULL DEFAULT '',
    "images" text[] NOT NULL DEFAULT '{}', -- 【改进】说说图片列表直接上数组，免去逗号分割字符串的痛苦解析
    "is_top" boolean NOT NULL DEFAULT FALSE,
    "status" smallint NOT NULL DEFAULT 1, -- 1公开 2私密
    "like_count" integer NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_talk_modtime BEFORE UPDATE ON "t_talk" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

-- =============================================================================
-- 3. 互动与系统控制模块（评论、网站配置、菜单）
-- =============================================================================

DROP TABLE IF EXISTS "t_comment" CASCADE;
CREATE TABLE "t_comment" (
    "id" serial NOT NULL,
    "user_id" varchar(64) NOT NULL DEFAULT '',
    "terminal_id" varchar(64) NOT NULL DEFAULT '',
    "topic_id" integer NOT NULL DEFAULT 0,
    "parent_id" integer NOT NULL DEFAULT 0,
    "reply_id" integer NOT NULL DEFAULT 0,
    "reply_user_id" varchar(255) NOT NULL DEFAULT '',
    "comment_content" text NOT NULL,
    "type" smallint NOT NULL DEFAULT 0, -- 1.文章 2.友链 3.说说
    "status" smallint NOT NULL DEFAULT 0, -- 0正常 1已编辑 2已删除
    "like_count" integer NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_comment_modtime BEFORE UPDATE ON "t_comment" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
CREATE INDEX "idx_comment_parent" ON "t_comment" ("parent_id");

DROP TABLE IF EXISTS "t_menu" CASCADE;
CREATE TABLE "t_menu" (
    "id" serial NOT NULL,
    "parent_id" integer NOT NULL DEFAULT 0,
    "path" varchar(64) NOT NULL DEFAULT '',
    "name" varchar(64) NOT NULL DEFAULT '',
    "component" varchar(256) NOT NULL DEFAULT '',
    "redirect" varchar(256) NOT NULL DEFAULT '',
    "type" varchar(64) NOT NULL DEFAULT '0',
    "title" varchar(64) NOT NULL DEFAULT '',
    "icon" varchar(64) NOT NULL DEFAULT '',
    "rank" integer NOT NULL DEFAULT 0,
    "perm" varchar(64) NOT NULL DEFAULT '',
    "params" varchar(256) NOT NULL DEFAULT '',
    "keep_alive" boolean NOT NULL DEFAULT FALSE,
    "always_show" boolean NOT NULL DEFAULT FALSE,
    "visible" boolean NOT NULL DEFAULT FALSE,
    "status" boolean NOT NULL DEFAULT FALSE,
    "extra" jsonb NOT NULL DEFAULT '{}', -- 【核心改进】前端多元路由数据直接上 JSONB
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "uk_menu_path_perm" UNIQUE ("path", "perm")
);

CREATE TRIGGER update_menu_modtime BEFORE UPDATE ON "t_menu" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

DROP TABLE IF EXISTS "t_website_config" CASCADE;
CREATE TABLE "t_website_config" (
    "id" serial NOT NULL,
    "key" varchar(32) NOT NULL DEFAULT '',
    "config" jsonb NOT NULL DEFAULT '{}', -- 【核心改进】网站动态配置彻头上 JSONB
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "uk_website_config_key" UNIQUE ("key")
);

CREATE TRIGGER update_website_config_modtime BEFORE UPDATE ON "t_website_config" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
-- 为配置的 JSONB 的任意深度查询注入 GIN 性能加速
CREATE INDEX "idx_website_config_jsonb" ON "t_website_config" USING GIN ("config");

```

---

## 四、 后续行动指南（防止工具链摩擦）

将此 SQL 导入 PostgreSQL 后，请让 AI Agent 在执行 `goctl model pgsql` 时注意以下细节：

1. **删除冗余实体**：在 `service/model/` 目录下，不需要再为 `t_article_tag` 生成任何代码，因为它已经在代码世界和数据库世界中蒸发了。
2. **标签增删逻辑简化**：在逻辑层（Logic）为文章添加或删除标签时，直接对 Go 中的 `[]string` 切片进行操作即可（GORM 或 go-zero 的底层绑定会直接将其翻译为 PgSQL 识别的 `{"Go","后端"}` 数组语法格式输入）。