package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_name := os.Args[1]                        // Reading the first argument after program name
	dat, _ := os.ReadFile(file_name)               // Read the file and get the data as byte[]
	string_list := strings.Split(string(dat), " ") // Conv string to a List of strings
	fmt.Println(string_list)

	for idx, word := range string_list {
		if word == "(cap)" {
			str := string_list[idx-1]
			string_list[idx-1] = strings.ToUpper(string(str[0])) + str[1:] // Select the first letter and upper case it then add it to the string
			string_list[idx] = ""                                          // Replece with empty string to mark it as deleted
		} else if word == "(up)" {
			string_list[idx-1] = strings.ToUpper(string_list[idx-1]) // Same as above, but convert all the word to upperCase
			string_list[idx] = ""
		} else if word == "(low)" {
			string_list[idx-1] = strings.ToLower(string_list[idx-1])
			string_list[idx] = ""
		} else if word == "(hex)" {
			num, _ := strconv.ParseInt(string_list[idx-1], 16, 64) // Take the number in hex, convert it to decimal integer
			string_list[idx-1] = strconv.Itoa(int(num))            // Conver the number to string
			string_list[idx] = ""
		} else if word == "(bin)" {
			num, _ := strconv.ParseInt(string_list[idx-1], 2, 64)
			string_list[idx-1] = strconv.Itoa(int(num))
			string_list[idx] = ""
		} else if word == "(low," {
			num_str := strings.Split(string_list[idx+1], ")") // Seperate the number from the bracket ')'. (low, <number>)
			num, _ := strconv.Atoi(num_str[0])                // Convert the number to integer, so we can use it to lowerCase number of strings behind

			for i := idx; i > idx-num; i-- {
				string_list[i-1] = strings.ToLower(string_list[i-1])
			}
			string_list[idx] = ""
			string_list[idx+1] = ""
		} else if word == "(up," {
			num_str := strings.Split(string_list[idx+1], ")")
			num, _ := strconv.Atoi(num_str[0])

			for i := idx; i > idx-num; i-- {
				string_list[i-1] = strings.ToUpper(string_list[i-1])
			}
			string_list[idx] = ""
		} else if word == "(cap," {
			num_str := strings.Split(string_list[idx+1], ")")
			num, _ := strconv.Atoi(num_str[0])

			for i := idx; i > idx-num; i-- {
				str := string_list[i-1]
				string_list[i-1] = strings.ToUpper(string(str[0])) + str[1:]
			}
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
	fmt.Println("")
	fmt.Println(str_out)

}
