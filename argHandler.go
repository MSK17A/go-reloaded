package main

func argHandler(string_list []string) {
	for idx, word := range string_list {
		if word == "(cap)" {
			Capitalize(string_list, idx)
			string_list[idx] = "" // Replece with empty string to mark it as deleted
		} else if word == "(up)" {
			Upper(string_list, idx)
			string_list[idx] = ""
		} else if word == "(low)" {
			Lower(string_list, idx)
			string_list[idx] = ""
		} else if word == "(hex)" {
			HexToDec(string_list, idx)
			string_list[idx] = ""
		} else if word == "(bin)" {
			BinToDec(string_list, idx)
			string_list[idx] = ""
		} else if word == "(low," {
			nLow(string_list, idx)
		} else if word == "(up," {
			nUp(string_list, idx)
		} else if word == "(cap," {
			nCapitalize(string_list, idx)
		}
	}
}
