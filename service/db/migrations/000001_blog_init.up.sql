SET client_encoding = 'UTF8';

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS trigger AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE t_album (
    id bigserial PRIMARY KEY,
    album_name varchar(64) NOT NULL DEFAULT '',
    album_desc varchar(128) NOT NULL DEFAULT '',
    album_cover varchar(255) NOT NULL DEFAULT '',
    is_delete boolean NOT NULL DEFAULT false,
    status smallint NOT NULL DEFAULT 1,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_album IS '相册';
COMMENT ON COLUMN t_album.id IS '主键';
COMMENT ON COLUMN t_album.album_name IS '相册名';
COMMENT ON COLUMN t_album.album_desc IS '相册描述';
COMMENT ON COLUMN t_album.album_cover IS '相册封面';
COMMENT ON COLUMN t_album.is_delete IS '是否删除';
COMMENT ON COLUMN t_album.status IS '状态值 1公开 2私密';
COMMENT ON COLUMN t_album.created_at IS '创建时间';
COMMENT ON COLUMN t_album.updated_at IS '更新时间';

CREATE TABLE t_api (
    id bigserial PRIMARY KEY,
    parent_id integer NOT NULL DEFAULT 0,
    name varchar(128) NOT NULL DEFAULT '',
    path varchar(128) NOT NULL DEFAULT '',
    method varchar(16) NOT NULL DEFAULT '',
    traceable smallint NOT NULL DEFAULT 0,
    status smallint NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_api_path_method_name UNIQUE (path, method, name)
);
COMMENT ON TABLE t_api IS '接口';
COMMENT ON COLUMN t_api.id IS '主键id';
COMMENT ON COLUMN t_api.parent_id IS '分组id';
COMMENT ON COLUMN t_api.name IS 'api名称';
COMMENT ON COLUMN t_api.path IS 'api路径';
COMMENT ON COLUMN t_api.method IS 'api请求方法';
COMMENT ON COLUMN t_api.traceable IS '是否追溯操作记录 0需要，1是';
COMMENT ON COLUMN t_api.status IS '是否禁用 0否 1是';
COMMENT ON COLUMN t_api.created_at IS '创建时间';
COMMENT ON COLUMN t_api.updated_at IS '更新时间';

CREATE TABLE t_article (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    category_id integer NOT NULL DEFAULT 0,
    article_cover varchar(1024) NOT NULL DEFAULT '',
    article_title varchar(64) NOT NULL DEFAULT '',
    article_content text NOT NULL,
    article_type smallint NOT NULL DEFAULT 0,
    original_url varchar(255) NOT NULL DEFAULT '',
    tags text[] NOT NULL DEFAULT '{}',
    metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
    is_top boolean NOT NULL DEFAULT false,
    is_delete boolean NOT NULL DEFAULT false,
    status smallint NOT NULL DEFAULT 1,
    like_count integer NOT NULL DEFAULT 0,
    view_count integer NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_article IS '文章';
COMMENT ON COLUMN t_article.id IS 'id';
COMMENT ON COLUMN t_article.user_id IS '作者';
COMMENT ON COLUMN t_article.category_id IS '文章分类';
COMMENT ON COLUMN t_article.article_cover IS '文章缩略图';
COMMENT ON COLUMN t_article.article_title IS '标题';
COMMENT ON COLUMN t_article.article_content IS '内容';
COMMENT ON COLUMN t_article.article_type IS '文章类型 1原创 2转载 3翻译';
COMMENT ON COLUMN t_article.original_url IS '原文链接';
COMMENT ON COLUMN t_article.tags IS '标签列表';
COMMENT ON COLUMN t_article.metadata IS '文章扩展元数据';
COMMENT ON COLUMN t_article.is_top IS '是否置顶 0否 1是';
COMMENT ON COLUMN t_article.is_delete IS '是否删除 0否 1是';
COMMENT ON COLUMN t_article.status IS '状态值 1公开 2私密 3草稿 4评论可见';
COMMENT ON COLUMN t_article.like_count IS '点赞数';
COMMENT ON COLUMN t_article.view_count IS '查看数';
COMMENT ON COLUMN t_article.created_at IS '发表时间';
COMMENT ON COLUMN t_article.updated_at IS '更新时间';

CREATE TABLE t_category (
    id bigserial PRIMARY KEY,
    category_name varchar(32) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_category_name UNIQUE (category_name)
);
COMMENT ON TABLE t_category IS '文章分类';
COMMENT ON COLUMN t_category.id IS 'id';
COMMENT ON COLUMN t_category.category_name IS '分类名';
COMMENT ON COLUMN t_category.created_at IS '创建时间';
COMMENT ON COLUMN t_category.updated_at IS '更新时间';

CREATE TABLE t_comment (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    terminal_id varchar(64) NOT NULL DEFAULT '',
    topic_id integer NOT NULL DEFAULT 0,
    parent_id integer NOT NULL DEFAULT 0,
    reply_id integer NOT NULL DEFAULT 0,
    reply_user_id varchar(255) NOT NULL DEFAULT '',
    comment_content text NOT NULL,
    type smallint NOT NULL DEFAULT 0,
    status smallint NOT NULL DEFAULT 0,
    like_count integer NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_comment IS '评论';
COMMENT ON COLUMN t_comment.id IS '主键';
COMMENT ON COLUMN t_comment.user_id IS '评论用户id';
COMMENT ON COLUMN t_comment.terminal_id IS '设备id';
COMMENT ON COLUMN t_comment.topic_id IS '主题id';
COMMENT ON COLUMN t_comment.parent_id IS '父评论id';
COMMENT ON COLUMN t_comment.reply_id IS '回复评论id';
COMMENT ON COLUMN t_comment.reply_user_id IS '评论回复用户id';
COMMENT ON COLUMN t_comment.comment_content IS '评论内容';
COMMENT ON COLUMN t_comment.type IS '评论类型 1.文章 2.友链 3.说说';
COMMENT ON COLUMN t_comment.status IS '状态 0.正常 1.已编辑 2.已删除';
COMMENT ON COLUMN t_comment.like_count IS '评论点赞数量';
COMMENT ON COLUMN t_comment.created_at IS '创建时间';
COMMENT ON COLUMN t_comment.updated_at IS '更新时间';

CREATE TABLE t_file_log (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    terminal_id varchar(64) NOT NULL DEFAULT '',
    file_path varchar(128) NOT NULL DEFAULT '',
    file_name varchar(128) NOT NULL DEFAULT '',
    file_type varchar(128) NOT NULL DEFAULT '',
    file_size integer NOT NULL DEFAULT 0,
    file_md5 varchar(128) NOT NULL DEFAULT '',
    file_url varchar(256) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_file_log IS '文件上传记录';
COMMENT ON COLUMN t_file_log.id IS 'id';
COMMENT ON COLUMN t_file_log.user_id IS '用户id';
COMMENT ON COLUMN t_file_log.terminal_id IS '设备id';
COMMENT ON COLUMN t_file_log.file_path IS '文件路径';
COMMENT ON COLUMN t_file_log.file_name IS '文件名称';
COMMENT ON COLUMN t_file_log.file_type IS '文件类型';
COMMENT ON COLUMN t_file_log.file_size IS '文件大小';
COMMENT ON COLUMN t_file_log.file_md5 IS '文件md5值';
COMMENT ON COLUMN t_file_log.file_url IS '上传路径';
COMMENT ON COLUMN t_file_log.created_at IS '创建时间';
COMMENT ON COLUMN t_file_log.updated_at IS '更新时间';

CREATE TABLE t_friend (
    id bigserial PRIMARY KEY,
    link_name varchar(32) NOT NULL DEFAULT '',
    link_avatar varchar(255) NOT NULL DEFAULT '',
    link_address varchar(64) NOT NULL DEFAULT '',
    link_intro varchar(100) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_friend_name UNIQUE (link_name)
);
COMMENT ON TABLE t_friend IS '友链';
COMMENT ON COLUMN t_friend.id IS 'id';
COMMENT ON COLUMN t_friend.link_name IS '链接名';
COMMENT ON COLUMN t_friend.link_avatar IS '链接头像';
COMMENT ON COLUMN t_friend.link_address IS '链接地址';
COMMENT ON COLUMN t_friend.link_intro IS '链接介绍';
COMMENT ON COLUMN t_friend.created_at IS '创建时间';
COMMENT ON COLUMN t_friend.updated_at IS '更新时间';

CREATE TABLE t_login_log (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    terminal_id varchar(64) NOT NULL DEFAULT '',
    login_type varchar(64) NOT NULL DEFAULT '',
    app_name varchar(64) NOT NULL DEFAULT '',
    login_at timestamptz NOT NULL DEFAULT now(),
    logout_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_login_log IS '用户登录历史';
COMMENT ON COLUMN t_login_log.id IS 'id';
COMMENT ON COLUMN t_login_log.user_id IS '用户id';
COMMENT ON COLUMN t_login_log.terminal_id IS '设备id';
COMMENT ON COLUMN t_login_log.login_type IS '登录类型';
COMMENT ON COLUMN t_login_log.app_name IS 'app名称';
COMMENT ON COLUMN t_login_log.login_at IS '登录时间';
COMMENT ON COLUMN t_login_log.logout_at IS '登出时间';
COMMENT ON COLUMN t_login_log.created_at IS '创建时间';
COMMENT ON COLUMN t_login_log.updated_at IS '更新时间';

CREATE TABLE t_menu (
    id bigserial PRIMARY KEY,
    parent_id integer NOT NULL DEFAULT 0,
    path varchar(64) NOT NULL DEFAULT '',
    name varchar(64) NOT NULL DEFAULT '',
    component varchar(256) NOT NULL DEFAULT '',
    redirect varchar(256) NOT NULL DEFAULT '',
    type varchar(64) NOT NULL DEFAULT '0',
    title varchar(64) NOT NULL DEFAULT '',
    icon varchar(64) NOT NULL DEFAULT '',
    rank integer NOT NULL DEFAULT 0,
    perm varchar(64) NOT NULL DEFAULT '',
    params varchar(256) NOT NULL DEFAULT '',
    keep_alive boolean NOT NULL DEFAULT false,
    always_show boolean NOT NULL DEFAULT false,
    visible boolean NOT NULL DEFAULT false,
    status boolean NOT NULL DEFAULT false,
    extra jsonb NOT NULL DEFAULT '{}'::jsonb,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_menu_path_perm UNIQUE (path, perm)
);
COMMENT ON TABLE t_menu IS '菜单';
COMMENT ON COLUMN t_menu.id IS '主键';
COMMENT ON COLUMN t_menu.parent_id IS '父id';
COMMENT ON COLUMN t_menu.path IS '路由路径';
COMMENT ON COLUMN t_menu.name IS '路由名称';
COMMENT ON COLUMN t_menu.component IS '路由组件';
COMMENT ON COLUMN t_menu.redirect IS '路由重定向';
COMMENT ON COLUMN t_menu.type IS '菜单类型';
COMMENT ON COLUMN t_menu.title IS '菜单标题';
COMMENT ON COLUMN t_menu.icon IS '菜单图标';
COMMENT ON COLUMN t_menu.rank IS '排序';
COMMENT ON COLUMN t_menu.perm IS '权限标识';
COMMENT ON COLUMN t_menu.params IS '路由参数';
COMMENT ON COLUMN t_menu.keep_alive IS '是否缓存';
COMMENT ON COLUMN t_menu.always_show IS '是否一直显示菜单';
COMMENT ON COLUMN t_menu.visible IS '菜单是否可见';
COMMENT ON COLUMN t_menu.status IS '是否禁用';
COMMENT ON COLUMN t_menu.extra IS '菜单元数据';
COMMENT ON COLUMN t_menu.created_at IS '创建时间';
COMMENT ON COLUMN t_menu.updated_at IS '更新时间';

CREATE TABLE t_message (
    id bigserial PRIMARY KEY,
    user_id varchar(255) NOT NULL DEFAULT '0',
    terminal_id varchar(255) NOT NULL DEFAULT '',
    message_content varchar(255) NOT NULL DEFAULT '',
    status integer NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_message IS '留言';
COMMENT ON COLUMN t_message.id IS '主键id';
COMMENT ON COLUMN t_message.user_id IS '用户id';
COMMENT ON COLUMN t_message.terminal_id IS '终端id';
COMMENT ON COLUMN t_message.message_content IS '留言内容';
COMMENT ON COLUMN t_message.status IS '状态:0正常 1编辑 2撤回 3删除';
COMMENT ON COLUMN t_message.created_at IS '发布时间';
COMMENT ON COLUMN t_message.updated_at IS '更新时间';

CREATE TABLE t_operation_log (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    terminal_id varchar(64) NOT NULL DEFAULT '',
    opt_module varchar(32) NOT NULL DEFAULT '',
    opt_desc varchar(255) NOT NULL DEFAULT '',
    request_uri varchar(255) NOT NULL DEFAULT '',
    request_method varchar(32) NOT NULL DEFAULT '',
    request_data text NOT NULL DEFAULT '',
    response_data text NOT NULL DEFAULT '',
    response_status integer NOT NULL DEFAULT 0,
    cost varchar(32) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_operation_log IS '操作记录';
COMMENT ON COLUMN t_operation_log.id IS '主键id';
COMMENT ON COLUMN t_operation_log.user_id IS '用户id';
COMMENT ON COLUMN t_operation_log.terminal_id IS '设备id';
COMMENT ON COLUMN t_operation_log.opt_module IS '操作模块';
COMMENT ON COLUMN t_operation_log.opt_desc IS '操作描述';
COMMENT ON COLUMN t_operation_log.request_uri IS '请求地址';
COMMENT ON COLUMN t_operation_log.request_method IS '请求方式';
COMMENT ON COLUMN t_operation_log.request_data IS '请求参数';
COMMENT ON COLUMN t_operation_log.response_data IS '返回数据';
COMMENT ON COLUMN t_operation_log.response_status IS '响应状态码';
COMMENT ON COLUMN t_operation_log.cost IS '耗时（ms）';
COMMENT ON COLUMN t_operation_log.created_at IS '创建时间';
COMMENT ON COLUMN t_operation_log.updated_at IS '更新时间';

CREATE TABLE t_page (
    id bigserial PRIMARY KEY,
    page_name varchar(32) NOT NULL DEFAULT '',
    page_label varchar(32) NOT NULL DEFAULT '',
    page_cover varchar(255) NOT NULL DEFAULT '',
    is_carousel boolean NOT NULL DEFAULT false,
    carousel_covers varchar(1024) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_page IS '页面';
COMMENT ON COLUMN t_page.id IS '页面id';
COMMENT ON COLUMN t_page.page_name IS '页面名';
COMMENT ON COLUMN t_page.page_label IS '页面标签';
COMMENT ON COLUMN t_page.page_cover IS '页面封面';
COMMENT ON COLUMN t_page.is_carousel IS '是否轮播';
COMMENT ON COLUMN t_page.carousel_covers IS '轮播图片列表';
COMMENT ON COLUMN t_page.created_at IS '创建时间';
COMMENT ON COLUMN t_page.updated_at IS '更新时间';

CREATE TABLE t_photo (
    id bigserial PRIMARY KEY,
    album_id integer NOT NULL DEFAULT 0,
    photo_name varchar(64) NOT NULL DEFAULT '',
    photo_desc varchar(128) NOT NULL DEFAULT '',
    photo_src varchar(255) NOT NULL DEFAULT '',
    is_delete boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_photo IS '照片';
COMMENT ON COLUMN t_photo.id IS '主键';
COMMENT ON COLUMN t_photo.album_id IS '相册id';
COMMENT ON COLUMN t_photo.photo_name IS '照片名';
COMMENT ON COLUMN t_photo.photo_desc IS '照片描述';
COMMENT ON COLUMN t_photo.photo_src IS '照片地址';
COMMENT ON COLUMN t_photo.is_delete IS '是否删除';
COMMENT ON COLUMN t_photo.created_at IS '创建时间';
COMMENT ON COLUMN t_photo.updated_at IS '更新时间';

CREATE TABLE t_role (
    id bigserial PRIMARY KEY,
    role_key varchar(64) NOT NULL DEFAULT '',
    role_comment varchar(64) NOT NULL DEFAULT '',
    status smallint NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_role_key UNIQUE (role_key),
    CONSTRAINT ck_role_key_enum CHECK (role_key IN ('root', 'super_admin', 'visitor'))
);
COMMENT ON TABLE t_role IS '角色';
COMMENT ON COLUMN t_role.id IS '主键id';
COMMENT ON COLUMN t_role.role_key IS '角色枚举名';
COMMENT ON COLUMN t_role.role_comment IS '角色备注';
COMMENT ON COLUMN t_role.status IS '状态 0正常 1禁用';
COMMENT ON COLUMN t_role.created_at IS '创建时间';
COMMENT ON COLUMN t_role.updated_at IS '更新时间';

CREATE TABLE t_role_api (
    id bigserial PRIMARY KEY,
    role_id integer NOT NULL DEFAULT 0,
    api_id integer NOT NULL DEFAULT 0
);
COMMENT ON TABLE t_role_api IS '角色-api关联';
COMMENT ON COLUMN t_role_api.id IS '主键id';
COMMENT ON COLUMN t_role_api.role_id IS '角色id';
COMMENT ON COLUMN t_role_api.api_id IS '接口id';

CREATE TABLE t_role_menu (
    id bigserial PRIMARY KEY,
    role_id integer NOT NULL DEFAULT 0,
    menu_id integer NOT NULL DEFAULT 0
);
COMMENT ON TABLE t_role_menu IS '角色-菜单关联';
COMMENT ON COLUMN t_role_menu.id IS '主键id';
COMMENT ON COLUMN t_role_menu.role_id IS '角色id';
COMMENT ON COLUMN t_role_menu.menu_id IS '菜单id';

CREATE TABLE t_system_notice (
    id bigserial PRIMARY KEY,
    title varchar(128) NOT NULL DEFAULT '',
    content text NOT NULL,
    type varchar(16) NOT NULL DEFAULT 'system',
    level varchar(16) NOT NULL DEFAULT 'info',
    app_name varchar(64) NOT NULL DEFAULT '',
    publisher_id varchar(64) NOT NULL DEFAULT '',
    publish_status smallint NOT NULL DEFAULT 1,
    publish_time timestamptz,
    revoke_time timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_system_notice IS '系统通知表';
COMMENT ON COLUMN t_system_notice.id IS '主键ID';
COMMENT ON COLUMN t_system_notice.title IS '通知标题';
COMMENT ON COLUMN t_system_notice.content IS '通知内容';
COMMENT ON COLUMN t_system_notice.type IS '通知类型 system-系统公告 maintenance-维护通知 update-功能更新 remind-重要提醒';
COMMENT ON COLUMN t_system_notice.level IS '通知等级 info-普通 notice-提醒 warning-警告 error-紧急';
COMMENT ON COLUMN t_system_notice.app_name IS '目标应用名称（如：blog-前台、admin-后台）';
COMMENT ON COLUMN t_system_notice.publisher_id IS '发布人ID';
COMMENT ON COLUMN t_system_notice.publish_status IS '发布状态 1-草稿 2-已发布 3-已撤回';
COMMENT ON COLUMN t_system_notice.publish_time IS '发布时间';
COMMENT ON COLUMN t_system_notice.revoke_time IS '撤回时间';
COMMENT ON COLUMN t_system_notice.created_at IS '创建时间';
COMMENT ON COLUMN t_system_notice.updated_at IS '更新时间';

CREATE TABLE t_tag (
    id bigserial PRIMARY KEY,
    tag_name varchar(32) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_tag_name UNIQUE (tag_name)
);
COMMENT ON TABLE t_tag IS '标签';
COMMENT ON COLUMN t_tag.id IS 'id';
COMMENT ON COLUMN t_tag.tag_name IS '标签名';
COMMENT ON COLUMN t_tag.created_at IS '创建时间';
COMMENT ON COLUMN t_tag.updated_at IS '更新时间';

CREATE TABLE t_talk (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    content varchar(2048) NOT NULL DEFAULT '',
    images text[] NOT NULL DEFAULT '{}',
    is_top boolean NOT NULL DEFAULT false,
    status smallint NOT NULL DEFAULT 1,
    like_count integer NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_talk IS '说说';
COMMENT ON COLUMN t_talk.id IS '说说id';
COMMENT ON COLUMN t_talk.user_id IS '用户id';
COMMENT ON COLUMN t_talk.content IS '说说内容';
COMMENT ON COLUMN t_talk.images IS '图片';
COMMENT ON COLUMN t_talk.is_top IS '是否置顶';
COMMENT ON COLUMN t_talk.status IS '状态 1.公开 2.私密';
COMMENT ON COLUMN t_talk.like_count IS '点赞数';
COMMENT ON COLUMN t_talk.created_at IS '创建时间';
COMMENT ON COLUMN t_talk.updated_at IS '更新时间';

CREATE TABLE t_user (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    username varchar(64) NOT NULL DEFAULT '',
    password varchar(128) NOT NULL DEFAULT '',
    nickname varchar(64) NOT NULL DEFAULT '',
    avatar varchar(255) NOT NULL DEFAULT '',
    email varchar(64) NOT NULL DEFAULT '',
    phone varchar(64) NOT NULL DEFAULT '',
    info varchar(1024) NOT NULL DEFAULT '',
    status smallint NOT NULL DEFAULT 0,
    register_type varchar(64) NOT NULL DEFAULT '',
    ip_address varchar(255) NOT NULL DEFAULT '',
    ip_source varchar(255) NOT NULL DEFAULT '',
    role_id integer NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_user_uid UNIQUE (user_id),
    CONSTRAINT uk_user_username UNIQUE (username)
);
COMMENT ON TABLE t_user IS '用户登录信息';
COMMENT ON COLUMN t_user.id IS 'id';
COMMENT ON COLUMN t_user.user_id IS '用户id';
COMMENT ON COLUMN t_user.username IS '用户名';
COMMENT ON COLUMN t_user.password IS '用户密码';
COMMENT ON COLUMN t_user.nickname IS '用户昵称';
COMMENT ON COLUMN t_user.avatar IS '用户头像';
COMMENT ON COLUMN t_user.email IS '邮箱';
COMMENT ON COLUMN t_user.phone IS '手机号';
COMMENT ON COLUMN t_user.info IS '用户信息';
COMMENT ON COLUMN t_user.status IS '状态: -1删除 0正常 1禁用';
COMMENT ON COLUMN t_user.register_type IS '注册方式';
COMMENT ON COLUMN t_user.ip_address IS '注册ip';
COMMENT ON COLUMN t_user.ip_source IS '注册ip 源';
COMMENT ON COLUMN t_user.role_id IS '用户角色id';
COMMENT ON COLUMN t_user.created_at IS '创建时间';
COMMENT ON COLUMN t_user.updated_at IS '更新时间';

CREATE TABLE t_user_oauth (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    platform varchar(64) NOT NULL DEFAULT '',
    open_id varchar(128) NOT NULL DEFAULT '',
    nickname varchar(128) NOT NULL DEFAULT '',
    avatar varchar(256) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_user_oauth_uid_platform UNIQUE (user_id, platform),
    CONSTRAINT uk_user_oauth_openid_platform UNIQUE (open_id, platform)
);
COMMENT ON TABLE t_user_oauth IS '第三方登录信息';
COMMENT ON COLUMN t_user_oauth.id IS 'id';
COMMENT ON COLUMN t_user_oauth.user_id IS '用户id';
COMMENT ON COLUMN t_user_oauth.platform IS '平台:手机号、邮箱、微信、飞书';
COMMENT ON COLUMN t_user_oauth.open_id IS '第三方平台id，标识唯一用户';
COMMENT ON COLUMN t_user_oauth.nickname IS '第三方平台昵称';
COMMENT ON COLUMN t_user_oauth.avatar IS '第三方平台头像';
COMMENT ON COLUMN t_user_oauth.created_at IS '创建时间';
COMMENT ON COLUMN t_user_oauth.updated_at IS '更新时间';

CREATE TABLE t_visit_daily_stats (
    id bigserial PRIMARY KEY,
    date varchar(10) NOT NULL DEFAULT '',
    view_count integer NOT NULL DEFAULT 0,
    visit_type smallint NOT NULL DEFAULT 1,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_visit_daily_stats_date_type UNIQUE (date, visit_type)
);
COMMENT ON TABLE t_visit_daily_stats IS '页面访问数量';
COMMENT ON COLUMN t_visit_daily_stats.id IS 'id';
COMMENT ON COLUMN t_visit_daily_stats.date IS '日期';
COMMENT ON COLUMN t_visit_daily_stats.view_count IS '访问量';
COMMENT ON COLUMN t_visit_daily_stats.visit_type IS '1 访客数 2 浏览数';
COMMENT ON COLUMN t_visit_daily_stats.created_at IS '创建时间';
COMMENT ON COLUMN t_visit_daily_stats.updated_at IS '更新时间';

CREATE TABLE t_visit_log (
    id bigserial PRIMARY KEY,
    user_id varchar(64) NOT NULL DEFAULT '',
    terminal_id varchar(64) NOT NULL DEFAULT '',
    page_name varchar(64) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON TABLE t_visit_log IS '访问日志';
COMMENT ON COLUMN t_visit_log.id IS 'id';
COMMENT ON COLUMN t_visit_log.user_id IS '用户id';
COMMENT ON COLUMN t_visit_log.terminal_id IS '设备id';
COMMENT ON COLUMN t_visit_log.page_name IS '访问页面';
COMMENT ON COLUMN t_visit_log.created_at IS '创建时间';
COMMENT ON COLUMN t_visit_log.updated_at IS '更新时间';

CREATE TABLE t_visitor (
    id bigserial PRIMARY KEY,
    terminal_id varchar(64) NOT NULL DEFAULT '',
    os varchar(50) NOT NULL DEFAULT '',
    browser varchar(50) NOT NULL DEFAULT '',
    ip_address varchar(255) NOT NULL DEFAULT '',
    ip_source varchar(255) NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_visitor_terminal_id UNIQUE (terminal_id)
);
COMMENT ON TABLE t_visitor IS '访客';
COMMENT ON COLUMN t_visitor.id IS 'id';
COMMENT ON COLUMN t_visitor.terminal_id IS '设备id';
COMMENT ON COLUMN t_visitor.os IS '操作系统';
COMMENT ON COLUMN t_visitor.browser IS '浏览器';
COMMENT ON COLUMN t_visitor.ip_address IS '操作ip';
COMMENT ON COLUMN t_visitor.ip_source IS '操作地址';
COMMENT ON COLUMN t_visitor.created_at IS '创建时间';
COMMENT ON COLUMN t_visitor.updated_at IS '更新时间';

CREATE TABLE t_website_config (
    id bigserial PRIMARY KEY,
    key varchar(32) NOT NULL DEFAULT '',
    config jsonb NOT NULL DEFAULT '{}'::jsonb,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT uk_website_config_key UNIQUE (key)
);
COMMENT ON TABLE t_website_config IS '网站配置表';
COMMENT ON COLUMN t_website_config.id IS 'id';
COMMENT ON COLUMN t_website_config.key IS '关键词';
COMMENT ON COLUMN t_website_config.config IS '配置信息';
COMMENT ON COLUMN t_website_config.created_at IS '创建时间';
COMMENT ON COLUMN t_website_config.updated_at IS '更新时间';

CREATE INDEX idx_article_tags ON t_article USING gin (tags);
CREATE INDEX idx_article_fulltext ON t_article USING gin (to_tsvector('simple', article_title || ' ' || article_content));
CREATE INDEX idx_comment_parent ON t_comment (parent_id);
CREATE INDEX idx_file_log_user_id ON t_file_log (user_id);
CREATE INDEX idx_file_log_path ON t_file_log (file_path);
CREATE INDEX idx_login_log_user_id ON t_login_log (user_id);
CREATE INDEX idx_operation_log_user_id ON t_operation_log (user_id);
CREATE INDEX idx_system_notice_publish_status ON t_system_notice (publish_status);
CREATE INDEX idx_system_notice_type ON t_system_notice (type);
CREATE INDEX idx_system_notice_publish_time ON t_system_notice (publish_time);
CREATE INDEX idx_visit_log_user_id ON t_visit_log (user_id);
CREATE INDEX idx_website_config_jsonb ON t_website_config USING gin (config);

CREATE TRIGGER trg_t_album_set_updated_at BEFORE UPDATE ON t_album
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_api_set_updated_at BEFORE UPDATE ON t_api
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_article_set_updated_at BEFORE UPDATE ON t_article
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_category_set_updated_at BEFORE UPDATE ON t_category
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_comment_set_updated_at BEFORE UPDATE ON t_comment
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_file_log_set_updated_at BEFORE UPDATE ON t_file_log
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_friend_set_updated_at BEFORE UPDATE ON t_friend
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_login_log_set_updated_at BEFORE UPDATE ON t_login_log
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_menu_set_updated_at BEFORE UPDATE ON t_menu
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_message_set_updated_at BEFORE UPDATE ON t_message
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_operation_log_set_updated_at BEFORE UPDATE ON t_operation_log
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_page_set_updated_at BEFORE UPDATE ON t_page
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_photo_set_updated_at BEFORE UPDATE ON t_photo
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_role_set_updated_at BEFORE UPDATE ON t_role
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_system_notice_set_updated_at BEFORE UPDATE ON t_system_notice
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_tag_set_updated_at BEFORE UPDATE ON t_tag
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_talk_set_updated_at BEFORE UPDATE ON t_talk
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_user_set_updated_at BEFORE UPDATE ON t_user
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_user_oauth_set_updated_at BEFORE UPDATE ON t_user_oauth
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_visit_daily_stats_set_updated_at BEFORE UPDATE ON t_visit_daily_stats
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_visit_log_set_updated_at BEFORE UPDATE ON t_visit_log
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_visitor_set_updated_at BEFORE UPDATE ON t_visitor
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_t_website_config_set_updated_at BEFORE UPDATE ON t_website_config
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
