BEGIN;

INSERT INTO t_role (
    id,
    role_key,
    role_comment,
    status
) VALUES
    (1, 'root', 'System Owner', 0),
    (2, 'super_admin', 'super admin', 0),
    (3, 'user', 'default registered user', 0)
ON CONFLICT (id) DO UPDATE SET
    role_key = EXCLUDED.role_key,
    role_comment = EXCLUDED.role_comment,
    status = EXCLUDED.status,
    updated_at = now();

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
    ip_source,
    role_id
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
    'local',
    1
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
    role_id = EXCLUDED.role_id,
    updated_at = now();

SELECT setval(pg_get_serial_sequence('t_role', 'id'), COALESCE((SELECT MAX(id) FROM t_role), 1), true);

COMMIT;
