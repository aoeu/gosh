package main

func strikethrough(input string) string {
	output := make([]rune, 0)
	for _, r := range input {
		output = append(output, []rune{r, 822}...)
	}
	return string(output)
}
