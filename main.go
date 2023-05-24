package main

import (
	"fmt"
	"regexp"
	//"strconv"
	"strings"
	"os"
)

func main() {
	file_name := os.Args[1]                        // Reading the first argument after program name
	dat, _ := os.ReadFile(file_name)               // Read the file and get the data as byte[]
	str := string(dat)

	reCap := regexp.MustCompile(`([^[:space:]]+)\s+\(cap\)`) // 
	modified := reCap.ReplaceAllStringFunc(str, capitalize)

	reUp := regexp.MustCompile(`([^\s]+)\s+\(up\)`)
	modified = reUp.ReplaceAllStringFunc(modified, ToUp)

	reQaws := regexp.MustCompile(`\s*\((\w+),\s*(\d+)\)`) // spcae (word, number)
	modified = reQaws.ReplaceAllStringFunc(modified, qawsHandler)
	fmt.Println(modified)
}


func capitalize(match string) string {
	return strings.Title(match[0 : len(match)-6])
}

func ToUp(match string) string {
	return strings.ToUpper(match[0 : len(match)-5])
}

func qawsHandler(match string) string {
	fmt.Println(match[0: len(match)-])
	return "QAWS"
}