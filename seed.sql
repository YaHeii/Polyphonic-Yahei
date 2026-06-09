/*
 Minimal seed data for local admin/backend startup.

 Default admin credentials:
 - username: admin
 - password: 123456
*/

BEGIN;

-- ---------------------------------------------------------------------------
-- users / roles / permissions
-- ---------------------------------------------------------------------------

INSERT INTO t_user (
    user_id,
    username,
    password,
    nickname,
    avatar,
    email,
    phone,
    info,
    status,
    register_type,
    ip_address,
    ip_source
) VALUES (
    'admin-001',
    'admin',
    '$2a$10$DNSq5UwzRTn8Xco1c8.jku1gfz2YnL1vl4WjT5WASYljPPat7pYiO',
    'Administrator',
    '',
    'admin@example.com',
    '13800000000',
    '{"gender":1,"intro":"System administrator","website":"https://yahei.local"}',
    0,
    'username',
    '127.0.0.1',
    'local'
) ON CONFLICT (user_id) DO UPDATE SET
    username = EXCLUDED.username,
    password = EXCLUDED.password,
    nickname = EXCLUDED.nickname,
    avatar = EXCLUDED.avatar,
    email = EXCLUDED.email,
    phone = EXCLUDED.phone,
    info = EXCLUDED.info,
    status = EXCLUDED.status,
    register_type = EXCLUDED.register_type,
    ip_address = EXCLUDED.ip_address,
    ip_source = EXCLUDED.ip_source,
    updated_at = now();

INSERT INTO t_role (
    id,
    parent_id,
    role_key,
    role_label,
    role_comment,
    is_default,
    status
) VALUES (
    1,
    0,
    'super_admin',
    'Super Admin',
    'local seed super admin',
    true,
    0
) ON CONFLICT (id) DO UPDATE SET
    parent_id = EXCLUDED.parent_id,
    role_key = EXCLUDED.role_key,
    role_label = EXCLUDED.role_label,
    role_comment = EXCLUDED.role_comment,
    is_default = EXCLUDED.is_default,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO t_user_role (
    id,
    user_id,
    role_id
) VALUES (
    1,
    'admin-001',
    1
) ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    role_id = EXCLUDED.role_id;

INSERT INTO t_menu (
    id,
    parent_id,
    path,
    name,
    component,
    redirect,
    type,
    title,
    icon,
    rank,
    perm,
    params,
    keep_alive,
    always_show,
    visible,
    status,
    extra
) VALUES
    (1, 0, '/dashboard', 'Dashboard', 'dashboard/index', '', '1', '仪表盘', 'LayoutDashboard', 1, 'dashboard:view', '', true, false, true, false, '{}'::jsonb),
    (2, 0, '/system', 'System', 'Layout', '/system/account', '0', '系统管理', 'Settings', 2, 'system', '', false, true, true, false, '{}'::jsonb),
    (3, 2, 'account', 'Account', 'system/account/index', '', '1', '用户管理', 'Users', 1, 'account:list', '', true, false, true, false, '{}'::jsonb),
    (4, 2, 'role', 'Role', 'system/role/index', '', '1', '角色管理', 'Shield', 2, 'role:list', '', true, false, true, false, '{}'::jsonb),
    (5, 2, 'menu', 'Menu', 'system/menu/index', '', '1', '菜单管理', 'PanelsTopLeft', 3, 'menu:list', '', true, false, true, false, '{}'::jsonb),
    (6, 2, 'api', 'Api', 'system/api/index', '', '1', '接口管理', 'Waypoints', 4, 'api:list', '', true, false, true, false, '{}'::jsonb),
    (7, 0, '/content', 'Content', 'Layout', '/content/article', '0', '内容管理', 'Files', 3, 'content', '', false, true, true, false, '{}'::jsonb),
    (8, 7, 'article', 'Article', 'content/article/index', '', '1', '文章管理', 'FileText', 1, 'article:list', '', true, false, true, false, '{}'::jsonb),
    (9, 7, 'category', 'Category', 'content/category/index', '', '1', '分类管理', 'FolderTree', 2, 'category:list', '', true, false, true, false, '{}'::jsonb),
    (10, 7, 'tag', 'Tag', 'content/tag/index', '', '1', '标签管理', 'Tags', 3, 'tag:list', '', true, false, true, false, '{}'::jsonb),
    (11, 0, '/site', 'Site', 'Layout', '/site/config', '0', '站点管理', 'Globe', 4, 'site', '', false, true, true, false, '{}'::jsonb),
    (12, 11, 'config', 'Config', 'site/config/index', '', '1', '网站配置', 'SlidersHorizontal', 1, 'website:config', '', true, false, true, false, '{}'::jsonb),
    (13, 11, 'message', 'Message', 'site/message/index', '', '1', '留言管理', 'MessagesSquare', 2, 'message:list', '', true, false, true, false, '{}'::jsonb),
    (14, 11, 'notice', 'Notice', 'site/notice/index', '', '1', '通知管理', 'Bell', 3, 'notice:list', '', true, false, true, false, '{}'::jsonb)
