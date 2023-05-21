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
				string_list[idx-1] = strings.ToUpper(string(str[0])) + str[1:]
				string_list[idx] = ""
			}
		case "(up)":
			{
				string_list[idx-1] = strings.ToUpper(string_list[idx-1])
				string_list[idx] = ""
			}
		case "(low)":
			{
				string_list[idx-1] = strings.ToLower(string_list[idx-1])
				string_list[idx] = ""
			}
		case "(hex)":
			{
				num, _ := strconv.ParseInt(string_list[idx-1], 16, 64)
				string_list[idx-1] = strconv.Itoa(int(num))
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
				num_str := strings.Split(string_list[idx+1], ")")
				num, _ := strconv.Atoi(num_str[0])

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
