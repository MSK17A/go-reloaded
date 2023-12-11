package pkgs

import (
	"strings"
)

func NUp(match string) string {
	matches := strings.Fields(match)
	result := ""
	for idx := range matches {
		result += strings.ToUpper(matches[idx]) + " "
	}
	return result
}
