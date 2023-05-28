package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	//"strconv"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Add a file as argument")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Error: Add only one file as an argument")
		return
	}

	file_name := os.Args[1]            // Reading the first argument after program name
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
	modified := reCap.ReplaceAllStringFunc(str, capitalize)

	reUp := regexp.MustCompile(`([^\s]+)\s+\(up\)\s*`) // Selects this expression " (up) "
	modified = reUp.ReplaceAllStringFunc(modified, ToUp)

	reLow := regexp.MustCompile(`([^\s]+)\s+\(low\)\s*`) // Selects this expression " (low) "
	modified = reLow.ReplaceAllStringFunc(modified, ToLow)

	reBin := regexp.MustCompile(`([^\s]+)\s+\(bin\)`) // Selects this expression " (bin) "
	modified = reBin.ReplaceAllStringFunc(modified, bin_to_dec)

	reHex := regexp.MustCompile(`([^\s]+)\s+\(hex\)`) // Selects this expression " (hex) "
	modified = reHex.ReplaceAllStringFunc(modified, hex_to_dec)

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

	f_output, err := os.Create("result.txt")
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

/*func qawsHandler(match string) string {
	//fmt.Println(match[0:len(match)])
	return ""
}*/

func Re_Punct(match string) string {
	if strings.ContainsAny(match, "(){}&^'") { // If it is not like these punctionans ?!,... just return them and skip
		return match
	}
	the_isolated_punct := strings.Trim(match, " ") // A placeholder to isolate the actual punctuation from the spaces
	/*for _, char := range match {
		if char != ' ' { // Do not add the spaces
			the_isolated_punct += string(char)
		}
	}*/
	//fmt.Println(the_isolated_punct)

	return the_isolated_punct + " " // add one space after the punctuation
}

func Re_Quote(match string) string {
	inside_quote := match[1 : len(match)-1]
	inside_quote = strings.Trim(inside_quote, " ")
	//fmt.Println(inside_quote)
	return "'" + inside_quote + "'"
}

func Re_Vowls(match string) string {
	return match[0:1] + "an" + match[2:]
}
