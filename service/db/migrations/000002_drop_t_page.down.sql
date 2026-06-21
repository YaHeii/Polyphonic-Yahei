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

CREATE TRIGGER trg_t_page_set_updated_at BEFORE UPDATE ON t_page
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_t_tag_set_updated_at BEFORE UPDATE ON t_tag
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
