package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file_name := os.Args[1:]						// Reading the first argument after program name
	dat,_ := os.ReadFile(string(file_name[0]))		// Read the file and get the data as string
	string_list := strings.Split(string(dat), " ")	// Conv string to a List
	fmt.Println(string_list)
}
