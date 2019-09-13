package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func scanInput() string {
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

func main() {
	// TODO(aoeu): Implement
	args := struct {
		input    string
		funcName string
	}{}
	flag.StringVar(&args.funcName, "with", "rotate", "the name of the way to convert input text")
	flag.Parse()
	args.input = strings.Join(flag.Args(), " ")
	if args.input == "" {
		args.input = scanInput()
	}
	var s string
	switch args.funcName {
	case "rotate":
		s = rotate(args.input)
	case "czar":
		s = czar(args.input)
	case "tiny":
		s = tiny(args.input)
	case "strike":
		s = strikeText(args.input)
	}
	fmt.Println(s)
}