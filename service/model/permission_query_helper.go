package model

import (
	"database/sql"
	"fmt"
	"strings"
)

func buildPermissionWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	return buildPermissionWhereClauseWithStartIndex(conditions, 1, args...)
}

func buildPermissionWhereClauseWithStartIndex(conditions string, start int, args ...interface{}) (string, []interface{}) {
	if conditions == "" {
		return "", args
	}

	replaced := strings.ReplaceAll(conditions, "id in (?)", "id = any(?)")
	normalizedArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		switch v := arg.(type) {
		case []int64:
			normalizedArgs = append(normalizedArgs, v)
		default:
			normalizedArgs = append(normalizedArgs, arg)
		}
	}

	var builder strings.Builder
	index := start
	for _, ch := range replaced {
		if ch == '?' {
			builder.WriteString(fmt.Sprintf("$%d", index))
			index++
			continue
		}
		builder.WriteRune(ch)
	}

	return builder.String(), normalizedArgs
}

func rowsAffected(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func buildPermissionIDArray(ids []int64) interface{} {
	if len(ids) == 0 {
		return []int64{}
	}
	return ids
}
