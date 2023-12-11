package pkgs

func Words_Adder(regexp_qaws string, num_of_words int) string {
	result := regexp_qaws

	for i := 0; i < num_of_words; i++ {
		result = `\w*\s*` + result
	}

	return result
}
