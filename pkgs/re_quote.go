package pkgs

import "strings"

func Re_Quote(match string) string {
	inside_quote := match[1 : len(match)-1]
	inside_quote = strings.Trim(inside_quote, " ")
	//fmt.Println(inside_quote)
	return "'" + inside_quote + "'"
}
