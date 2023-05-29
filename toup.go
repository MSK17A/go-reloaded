package main

import "strings"

func ToUp(match string) string {
	return strings.ToUpper(match[0 : len(match)-6]) // Trims "(up) "
}
