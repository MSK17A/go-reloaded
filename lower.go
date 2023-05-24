package main

import "strings"

func Lower(string_list []string, idx int) {
	string_list[idx-1] = strings.ToLower(string_list[idx-1])
}
