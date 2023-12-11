package pkgs

import (
	"strings"
)

func NLow(match string) string {
	matches := strings.Fields(match)
	result := ""
	for idx := range matches {
		result += strings.ToLower(matches[idx]) + " "
	}
	return result
}
