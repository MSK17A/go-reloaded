package main

import (
	"strconv"
	"strings"
)

func nUp(string_list []string, idx int) {
	num_str := strings.Split(string_list[idx+1], ")") // Seperate the number from the bracket ')'. (low, <number>)
	num, _ := strconv.Atoi(num_str[0])                // Convert the number to integer, so we can use it to lowerCase number of strings behind

	for i := idx; i > idx-num; i-- {
		string_list[i-1] = strings.ToLower(string_list[i-1])
	}

	string_list[idx] = ""
	string_list[idx+1] = ""
	/*if unicode.IsPunct(rune(string_list[idx+2][0])) {
		string_list[idx+2] = ""
		string_list[idx-1] += ","
	}*/
}
