package main

import "strconv"

func BinToDec(string_list []string, idx int) {
	num, _ := strconv.ParseInt(string_list[idx-1], 2, 64)
	string_list[idx-1] = strconv.Itoa(int(num))
}
