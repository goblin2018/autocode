package g

import (
	"strings"
	"unicode"
)

func toLowerAndUnderscore(s string) string {
	var buf strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			buf.WriteByte('_')
		}
		buf.WriteRune(unicode.ToLower(r))
	}
	return buf.String()
}
