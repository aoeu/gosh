package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} -all [REGULAR EXPRESSION] -with [REPLACEMENT TEXT]

"{{.}}" reads text from standard input, searches for all text that matches 
a supplied regular expression, replaces the matching text with supplied replacement text,
and outputs the resulting text to standard output.

Example:

	$ echo '123. One Two Three' | {{.}} -all '^\d+\.' -with 'Testing:'
	> Testing: One Two Three

	$ echo '123. One Two Three' | sed -E 's/[0-9]{1,}\.[ ]{1,}/Numbers: /'
	> Numbers: One Two Three
	$ echo '123. One Two Three' | {{.}} -with 'Numbers: ' -all '\d+\.\s+' 
	> Numbers: One Two Three

	echo 'sed -E "s/[0-9]{1,}\.[ ]{1,}/Numbers: /"' | {{.}} -all 'sed.*'  -with '{{.}} -all "\d+\.\s+" -with "Numbers: "'

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		searchExp string
		repText   string
	}{}
	flag.StringVar(&args.searchExp, "all", "", "The regular expression to search for in the input text.")
	flag.StringVar(&args.repText, "with", "", "The literal text to replace any regular expression matches with.")
	flag.Parse()
	searchRegexp, err := regexp.Compile(args.searchExp)
	switch {
	case args.searchExp == "":
	case args.repText == "":
		flag.Usage()
	case err != nil:
		fmt.Fprintf(os.Stderr, "Invalid regular expression supplied as argument: %v\n", err)
		flag.Usage()
	}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Fprintln(os.Stdout, searchRegexp.ReplaceAllString(input.Text(), args.repText))
	}
}