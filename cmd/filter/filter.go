package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `usage: {{.}} [token]...

'{{.}}' removes lines of text from standard input that contain
text tokens provided in a space-separated list. Any lines of
text do not contain the provided text tokens are printed to
standard output.

examples:

	find . -name '*.yava' | {{.}} generated-sources target test
	cat works_of_shakespeare.txt | {{.}} thou thee thine

flags:

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
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
