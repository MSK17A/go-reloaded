package pkgs

import "strings"

func ToLow(match string) string {
	return strings.ToLower(match[0 : len(match)-6]) // Trims "(low) "
}
