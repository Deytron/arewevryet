package utils

import "strings"

func Sanstr(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
