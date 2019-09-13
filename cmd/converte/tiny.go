package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var runeMap = map[rune][]rune{
	'a': {' ', 867},
	'c': {' ', 872},
	'd': {' ', ' ', 873},
	'e': {' ', 868},
	'h': {' ', 874},
	'i': {' ', 869},
	'l': {' ', ' ', 7646},
	'm': {' ', 875},
	'o': {' ', ' ', 870},
	'r': {' ', 876},
	't': {' ', 877},
	'u': {' ', 871},
	'v': {' ', 878},
	'x': {' ', 879},
}

func shrink(in string) (out string) {
	for _, r := range in {
		r2, ok := runeMap[r]
		if !ok {
			r2 = []rune{r}
		}
		out += string(r2)
	}
	return
}

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

var argInput = flag.String("i", "", "A string of text to transform.")

func main() {
	flag.Parse()
	input := *argInput
	if *argInput == "" {
		input = getInput()
	}
	shrunk := shrink(input)
	fmt.Printf("%s\n", shrunk)
}
