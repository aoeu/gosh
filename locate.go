package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var usage string = `
Usage: '%s regexp', where "regexp" is a regular expression to match a file name.
Example: %s 'example.*.txt'
`

var filenameRegexp *regexp.Regexp
var locatedPath string

func main() {
	if len(os.Args) != 2 {
		s := os.Args[0]
		fmt.Fprintf(os.Stderr, usage, s, s)
		os.Exit(1)
	}
	filenameRegexp = regexp.MustCompile(os.Args[1])
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = filepath.Walk(wd, Mark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch locatedPath {
	case "":
		fmt.Fprintf(os.Stderr, "File not found with name matching '%s'\n", os.Args[1])
	default:
		fmt.Println(filepath.Dir(locatedPath))
	}
}

func Mark(path string, info os.FileInfo, err error) error {
	switch {
	case locatedPath != "" || info.IsDir():
		return nil
	case err != nil && err != os.ErrPermission:
		return err
	case filenameRegexp.Match([]byte(info.Name())):
		locatedPath = path
	}
	return nil
}
