package uri

import (
	strs "strings"
)

func ParseQueryStr(queryStr string) *map[string]string {
	queryParams := make(map[string]string)

	for item := range strs.SplitSeq(queryStr, "&") {
		str := strs.Split(item, "=")

		if len(str) == 2 {
			queryParams[str[0]] = str[1]
		}
	}

	return &queryParams
}