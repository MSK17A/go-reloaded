package main

import (
	"fmt"
	"regexp"
<<<<<<< HEAD
	//"strconv"
=======
	"strconv"

	//"strconv"
	"os"
>>>>>>> 6387139 (Added some text manioulation functions using by regexp)
	"strings"
	"os"
)

func main() {
<<<<<<< HEAD
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
=======
	file_name := os.Args[1]          // Reading the first argument after program name
	dat, _ := os.ReadFile(file_name) // Read the file and get the data as byte[]
	str := string(dat)

	reCap := regexp.MustCompile(`([^[:space:]]+)\s+\(cap\)\s*`) // Selects this expression " (cap) "
	modified := reCap.ReplaceAllStringFunc(str, capitalize)

	reUp := regexp.MustCompile(`([^\s]+)\s+\(up\)\s*`) //
	modified = reUp.ReplaceAllStringFunc(modified, ToUp)

	reLow := regexp.MustCompile(`([^\s]+)\s+\(low\)\s*`) //
	modified = reLow.ReplaceAllStringFunc(modified, ToLow)

	reBin := regexp.MustCompile(`([^\s]+)\s+\(bin\)`) //
	modified = reBin.ReplaceAllStringFunc(modified, bin_to_dec)

	reHex := regexp.MustCompile(`([^\s]+)\s+\(hex\)`) //
	modified = reHex.ReplaceAllStringFunc(modified, hex_to_dec)

	/*reQaws := regexp.MustCompile(`\s*\((\w+),\s*(\d+)\)`) // spcae (word, number)
	modified = reQaws.ReplaceAllStringFunc(modified, qawsHandler)*/
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
	fmt.Println(match[0:len(match)])
	return "QAWS"
>>>>>>> 6387139 (Added some text manioulation functions using by regexp)
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