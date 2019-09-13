package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func getInput() string {
	input := make([]string, 0)
	var token string
	for {
		_, err := fmt.Scan(&token)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}
		input = append(input, token)
	}
	return strings.Join(input, " ")
}

var argInput = flag.String("i", "", "A string of text to strike through.")

func main() {
	flag.Parse()
	input := *argInput
	if *argInput == "" {
		input = getInput()
	}
	output := make([]rune, 0)
	for _, r := range input {
		output = append(output, []rune{r, 822}...)
	}
	fmt.Printf("%s\n", string(output))
}
