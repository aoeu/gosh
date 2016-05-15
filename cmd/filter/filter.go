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
	find . -name '*.yava' | {{.}} generated-sources target test
	cat works_of_shakespeare.txt | {{.}} thou thee thine

flags:

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
	// TODO(aoeu): Print default flags before printing usage examples.
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	args := struct {
		matchAll bool
	}{}
	flag.BoolVar(&args.matchAll, "all", false, "Lines ommitted must match all filters (instead of any filter).")
	flag.Parse()
	filters := flag.Args()
	if args.matchAll && len(filters) == 0 {
		flag.Usage()
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		matchedAny := false
		matchedAll := true
		for _, f := range filters {
			if strings.Contains(t, f) {
				matchedAny = true
			} else {
				matchedAll = false
			}
		}
		switch {
		case args.matchAll && matchedAll:
		case !args.matchAll && matchedAny:
		default:
			fmt.Fprintln(os.Stdout, t)
		}
	}
}
