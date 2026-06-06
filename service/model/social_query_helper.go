package model

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func buildSocialWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	return buildSocialWhereClauseWithStartIndex(conditions, 1, args...)
}

func buildSocialWhereClauseWithStartIndex(conditions string, start int, args ...interface{}) (string, []interface{}) {
	if conditions == "" {
		return "", args
	}

	replaced := strings.ReplaceAll(conditions, "id in (?)", "id = any(?)")
	replaced = strings.ReplaceAll(replaced, "topic_id in (?)", "topic_id = any(?)")

	normalizedArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		switch v := arg.(type) {
		case []int64:
			normalizedArgs = append(normalizedArgs, pq.Array(v))
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
