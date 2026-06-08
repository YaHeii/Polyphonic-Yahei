package model

import (
	"fmt"
	"strings"
)

func buildSyslogWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	return buildSyslogWhereClauseWithStartIndex(conditions, 1, args...)
}

func buildSyslogWhereClauseWithStartIndex(conditions string, start int, args ...interface{}) (string, []interface{}) {
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
