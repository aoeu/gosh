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

var argInput = flag.String("i", "", "A string of text to transform.")

func main() {
	// TODO(aoeu): Implement
}