package main

import (
	"bufio"
	"fmt"
	"goReloaded/pkgs"
	"regexp"
	"strings"

	//"strconv"
	"os"
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

	reCap := regexp.MustCompile(`([^\s]+)\s+\(cap\)\s+`) // Selects this expression " (cap) "
	modified := reCap.ReplaceAllStringFunc(str, pkgs.Capitalize)

	reUp := regexp.MustCompile(`([^\s]+)\s+\(up\)\s+`) // Selects this expression " (up) "
	modified = reUp.ReplaceAllStringFunc(modified, pkgs.ToUp)

	reLow := regexp.MustCompile(`([^\s]+)\s+\(low\)\s+`) // Selects this expression " (low) "
	modified = reLow.ReplaceAllStringFunc(modified, pkgs.ToLow)

	reBin := regexp.MustCompile(`([^\s]+)\s+\(bin\)\s+`) // Selects this expression " (bin) "
	modified = reBin.ReplaceAllStringFunc(modified, pkgs.BinToDec)

	reHex := regexp.MustCompile(`([^\s]+)\s+\(hex\)\s+`) // Selects this expression " (hex) "
	modified = reHex.ReplaceAllStringFunc(modified, pkgs.HexToDec)

	modified = pkgs.Cap_qaws_with_number(modified) // Numberd brackets CAP
	modified = pkgs.Low_qaws_with_number(modified) // Numberd brackets LOW
	modified = pkgs.Up_qaws_with_number(modified)  // Numberd brackets UP

	rePunct := regexp.MustCompile(`\s+([[:punct:]]{1,})\s*`) // handle the spaces before and after the punctuation, also works for group of puncts.
	modified = rePunct.ReplaceAllStringFunc(modified, pkgs.Re_Punct)

	reQuote := regexp.MustCompile(`\'[^']+\'`) // handle the quotation marks ''
	modified = reQuote.ReplaceAllStringFunc(modified, pkgs.Re_Quote)

	vowels := "aAeEiIoOuUhH"
	// Loop through all vowels
	for _, vowel := range vowels {
		reVowls := regexp.MustCompile(`\sa\s` + string(vowel)) // handle the vowels
		modified = reVowls.ReplaceAllStringFunc(modified, pkgs.Re_Vowls)
	}

	modified = strings.TrimSpace(modified)
	// fmt.Println(str_out)

	f_output, err := os.Create(output_name)
	if err != nil {
		fmt.Println("Error creating file!")
		return
	}
	w := bufio.NewWriter(f_output)
	_, err = w.WriteString(modified)
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
