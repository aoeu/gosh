package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [regular expression]...

'{{.}}' prints lines of text from standard input that match
regular expressions provided in a space-separated list.

Any lines of text that do not match the regular expressions are
ommitted from standard output.

Examples:

	find . | {{.}} '.*\..ava' > yava_and_yavascript_filenames.txt

	cat << EOF || {{.}} dog > /tmp/no_cats.txt
		cat
		dog
		cat dog
		dog cat
		EOF

Flags:

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		matchAll bool
	}{}
	flag.BoolVar(&args.matchAll, "all", false, "Lines accepted must match all regular expressions provided (instead of any).")
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
				fmt.Fprintf(os.Stderr, "Regular expression '%v' produced error %v\n", f, err)
				os.Exit(1)
			case match:
				matchedAny = true
			default:
				matchedAll = false
			}
		}
		switch {
		case args.matchAll && !matchedAll:
		case !matchedAny:
		default:
			fmt.Fprintln(os.Stdout, t)
		}
	}
}
