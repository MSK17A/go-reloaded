package pkgs

import "strings"

func Capitalize(match string) string {
	str := match[0 : len(match)-6]                  // Trims "(cap) "
	str = strings.ToUpper(string(str[0])) + str[1:] // Select the first letter and upper case it then add it to the string
	return str
}