ON CONFLICT (id) DO UPDATE SET
    parent_id = EXCLUDED.parent_id,
    path = EXCLUDED.path,
    name = EXCLUDED.name,
    component = EXCLUDED.component,
    redirect = EXCLUDED.redirect,
    type = EXCLUDED.type,
    title = EXCLUDED.title,
    icon = EXCLUDED.icon,
    rank = EXCLUDED.rank,
    perm = EXCLUDED.perm,
    params = EXCLUDED.params,
    keep_alive = EXCLUDED.keep_alive,
    always_show = EXCLUDED.always_show,
    visible = EXCLUDED.visible,
    status = EXCLUDED.status,
    extra = EXCLUDED.extra,
    updated_at = now();

INSERT INTO t_api (
    id,
    parent_id,
    name,
    path,
    method,
    traceable,
    status
) VALUES
    (1, 0, 'Ping', '/admin-api/v1/ping', 'GET', 0, 0),
    (2, 0, 'Login', '/admin-api/v1/login', 'POST', 0, 0),
    (3, 0, 'GetUserInfo', '/admin-api/v1/user/get_user_info', 'GET', 0, 0),
    (4, 0, 'GetUserMenus', '/admin-api/v1/user/get_user_menus', 'GET', 0, 0),
    (5, 0, 'GetUserRoles', '/admin-api/v1/user/get_user_roles', 'GET', 0, 0),
    (6, 0, 'GetUserApis', '/admin-api/v1/user/get_user_apis', 'GET', 0, 0),
    (7, 0, 'FindAccountList', '/admin-api/v1/account/find_account_list', 'POST', 1, 0),
    (8, 0, 'GetAdminHomeInfo', '/admin-api/v1/website/get_admin_home_info', 'GET', 0, 0),
    (9, 0, 'GetWebsiteConfig', '/admin-api/v1/website/get_website_config', 'GET', 0, 0),
    (10, 0, 'FindArticleList', '/admin-api/v1/article/find_article_list', 'POST', 1, 0),
    (11, 0, 'FindCategoryList', '/admin-api/v1/category/find_category_list', 'POST', 1, 0),
    (12, 0, 'FindTagList', '/admin-api/v1/tag/find_tag_list', 'POST', 1, 0),
    (13, 0, 'FindMessageList', '/admin-api/v1/message/find_message_list', 'POST', 1, 0),
    (14, 0, 'FindNoticeList', '/admin-api/v1/notice/find_notice_list', 'POST', 1, 0),
    (15, 0, 'UploadFile', '/admin-api/v1/upload/upload_file', 'POST', 1, 0),
    (16, 0, 'ListUploadFile', '/admin-api/v1/upload/list_upload_file', 'POST', 1, 0)
