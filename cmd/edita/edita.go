package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage '{{.}} filepath [filepath...]'

'{{.}}' opens the specified files in the specified editor, or the editor
referenced in the EDITOR environment variable, or exits with an
error if no editor was specified and the EDITOR environment variable is not set.

Example:

	EDITOR=$PLAN9PORT/bin/acme export EDITOR && {{.}} /tmp/file1.txt

	go get github.com/aoeu/acme/A && {{.}} -with A /tmp/file1.txt /tmp/file2.txt
`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		editorPath string
	}{}
	flag.StringVar(&args.editorPath, "with", "", "The text editor to edit text files with.")
	flag.Parse()
	if args.editorPath == "" {
		args.editorPath = os.Getenv("EDITOR")
	}
	switch {
	case args.editorPath == "":
		fmt.Fprintln(os.Stderr, "No text editor specificed as an argument and none set in EDITOR environment variable")
		flag.Usage()
	case len(flag.Args()) == 0:
		fmt.Fprintln(os.Stderr, "No text files were provided to edit.")
		flag.Usage()
	}
	// TODO(aoeu): Verify that the editor and arguments are valid file paths before executing a command.
	// TODO(aoeu): Can line editors like sam and ed be supported?
	cmd := exec.Command(args.editorPath, flag.Args()...)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when running editor %v to edit files %v: %v", args.editorPath, flag.Args(), err)
		os.Exit(1)
	}
}
