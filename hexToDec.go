package main

import "strconv"

func HexToDec(string_list []string, idx int) {
	num, _ := strconv.ParseInt(string_list[idx-1], 16, 64) // Take the number in hex, convert it to decimal integer
	string_list[idx-1] = strconv.Itoa(int(num))            // Conver the number to string
}
