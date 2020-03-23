package utils

import "strings"

func StrConcat(str ...string) string {
	return strings.Join(str, "")
}
