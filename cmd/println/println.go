package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: '{{.}} text'

'{{.}}' prints any provided text to standard output followed by a newline character.

Example:

	{{.}} "Hello, $PWD"
	{{.}} "Why not use" the echo command?  'http://www.in-ulm.de/~mascheck/various/echo+printf'

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	var s string
	if len(os.Args) > 1 {
		// TODO(aoeu): Preserve arbitrary spacing provided with runtime arguments.
		s += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(s)
}