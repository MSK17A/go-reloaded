package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//file_name := os.Args[1]                        // Reading the first argument after program name
	dat, _ := os.ReadFile("sample.txt")            // Read the file and get the data as byte[]
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
		} else if word == "(up," {
			nUp(string_list, idx)
		} else if word == "(cap," {
			nCapitalize(string_list, idx)
		} else if word != "" && idx < len(string_list)-1 {
			if unicode.IsPunct(rune(word[0])) && !unicode.IsLetter(rune(word[1])) {
				string_list[idx] = ""
				for rev_idx := idx - 1; rev_idx > 0; rev_idx-- {
					if string_list[rev_idx] != "" {
						string_list[rev_idx] += string(word[0:])
						break
					}
				}

			} else if unicode.IsPunct(rune(word[0])) && unicode.IsLetter(rune(word[1])) {
				temp_str := word[1:]
				string_list[idx] = ""
				for rev_idx := idx - 1; rev_idx > 0; rev_idx-- {
					if string_list[rev_idx] != "" {
						string_list[rev_idx] += string(word[0])
						break
					}
				}
				string_list[idx] = temp_str
			}
		}
	}

	str_out := ""

	for _, word := range string_list {
		if word != "" {
			/*if idx < len(string_list)-1 {
				if string_list[idx+1] != "" {

					if unicode.IsPunct(rune(string_list[idx+1][0])) {
						str_out += word
						continue
					}
				}
			}*/
			str_out += word
			str_out += " "
		}
	}
	fmt.Println("")
	fmt.Println(str_out)
}