ON CONFLICT (id) DO UPDATE SET
    parent_id = EXCLUDED.parent_id,
    name = EXCLUDED.name,
    path = EXCLUDED.path,
    method = EXCLUDED.method,
    traceable = EXCLUDED.traceable,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO t_role_menu (
    id,
    role_id,
    menu_id
) VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3),
    (4, 1, 4),
    (5, 1, 5),
    (6, 1, 6),
    (7, 1, 7),
    (8, 1, 8),
    (9, 1, 9),
    (10, 1, 10),
    (11, 1, 11),
    (12, 1, 12),
    (13, 1, 13),
    (14, 1, 14)
ON CONFLICT (id) DO UPDATE SET
    role_id = EXCLUDED.role_id,
    menu_id = EXCLUDED.menu_id;

INSERT INTO t_role_api (
    id,
    role_id,
    api_id
) VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3),
    (4, 1, 4),
    (5, 1, 5),
    (6, 1, 6),
    (7, 1, 7),
    (8, 1, 8),
    (9, 1, 9),
    (10, 1, 10),
    (11, 1, 11),
    (12, 1, 12),
    (13, 1, 13),
    (14, 1, 14),
    (15, 1, 15),
    (16, 1, 16)
ON CONFLICT (id) DO UPDATE SET
    role_id = EXCLUDED.role_id,
    api_id = EXCLUDED.api_id;

-- ---------------------------------------------------------------------------
-- website configuration
-- ---------------------------------------------------------------------------

INSERT INTO t_website_config (
    "key",
    config
) VALUES (
    'website_config',
    '{
      "admin_url": "http://127.0.0.1:9091",
      "websocket_url": "ws://127.0.0.1:9091/ws",
      "tourist_avatar": "",
      "user_avatar": "",
      "website_feature": {
        "is_ai_assistant": 0,
        "is_music_player": 0,
        "is_comment_review": 0,
        "is_email_notice": 0,
        "is_message_review": 0,
        "is_reward": 0
      },
      "website_info": {
        "website_author": "Yahei",
        "website_avatar": "",
        "website_create_time": "2026-01-01",
        "website_intro": "Polyphonic Yahei local seed environment",
        "website_name": "Polyphonic Yahei",
        "website_notice": "Local development environment",
        "website_record_no": ""
      },
      "reward_qr_code": {
        "alipay_qr_code": "",
        "weixin_qr_code": ""
      },
      "social_login_list": [
        {
          "name": "GitHub",
          "platform": "github",
          "authorize_url": "",
          "enabled": false
        }
      ],
      "social_url_list": [
        {
          "name": "GitHub",
          "platform": "github",
          "link_url": "https://github.com/YaHeii",
          "enabled": true
        }
      ]
    }'::jsonb
) ON CONFLICT ("key") DO UPDATE SET
    config = EXCLUDED.config,
    updated_at = now();

INSERT INTO t_website_config (
    "key",
    config
) VALUES (
    'about_me',
    '{"content":"This is a local seed profile used for admin/backend integration."}'::jsonb
) ON CONFLICT ("key") DO UPDATE SET
    config = EXCLUDED.config,
    updated_at = now();

-- ---------------------------------------------------------------------------
-- content samples
-- ---------------------------------------------------------------------------

INSERT INTO t_category (
    id,
    category_name
) VALUES
    (1, 'Go'),
    (2, 'System Design')
ON CONFLICT (id) DO UPDATE SET
    category_name = EXCLUDED.category_name,
    updated_at = now();

INSERT INTO t_tag (
    id,
    tag_name
) VALUES
    (1, 'go-zero'),
    (2, 'postgres'),
    (3, 'docker')
ON CONFLICT (id) DO UPDATE SET
    tag_name = EXCLUDED.tag_name,
    updated_at = now();

