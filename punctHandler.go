package main

func punctHandler(string_list []string) {
	for idx, word := range string_list {
		if word == "," || word == "?" || word == "!" || word == "..." {
			string_list[idx] = ""
			for i := idx - 1; i > 0; i-- {
				if string_list[i] != "" {
					string_list[i] += word
					break
				}
			}
		}
	}
}
