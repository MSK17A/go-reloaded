package main

import "strings"

func Capitalize(string_list []string, idx int) {
	str := string_list[idx-1]
	string_list[idx-1] = strings.ToUpper(string(str[0])) + str[1:] // Select the first letter and upper case it then add it to the string
	string_list[idx] = ""                                          // Replece with empty string to mark it as deleted
}
