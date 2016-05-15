package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var usageTemplate = `usage: {{.}} [token]...

'{{.}}' removes lines of text from standard input that contain any
of text tokens provided in a space-separated list. Any lines of
text do not contain the provided text tokens are printed to
standard output.

examples:
	find . -name '*.yava' | filter generated-sources target test

`

func usage() {
	var t *template.Template
	var err error
	if t, err = template.New("usage").Parse(usageTemplate); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := t.Execute(os.Stdout, os.Args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := t.Execute(os.Stdout, os.Args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	filters := os.Args
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		isMatch := false
		for _, f := range filters {
			if strings.Contains(t, f) {
				isMatch = true
			}
		}
		if !isMatch {
			fmt.Fprintln(os.Stdout, t)
		}
	}
}
