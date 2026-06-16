BEGIN;

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

COMMIT;
