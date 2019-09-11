package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}}

'{{.}}' outputs the name of the working directory being run from.

Examples:

	{{.}}

	sh -c 'test $({{.}}) = $(basename $PWD) && echo esta {{.}}'

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	d := filepath.Dir(os.Args[0])
	a, err := filepath.Abs(d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred obtaining absolute path of '%v': %v\n", d, err)
	}
	fmt.Println(filepath.Base(a))
}