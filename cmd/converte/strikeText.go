package main

import (
	"flag"
	"fmt"
)

func strikeText() {
	flag.Parse()
	input := *argInput
	if *argInput == "" {
		input = scanInput()
	}
	output := make([]rune, 0)
	for _, r := range input {
		output = append(output, []rune{r, 822}...)
	}
	fmt.Printf("%s\n", string(output))
}
