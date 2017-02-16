package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage '{{.}} filepath [filepath...]'

'{{.}}' opens the specified files in the specified editor, or the editor
referenced in the EDITOR environment variable, or exits with an
error if no editor was specified and the EDITOR environment variable is not set.

The absolute path of the text editor may be specified as an argument,
or just the name of the text editor if the text editor exists in any directory
specified within the PATH environment variable.

Example:

	EDITOR=$PLAN9PORT/bin/acme export EDITOR && {{.}} /tmp/file1.txt

	go get github.com/aoeu/acme/A && {{.}} -with A /tmp/file1.txt /tmp/file2.txt

	find . -name '*.go' | edita
`

func editorExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		editorPath string
	}{}
	flag.StringVar(&args.editorPath, "with", "", "The text editor to edit text files with.")
	flag.Parse()
	switch {
	case args.editorPath == "":
		args.editorPath = os.Getenv("EDITOR")
	case !editorExists(args.editorPath):
		args.editorPath, _ = exec.LookPath(args.editorPath)
	}
	files := flag.Args()
	if len(files) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if s := strings.Trim(input.Text(), "\t\n"); s != "" {
				files = append(files, s)
			}
		}
	}
	switch {
	case args.editorPath == "":
		fmt.Fprintln(os.Stderr, "No text editor specificed as an argument and none set in EDITOR environment variable")
		flag.Usage()
	case !editorExists(args.editorPath):
		fmt.Fprintf(os.Stderr, "No editor exists at the specified path: %v\n", args.editorPath)
		flag.Usage()
	case len(files) == 0:
		fmt.Fprintln(os.Stderr, "No text files were provided to edit.")
		flag.Usage()
	}
	invalidPaths := make([]string, 0)
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			p := strings.Split(f, "/")
			dir := strings.Join(p[0:len(p)-1], "/")
			if _, err = os.Stat(dir); err != nil {
				invalidPaths = append(invalidPaths, f)
			}
		}
	}
	if len(invalidPaths) != 0 {
		fmt.Fprintf(os.Stderr, "Some paths of files to edit are both files that do not exist and have paths with non-existant directories: %v\n", invalidPaths)
		os.Exit(1)
	}
	cmd := exec.Command(args.editorPath, files...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when running editor %v to edit files %v: %v", args.editorPath, files, err)
		os.Exit(1)
	}
}
