package main

import "strconv"

func HexToDec(match string) string {
	str := match[0 : len(match)-6]
	num, _ := strconv.ParseInt(str, 16, 64)
	return strconv.Itoa(int(num))
}
