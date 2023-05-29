package main

import (
	"bufio"
	"fmt"
	"regexp"

	//"strconv"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Add a file as argument")
		return
	}

	if len(os.Args) == 2 {
		fmt.Println("Error: Put a name for the output file in the second argument!")
		return
	}

	if len(os.Args) > 3 {
		fmt.Println("Error: Add only one input file and one output name.")
		return
	}

	file_name := os.Args[1] // Reading the first argument after program name
	output_name := os.Args[2]
	dat, err := os.ReadFile(file_name) // Read the file and get the data as byte[]
	if err != nil {
		fmt.Println("Error in reading the file!")
		return
	}
	if len(dat) < 1 {
		fmt.Println("Empty file!")
		return
	}
	str := string(dat)

	reCap := regexp.MustCompile(`([^\s]+)\s+\(cap\)\s*`) // Selects this expression " (cap) "
	modified := reCap.ReplaceAllStringFunc(str, Capitalize)

	reUp := regexp.MustCompile(`([^\s]+)\s+\(up\)\s*`) // Selects this expression " (up) "
	modified = reUp.ReplaceAllStringFunc(modified, ToUp)

	reLow := regexp.MustCompile(`([^\s]+)\s+\(low\)\s*`) // Selects this expression " (low) "
	modified = reLow.ReplaceAllStringFunc(modified, ToLow)

	reBin := regexp.MustCompile(`([^\s]+)\s+\(bin\)`) // Selects this expression " (bin) "
	modified = reBin.ReplaceAllStringFunc(modified, BinToDec)

	reHex := regexp.MustCompile(`([^\s]+)\s+\(hex\)`) // Selects this expression " (hex) "
	modified = reHex.ReplaceAllStringFunc(modified, HexToDec)

	string_list := strings.Split(string(modified), " ") // Conv string to a List of strings
	for idx, word := range string_list {
		if word == "(low," {
			nLow(string_list, idx)
		} else if word == "(up," {
			nUp(string_list, idx)
		} else if word == "(cap," {
			nCapitalize(string_list, idx)
		}
	}
	/*reQaws := regexp.MustCompile(`\s*\((\w+),\s*(\d+)\)`) // spcae (word, number)
	modified = reQaws.ReplaceAllStringFunc(modified, qawsHandler)*/

	// Re-construct the string from array
	str_out := ""
	for _, word := range string_list {
		if word != "" {
			str_out += word
			str_out += " "
		}
	}

	rePunct := regexp.MustCompile(`\s+([[:punct:]]{1,})\s*`) // handle the spaces before and after the punctuation, also works for group of puncts.
	str_out = rePunct.ReplaceAllStringFunc(str_out, Re_Punct)

	reQuote := regexp.MustCompile(`\'[^']+\'`) // handle the quotation marks ''
	str_out = reQuote.ReplaceAllStringFunc(str_out, Re_Quote)

	vowels := "aAeEiIoOuUhH"
	// Loop through all vowels
	for _, vowel := range vowels {
		reVowls := regexp.MustCompile(`\sa\s` + string(vowel)) // handle the vowels
		str_out = reVowls.ReplaceAllStringFunc(str_out, Re_Vowls)
	}

	// Remove last space
	if str_out[len(str_out)-1] == ' ' {
		str_out = str_out[0 : len(str_out)-1]
	}
	// fmt.Println(str_out)

	f_output, err := os.Create(output_name)
	if err != nil {
		fmt.Println("Error creating file!")
		return
	}
	w := bufio.NewWriter(f_output)
	_, err = w.WriteString(str_out)
	if err != nil {
		fmt.Println("Error writing file!")
		return
	}

	w.Flush()
}

/*func qawsHandler(match string) string {
	//fmt.Println(match[0:len(match)])
	return ""
}*/
