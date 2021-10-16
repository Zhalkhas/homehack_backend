package utils

import (
	"strings"
	"unicode"
)

func ExtractPrice(str string) string {
	builder := strings.Builder{}
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
