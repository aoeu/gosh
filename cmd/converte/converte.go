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
		amount   int
		decrypt  bool
		emote    bool
	}{}
	flag.StringVar(&args.funcName, "with", "rotate", "the name of the way to convert input text")
	flag.IntVar(&args.amount, "amount", 0, `czar: Amount of letter offsets to transpose by.
trails: The amount of trails to add`)
	flag.BoolVar(&args.decrypt, "decrypt", false, "czar: Decrypt text.")
	flag.BoolVar(&args.emote, "emote", false, "rotate: Flip text with added emoji")
	flag.Parse()
	args.input = strings.Join(flag.Args(), " ")
	if args.input == "" {
		args.input = scanInput()
	}
	var s string
	switch args.funcName {
	case "rotate":
		s = rotate(args.input, args.emote)
	case "czar":
		s = czar(args.input, args.amount, args.decrypt)
	case "tiny":
		s = tiny(args.input)
	case "strikethrough":
		s = strikethrough(args.input)
	case "trails":
		s = addTrails(args.input, args.amount)
	}
	fmt.Println(s)
}
