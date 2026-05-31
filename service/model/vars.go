package model

import (
	"errors"

	// 核心桥接：利用空白导入（_）将 pgx 驱动注册进 database/sql 标准库
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

// 统一包装，防止底层 SQL 错误在 logic 层引起跨层级污染
var ErrInvalidMetadata = errors.New("metadata 字段解析失败，数据格式不兼容")
