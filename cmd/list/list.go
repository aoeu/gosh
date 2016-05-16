package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aoeu/gosh"
)

var usageTemplate = `usage: {{.}} [file]...

'{{.}}' lists the files in the current directory in an actual list,
instead of columns, which is dissimilar from the 'ls' command in 
Unix-like systems, but is similar to 'ls' of the Plan9 
operating system (or the 'ls -1' command in Unix-like systems).

A glob expression or arbirtary list of files may be provided as arguments.

examples:

	{{.}} 
	{{.}} a*
	{{.}} *.txt
	{{.}} foo.bar *.fiz qux.baz *.buz

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	var files []string
	var err error
	switch len(os.Args) {
	case 1:
		files, err = filepath.Glob("*")
	case 2:
		files, err = filepath.Glob(os.Args[1])
	default:
		files = os.Args[1:]
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, f := range files {
		fmt.Println(f)
	}
}
