package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [token]...

'{{.}}' removes lines of text from standard input that
contain regular expressions provided in a space-separated list.
Any lines of text that match the filter(s) and constraints
are printed standard output.

Examples:

	find . -name '*.yava' | {{.}} generated-sources target test

	cat works_of_shakespeare.txt | {{.}} thou thee thine

	cat << EOF | filter -accept -all cat dog
		o
		cat
		dog
		cat dog
		dog cat
		tacocat
		dogmaomagod
		grep -v dog | grep -v cat | grep -v 'cat.*dog'
		EOF

Flags:

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		matchAll bool
	}{}
	flag.BoolVar(&args.matchAll, "all", false, "Lines ommitted must match all filters (instead of any filter).")
	flag.Parse()
	filters := flag.Args()
	if len(filters) == 0 {
		flag.Usage()
	}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		matchedAny := false
		matchedAll := true
		for _, f := range filters {
			match, err := regexp.MatchString(f, t)
			switch {
			case err != nil:
				fmt.Println("Filter expression %v produced error %v", f, err)
				os.Exit(1)
			case match:
				matchedAny = true
			default:
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
