BEGIN;

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
    (13, 1, 13),
    (14, 1, 14),
    (15, 1, 15),
    (16, 1, 16)
ON CONFLICT (id) DO UPDATE SET
    role_id = EXCLUDED.role_id,
    api_id = EXCLUDED.api_id;

SELECT setval(pg_get_serial_sequence('t_menu', 'id'), COALESCE((SELECT MAX(id) FROM t_menu), 1), true);
SELECT setval(pg_get_serial_sequence('t_api', 'id'), COALESCE((SELECT MAX(id) FROM t_api), 1), true);
SELECT setval(pg_get_serial_sequence('t_role_menu', 'id'), COALESCE((SELECT MAX(id) FROM t_role_menu), 1), true);
SELECT setval(pg_get_serial_sequence('t_role_api', 'id'), COALESCE((SELECT MAX(id) FROM t_role_api), 1), true);

COMMIT;