INSERT INTO t_article (
    id,
    user_id,
    category_id,
    article_cover,
    article_title,
    article_content,
    article_type,
    original_url,
    tags,
    metadata,
    is_top,
    is_delete,
    status,
    like_count,
    view_count,
    created_at,
    updated_at
) VALUES
    (
        1,
        'admin-001',
        1,
        '',
        'Getting Started With go-zero',
        'Seed article for local admin compose testing.',
        1,
        '',
        ARRAY['go-zero', 'postgres']::text[],
        '{}'::jsonb,
        true,
        false,
        1,
        8,
        128,
        now() - interval '1 day',
        now() - interval '1 day'
    ),
    (
        2,
        'admin-001',
        2,
        '',
        'Docker Compose Notes',
        'A sample article used to verify category, tag and article queries.',
        1,
        '',
        ARRAY['docker']::text[],
        '{}'::jsonb,
        false,
        false,
        1,
        4,
        96,
        now(),
        now()
    )
ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    category_id = EXCLUDED.category_id,
    article_cover = EXCLUDED.article_cover,
    article_title = EXCLUDED.article_title,
    article_content = EXCLUDED.article_content,
    article_type = EXCLUDED.article_type,
    original_url = EXCLUDED.original_url,
    tags = EXCLUDED.tags,
    metadata = EXCLUDED.metadata,
    is_top = EXCLUDED.is_top,
    is_delete = EXCLUDED.is_delete,
    status = EXCLUDED.status,
    like_count = EXCLUDED.like_count,
    view_count = EXCLUDED.view_count,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;

INSERT INTO t_message (
    id,
    user_id,
    terminal_id,
    message_content,
    status,
    created_at,
    updated_at
) VALUES
    (1, 'admin-001', 'terminal-admin', 'Seeded message from admin user.', 0, now() - interval '2 hours', now() - interval '2 hours'),
    (2, '0', 'visitor-001', 'Seeded visitor message for list verification.', 0, now() - interval '1 hour', now() - interval '1 hour')
ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    terminal_id = EXCLUDED.terminal_id,
    message_content = EXCLUDED.message_content,
    status = EXCLUDED.status,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;

INSERT INTO t_system_notice (
    id,
    title,
    content,
    type,
    level,
    app_name,
    publisher_id,
    publish_status,
    publish_time
) VALUES (
    1,
    'Local Environment Ready',
    'This notice is seeded for admin notice list verification.',
    'system',
    'info',
    'admin',
    'admin-001',
    2,
    now()
) ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    content = EXCLUDED.content,
    type = EXCLUDED.type,
    level = EXCLUDED.level,
    app_name = EXCLUDED.app_name,
    publisher_id = EXCLUDED.publisher_id,
    publish_status = EXCLUDED.publish_status,
    publish_time = EXCLUDED.publish_time,
    updated_at = now();

INSERT INTO t_friend (
    id,
    link_name,
    link_avatar,
    link_address,
    link_intro
) VALUES (
    1,
    'Polyphonic Repo',
    '',
    'https://github.com/YaHeii/Polyphonic-Yahei',
    'Repository link seeded for friend CRUD.'
) ON CONFLICT (id) DO UPDATE SET
    link_name = EXCLUDED.link_name,
    link_avatar = EXCLUDED.link_avatar,
    link_address = EXCLUDED.link_address,
    link_intro = EXCLUDED.link_intro,
    updated_at = now();

INSERT INTO t_page (
    id,
    page_name,
    page_label,
    page_cover,
    is_carousel,
    carousel_covers
) VALUES (
    1,
    'about',
    'About',
    '',
    false,
    ''
) ON CONFLICT (id) DO UPDATE SET
    page_name = EXCLUDED.page_name,
    page_label = EXCLUDED.page_label,
    page_cover = EXCLUDED.page_cover,
    is_carousel = EXCLUDED.is_carousel,
    carousel_covers = EXCLUDED.carousel_covers,
    updated_at = now();

