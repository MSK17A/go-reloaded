package main

import (
	"fmt"
	"regexp"
	"strconv"

	//"strconv"
	"os"
	"strings"
)

func main() {
	file_name := os.Args[1]          // Reading the first argument after program name
	dat, _ := os.ReadFile(file_name) // Read the file and get the data as byte[]
	str := string(dat)

	reCap := regexp.MustCompile(`([^\s]+)\s+\(cap\)\s*`) // Selects this expression " (cap) "
	modified := reCap.ReplaceAllStringFunc(str, capitalize)

	reUp := regexp.MustCompile(`([^\s]+)\s+\(up\)\s*`) //
	modified = reUp.ReplaceAllStringFunc(modified, ToUp)

	reLow := regexp.MustCompile(`([^\s]+)\s+\(low\)\s*`) //
	modified = reLow.ReplaceAllStringFunc(modified, ToLow)

	reBin := regexp.MustCompile(`([^\s]+)\s+\(bin\)`) //
	modified = reBin.ReplaceAllStringFunc(modified, bin_to_dec)

	reHex := regexp.MustCompile(`([^\s]+)\s+\(hex\)`) //
	indx := reHex.FindStringIndex(modified)
	fmt.Println(indx)
	modified = reHex.ReplaceAllStringFunc(modified, hex_to_dec)

	reQaws := regexp.MustCompile(`\s*\((\w+),\s*(\d+)\)`) // spcae (word, number)
	modified = reQaws.ReplaceAllStringFunc(modified, qawsHandler)

	rePunct := regexp.MustCompile(`\s+,\s+`) // handle the spaces before and after the punctuation
	modified = rePunct.ReplaceAllStringFunc(modified, Re_Punct)
	fmt.Println(modified)
}

func capitalize(match string) string {
	str := match[0 : len(match)-6]                  // Trims "(cap) "
	str = strings.ToUpper(string(str[0])) + str[1:] // Select the first letter and upper case it then add it to the string
	return str
}

func ToUp(match string) string {
	return strings.ToUpper(match[0 : len(match)-6]) // Trims "(up) "
}

func ToLow(match string) string {
	return strings.ToLower(match[0 : len(match)-6]) // Trims "(low) "
}

func bin_to_dec(match string) string {
	str := match[0 : len(match)-6]
	num, _ := strconv.ParseInt(str, 2, 64)
	return strconv.Itoa(int(num))
}

func hex_to_dec(match string) string {
	str := match[0 : len(match)-6]
	num, _ := strconv.ParseInt(str, 16, 64)
	return strconv.Itoa(int(num))
}

func qawsHandler(match string) string {
	//fmt.Println(match[0:len(match)])
	return ""
}

func Re_Punct(match string) string {
	return ", "
}
