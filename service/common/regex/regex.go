package regex

import "strings"

func IsRegex(s string) bool {
	return strings.HasPrefix(s, "^") && strings.HasSuffix(s, "$")
}
