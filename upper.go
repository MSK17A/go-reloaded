package main

import "strings"

func Upper(string_list []string, idx int) {
	string_list[idx-1] = strings.ToUpper(string_list[idx-1]) // Same as above, but convert all the word to upperCase
}
