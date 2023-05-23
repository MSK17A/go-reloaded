package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file_name := os.Args[1]                        // Reading the first argument after program name
	dat, _ := os.ReadFile(file_name)               // Read the file and get the data as byte[]
	string_list := strings.Split(string(dat), " ") // Conv string to a List of strings
	fmt.Println(string_list)

	for idx, word := range string_list {
		if word == "(cap)" {
			Capitalize(string_list, idx)
			string_list[idx] = "" // Replece with empty string to mark it as deleted
		} else if word == "(up)" {
			Upper(string_list, idx)
			string_list[idx] = ""
		} else if word == "(low)" {
			Lower(string_list, idx)
			string_list[idx] = ""
		} else if word == "(hex)" {
			HexToDec(string_list, idx)
			string_list[idx] = ""
		} else if word == "(bin)" {
			BinToDec(string_list, idx)
			string_list[idx] = ""
		} else if word == "(low," {
			nLow(string_list, idx)
			string_list[idx] = ""
			string_list[idx+1] = ""
		} else if word == "(up," {
			nUp(string_list, idx)
			string_list[idx] = ""
		} else if word == "(cap," {
			nCapitalize(string_list, idx)
			string_list[idx] = ""

		} else if word == "," {
			string_list[idx-1] = ""
			string_list[idx] = ""
			string_list[idx-2] += ","
		}
	}

	str_out := ""

	for _, word := range string_list {
		if word != "" {
			str_out += word
			str_out += " "
		}
	}

	byte_str := []byte(str_out)

	for idx, char := range byte_str{
		if char == "," {
			if byte_str[idx-1] == " " {
				byte_str[idx-1] = ""
			}
		}
	}
	fmt.Println("")
	fmt.Println(str_out)
}
