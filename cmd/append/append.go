package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} -to [file] [text]

'{{.}}' takes quoted arguments and appends them to the specified file,
inserting newlines into the file after appending each argument.

Examples:

	{{.}} -to world hello

	{{.}} -to protips.txt 'files can be easily clobberred with a typo of > instead of >> ' \
		'when appending text to files via output redirection operators, ' \
		'such as:  $ "stuff >> filename.txt"'

	{{.}} -to /tmp/out "foo bar baz" 'zip ding pop' {do re mi}

	{{.}} -to ~/install_notes.txt 'gsettings set org.mate.session.required-components windowmanager i3' {gsettings set org.mate.session required-components-list "['windowmanager', 'panel']"}

Flags:

`

func main() {
	args := struct {
		outputFilepath string
		inputText      string
	}{}
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.StringVar(&args.outputFilepath, "to", "", "the path to the file to append text to")
	flag.Parse()
	args.inputText = strings.Join(flag.Args(), "\n")
	if args.inputText == "" {
		fmt.Fprintf(os.Stderr, "no input text was provided to append to a file\n")
		flag.Usage()
	}
	args.inputText += "\n"
	if args.outputFilepath == "" {
		fmt.Fprintf(os.Stderr, "an output file must be provided\n")
		flag.Usage()
	}
	f, err := os.OpenFile(args.outputFilepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open file at path '%v': %v\n", args.outputFilepath, err)
		os.Exit(1)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "could not cleanly close '%v': %v\n", args.outputFilepath, err)
		}
	}()
	if _, err := f.Write([]byte(args.inputText)); err != nil {
		fmt.Fprintf(os.Stderr, "could not write '%v' to '%v': %v\n", args.inputText, args.outputFilepath, err)
	}
}