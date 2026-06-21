DROP TRIGGER IF EXISTS trg_t_page_set_updated_at ON t_page;
DROP TRIGGER IF EXISTS trg_t_tag_set_updated_at ON t_tag;

DELETE FROM t_role_api
WHERE api_id IN (
    SELECT id
    FROM t_api
    WHERE path = '/admin-api/v1/tag/find_tag_list'
);

DELETE FROM t_api
WHERE path = '/admin-api/v1/tag/find_tag_list';

DELETE FROM t_role_menu
WHERE menu_id IN (
    SELECT id
    FROM t_menu
    WHERE parent_id = 7 AND path = 'tag'
);

DELETE FROM t_menu
WHERE parent_id = 7 AND path = 'tag';

DROP TABLE IF EXISTS t_page;
DROP TABLE IF EXISTS t_tag;