INSERT INTO t_album (
    id,
    album_name,
    album_desc,
    album_cover,
    is_delete,
    status
) VALUES (
    1,
    'Default Album',
    'Seed album for admin CRUD verification.',
    '',
    false,
    1
) ON CONFLICT (id) DO UPDATE SET
    album_name = EXCLUDED.album_name,
    album_desc = EXCLUDED.album_desc,
    album_cover = EXCLUDED.album_cover,
    is_delete = EXCLUDED.is_delete,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO t_photo (
    id,
    album_id,
    photo_name,
    photo_desc,
    photo_src,
    is_delete
) VALUES (
    1,
    1,
    'seed-photo',
    'Seed photo for local list verification.',
    '',
    false
) ON CONFLICT (id) DO UPDATE SET
    album_id = EXCLUDED.album_id,
    photo_name = EXCLUDED.photo_name,
    photo_desc = EXCLUDED.photo_desc,
    photo_src = EXCLUDED.photo_src,
    is_delete = EXCLUDED.is_delete,
    updated_at = now();

INSERT INTO t_talk (
    id,
    user_id,
    content,
    images,
    is_top,
    status,
    like_count
) VALUES (
    1,
    'admin-001',
    'Seeded talk content for admin list verification.',
    ARRAY[]::text[],
    false,
    1,
    2
) ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    content = EXCLUDED.content,
    images = EXCLUDED.images,
    is_top = EXCLUDED.is_top,
    status = EXCLUDED.status,
    like_count = EXCLUDED.like_count,
    updated_at = now();

INSERT INTO t_comment (
    id,
    user_id,
    terminal_id,
    topic_id,
    parent_id,
    reply_id,
    reply_user_id,
    comment_content,
    type,
    status,
    like_count
) VALUES (
    1,
    'admin-001',
    'terminal-admin',
    1,
    0,
    0,
    '',
    'Seed comment for article moderation list.',
    1,
    0,
    1
) ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    terminal_id = EXCLUDED.terminal_id,
    topic_id = EXCLUDED.topic_id,
    parent_id = EXCLUDED.parent_id,
    reply_id = EXCLUDED.reply_id,
    reply_user_id = EXCLUDED.reply_user_id,
    comment_content = EXCLUDED.comment_content,
    type = EXCLUDED.type,
    status = EXCLUDED.status,
    like_count = EXCLUDED.like_count,
    updated_at = now();

-- ---------------------------------------------------------------------------
-- logs / visitor / trend samples
-- ---------------------------------------------------------------------------

INSERT INTO t_visitor (
    id,
    terminal_id,
    os,
    browser,
    ip_address,
    ip_source
) VALUES (
    1,
    'visitor-001',
    'Linux',
    'Chrome',
    '127.0.0.1',
    'local'
) ON CONFLICT (id) DO UPDATE SET
    terminal_id = EXCLUDED.terminal_id,
    os = EXCLUDED.os,
    browser = EXCLUDED.browser,
    ip_address = EXCLUDED.ip_address,
    ip_source = EXCLUDED.ip_source,
    updated_at = now();

INSERT INTO t_visit_log (
    id,
    user_id,
    terminal_id,
    page_name,
    created_at,
    updated_at
) VALUES (
    1,
    'admin-001',
    'terminal-admin',
    '/article/1',
    now() - interval '3 hours',
    now() - interval '3 hours'
) ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    terminal_id = EXCLUDED.terminal_id,
    page_name = EXCLUDED.page_name,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;

INSERT INTO t_visit_daily_stats (
    id,
    date,
    view_count,
    visit_type
) VALUES
    (1, to_char(current_date - interval '1 day', 'YYYY-MM-DD'), 18, 1),
    (2, to_char(current_date - interval '1 day', 'YYYY-MM-DD'), 64, 2),
    (3, to_char(current_date, 'YYYY-MM-DD'), 24, 1),
    (4, to_char(current_date, 'YYYY-MM-DD'), 92, 2)
ON CONFLICT (id) DO UPDATE SET
    date = EXCLUDED.date,
    view_count = EXCLUDED.view_count,
    visit_type = EXCLUDED.visit_type,
    updated_at = now();

