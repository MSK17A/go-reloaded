package pkgs

import (
	"regexp"
	"strconv"
)

func Cap_qaws_with_number(str string) string {
	reQaws_regexp := `(^|\s+)\(\s*(cap)\s*,\s*(\d+)\s*\)(\s+|$)`
	reQaws := regexp.MustCompile(reQaws_regexp) // spcae (word, number) space
	qaws_match := reQaws.FindStringSubmatch(str)
	if len(qaws_match) < 4 {
		return str
	}
	digit, _ := strconv.Atoi(qaws_match[3])
	reQaws_regexp = Words_Adder(reQaws_regexp, digit)
	reQaws = regexp.MustCompile(reQaws_regexp)
	str_out := reQaws.ReplaceAllStringFunc(str, NCapitalize)
	reQaws_regexp = `(^|\s+)\(\s*(cap)\s*,\s*(\d+)\s*\)(\s+|$)`
	reQaws = regexp.MustCompile(reQaws_regexp)
	str_out = reQaws.ReplaceAllString(str_out, " ")

	return str_out
}

func Low_qaws_with_number(str string) string {
	reQaws_regexp := `(^|\s+)\(\s*(low)\s*,\s*(\d+)\s*\)(\s+|$)`
	reQaws := regexp.MustCompile(reQaws_regexp) // spcae (word, number) space
	qaws_match := reQaws.FindStringSubmatch(str)
	if len(qaws_match) < 4 {
		return str
	}
	digit, _ := strconv.Atoi(qaws_match[3])
	reQaws_regexp = Words_Adder(reQaws_regexp, digit)
	reQaws = regexp.MustCompile(reQaws_regexp)
	str_out := reQaws.ReplaceAllStringFunc(str, NLow)
	reQaws_regexp = `(^|\s+)\(\s*(low)\s*,\s*(\d+)\s*\)(\s+|$)`
	reQaws = regexp.MustCompile(reQaws_regexp)
	str_out = reQaws.ReplaceAllString(str_out, " ")

	return str_out
}
