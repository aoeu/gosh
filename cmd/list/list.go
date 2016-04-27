package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

var usageTemplate = `usage: {{.}} [ files ]

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
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
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