INSERT INTO t_login_log (
    id,
    user_id,
    terminal_id,
    login_type,
    app_name,
    login_at,
    logout_at,
    created_at,
    updated_at
) VALUES (
    1,
    'admin-001',
    'terminal-admin',
    'password',
    'admin-api',
    now() - interval '6 hours',
    now() - interval '5 hours',
    now() - interval '6 hours',
    now() - interval '5 hours'
) ON CONFLICT (id) DO UPDATE SET
    user_id = EXCLUDED.user_id,
    terminal_id = EXCLUDED.terminal_id,
    login_type = EXCLUDED.login_type,
    app_name = EXCLUDED.app_name,
    login_at = EXCLUDED.login_at,
    logout_at = EXCLUDED.logout_at,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;

-- ---------------------------------------------------------------------------
-- keep sequences in sync for explicit ids
-- ---------------------------------------------------------------------------

SELECT setval(pg_get_serial_sequence('t_role', 'id'), COALESCE((SELECT MAX(id) FROM t_role), 1), true);
SELECT setval(pg_get_serial_sequence('t_user_role', 'id'), COALESCE((SELECT MAX(id) FROM t_user_role), 1), true);
SELECT setval(pg_get_serial_sequence('t_menu', 'id'), COALESCE((SELECT MAX(id) FROM t_menu), 1), true);
SELECT setval(pg_get_serial_sequence('t_api', 'id'), COALESCE((SELECT MAX(id) FROM t_api), 1), true);
SELECT setval(pg_get_serial_sequence('t_role_menu', 'id'), COALESCE((SELECT MAX(id) FROM t_role_menu), 1), true);
SELECT setval(pg_get_serial_sequence('t_role_api', 'id'), COALESCE((SELECT MAX(id) FROM t_role_api), 1), true);
SELECT setval(pg_get_serial_sequence('t_category', 'id'), COALESCE((SELECT MAX(id) FROM t_category), 1), true);
SELECT setval(pg_get_serial_sequence('t_tag', 'id'), COALESCE((SELECT MAX(id) FROM t_tag), 1), true);
SELECT setval(pg_get_serial_sequence('t_article', 'id'), COALESCE((SELECT MAX(id) FROM t_article), 1), true);
SELECT setval(pg_get_serial_sequence('t_message', 'id'), COALESCE((SELECT MAX(id) FROM t_message), 1), true);
SELECT setval(pg_get_serial_sequence('t_system_notice', 'id'), COALESCE((SELECT MAX(id) FROM t_system_notice), 1), true);
SELECT setval(pg_get_serial_sequence('t_friend', 'id'), COALESCE((SELECT MAX(id) FROM t_friend), 1), true);
SELECT setval(pg_get_serial_sequence('t_page', 'id'), COALESCE((SELECT MAX(id) FROM t_page), 1), true);
SELECT setval(pg_get_serial_sequence('t_album', 'id'), COALESCE((SELECT MAX(id) FROM t_album), 1), true);
SELECT setval(pg_get_serial_sequence('t_photo', 'id'), COALESCE((SELECT MAX(id) FROM t_photo), 1), true);
SELECT setval(pg_get_serial_sequence('t_talk', 'id'), COALESCE((SELECT MAX(id) FROM t_talk), 1), true);
SELECT setval(pg_get_serial_sequence('t_comment', 'id'), COALESCE((SELECT MAX(id) FROM t_comment), 1), true);
SELECT setval(pg_get_serial_sequence('t_visitor', 'id'), COALESCE((SELECT MAX(id) FROM t_visitor), 1), true);
SELECT setval(pg_get_serial_sequence('t_visit_log', 'id'), COALESCE((SELECT MAX(id) FROM t_visit_log), 1), true);
SELECT setval(pg_get_serial_sequence('t_visit_daily_stats', 'id'), COALESCE((SELECT MAX(id) FROM t_visit_daily_stats), 1), true);
SELECT setval(pg_get_serial_sequence('t_login_log', 'id'), COALESCE((SELECT MAX(id) FROM t_login_log), 1), true);

COMMIT;
