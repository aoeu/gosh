package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var runeMap = map[rune]rune{
	'a': 867,
	'c': 872,
	'd': 873,
	'e': 868,
	'h': 874,
	'i': 869,
	'm': 875,
	'o': 870,
	'r': 876,
	't': 877,
	'u': 871,
	'v': 878,
	'x': 879,
	'z': 7654,
}

func trails(in string, num int) (out string) {
	num = 3
	for _, r := range in {
		r2, ok := runeMap[r]
		if ok {
			runes := []rune{r}
			for i := 0; i < num; i++ {
				runes = append(runes, r2)
			}
			out += string(runes)
		} else {
			out += string(r)
		}
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
var num = flag.Int("t", 1, "Number of times to transform each letter.")

func main() {
	flag.Parse()
	input := *argInput
	if *argInput == "" {
		input = getInput()
	}
	fmt.Printf("%s\n", trails(input, *num))
}
