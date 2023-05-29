package main

import "strconv"

func BinToDec(match string) string {
	str := match[0 : len(match)-7]
	num, _ := strconv.ParseInt(str, 2, 64)
	return strconv.Itoa(int(num)) + " "
}
