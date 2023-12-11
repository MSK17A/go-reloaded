package pkgs

import "strings"

/*
func NCapitalize(string_list []string, idx int) {
	num_str := strings.Split(string_list[idx+1], ")")
	num, _ := strconv.Atoi(num_str[0])

	for i := idx; i > idx-num; i-- {
		str := string_list[i-1]
		string_list[i-1] = strings.ToUpper(string(str[0])) + str[1:]
	}

	string_list[idx] = ""
	string_list[idx+1] = ""
}
*/

func NCapitalize(match string) string {
	matches := strings.Fields(match)
	result := ""
	for idx := range matches {
		result += strings.ToUpper(string(matches[idx][0])) + matches[idx][1:] + " "
	}
	return result
}
