BEGIN;

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

SELECT setval(pg_get_serial_sequence('t_role', 'id'), COALESCE((SELECT MAX(id) FROM t_role), 1), true);
SELECT setval(pg_get_serial_sequence('t_user_role', 'id'), COALESCE((SELECT MAX(id) FROM t_user_role), 1), true);

COMMIT;
