package pkgs

import (
	"strings"
)

func NLow(match string) string {
	matches := strings.Fields(match)
	result := ""
	for idx := range matches {
		result += strings.ToLower(string(matches[idx][0])) + matches[idx][1:] + " "
	}
	return result
}
