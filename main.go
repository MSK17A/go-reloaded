package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_name := os.Args[1:]                       // Reading the first argument after program name
	dat, _ := os.ReadFile(string(file_name[0]))    // Read the file and get the data as byte[]
	string_list := strings.Split(string(dat), " ") // Conv string to a List of strings
	fmt.Println(string_list)

	for idx, word := range string_list {
		switch word {
		case "(cap)":
			{
				str := string_list[idx-1]
				string_list[idx-1] = strings.ToUpper(string(str[0])) + str[1:] // Select the first letter and upper case it then add it to the string
				string_list[idx] = ""                                          // Replece with empty string to mark it as deleted
			}
		case "(up)":
			{
				string_list[idx-1] = strings.ToUpper(string_list[idx-1]) // Same as above, but convert all the word to upperCase
				string_list[idx] = ""
			}
		case "(low)":
			{
				string_list[idx-1] = strings.ToLower(string_list[idx-1])
				string_list[idx] = ""
			}
		case "(hex)":
			{
				num, _ := strconv.ParseInt(string_list[idx-1], 16, 64) // Take the number in hex, convert it to decimal integer
				string_list[idx-1] = strconv.Itoa(int(num))            // Conver the number to string
				string_list[idx] = ""
			}
		case "(bin)":
			{
				num, _ := strconv.ParseInt(string_list[idx-1], 2, 64)
				string_list[idx-1] = strconv.Itoa(int(num))
				string_list[idx] = ""
			}
		case "(low,":
			{
				num_str := strings.Split(string_list[idx+1], ")") // Seperate the number from the bracket ')'. (low, <number>)
				num, _ := strconv.Atoi(num_str[0])                // Convert the number to integer, so we can use it to lowerCase number of strings behind

				for i := idx; i > idx-num; i-- {
					string_list[i-1] = strings.ToLower(string_list[i-1])
				}
			}
		case "(up,":
			{
				num_str := strings.Split(string_list[idx+1], ")")
				num, _ := strconv.Atoi(num_str[0])

				for i := idx; i > idx-num; i-- {
					string_list[i-1] = strings.ToUpper(string_list[i-1])
				}
			}
		}
	}
	fmt.Println("")
	fmt.Println(string_list)

}
