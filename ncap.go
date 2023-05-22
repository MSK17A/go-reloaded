package main

import (
	"strconv"
	"strings"
)

func nCapitalize(string_list []string, idx int) {
	num_str := strings.Split(string_list[idx+1], ")")
	num, _ := strconv.Atoi(num_str[0])

	for i := idx; i > idx-num; i-- {
		str := string_list[i-1]
		string_list[i-1] = strings.ToUpper(string(str[0])) + str[1:]
	}
}
